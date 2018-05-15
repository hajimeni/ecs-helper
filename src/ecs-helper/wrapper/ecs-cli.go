package wrapper

import (
"os/exec"
"bufio"
"errors"
"net/http"
"path/filepath"
"os"
"io"
"log"
"syscall"
"runtime"
)

const EcsCliCommand = "ecs-cli"
const BinPath = ".bin"

func ExecuteEcsCli(args []string) (int, error) {
    ecsCliBinFilePath, err := GetEcsCliBin()
    if err != nil {
        return -1, err
    }
    log.Println("ecs-cli command args: ", args)

    ecsCmd := exec.Command(ecsCliBinFilePath, args...)

    stdout, err := ecsCmd.StdoutPipe()
    if err != nil {
        return -1, err
    }

    stderr, err := ecsCmd.StderrPipe()
    if err != nil {
        return -1, err
    }

    err = ecsCmd.Start()
    if err != nil {
        return -1, err
    }

    streamReader := func(scanner *bufio.Scanner, outputChan chan string, doneChan chan bool) {
        defer close(outputChan)
        defer close(doneChan)
        for scanner.Scan() {
            outputChan <- scanner.Text()
        }
        doneChan <- true
    }
    stdoutScanner := bufio.NewScanner(stdout)
    stdoutOutputChan := make(chan string)
    stdoutDoneChan := make(chan bool)

    stderrScanner := bufio.NewScanner(stderr)
    stderrOutputChan := make(chan string)
    stderrDoneChan := make(chan bool)

    go streamReader(stdoutScanner, stdoutOutputChan, stdoutDoneChan)
    go streamReader(stderrScanner, stderrOutputChan, stderrDoneChan)

    stillGoing := true
    for stillGoing {
        select {
        case <- stdoutDoneChan:
            stillGoing = false
        case line := <-stdoutOutputChan:
            log.Println(line)
        case line := <-stderrOutputChan:
            log.Println(line)
        }
    }
    err = ecsCmd.Wait()
    if err != nil {
        if exiterr, ok := err.(*exec.ExitError); ok {
            if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
                return status.ExitStatus(), nil
            }
        } else {
            return -1, err
        }
    }
    return 0, nil
}

func GetEcsCliBin() (string, error) {
    if !existsEcsCliBin() {
        downloadUrl := ""
        switch runtime.GOOS {
        case "darwin":
            downloadUrl = "https://s3.amazonaws.com/amazon-ecs-cli/ecs-cli-darwin-amd64-latest"
        case "linux":
            downloadUrl = "https://s3.amazonaws.com/amazon-ecs-cli/ecs-cli-linux-amd64-latest"
        }
        if downloadUrl == "" {
            return "", errors.New("This OS is not supported !")
        }
        log.Printf("Download ecs-cli from %s", downloadUrl)

        response, err := http.Get(downloadUrl)
        if err != nil {
            return "", err
        }

        defer response.Body.Close()
        binPath, err := getBinPath()
        if err:= os.MkdirAll(binPath, 0777); err != nil {
            return "", err
        }
        file, err := os.Create(filepath.Join(binPath, EcsCliCommand))
        if err != nil {
            return "", err
        }

        defer file.Close()
        io.Copy(file, response.Body)
        os.Chmod(filepath.Join(binPath, EcsCliCommand), 0755)
        log.Printf("Downloaded ecs-cli to %s", filepath.Join(binPath, EcsCliCommand))
    }

    binPath, err := getBinPath()
    if err != nil {
        return "", err
    }
    return filepath.Join(binPath, EcsCliCommand), nil
}

func existsEcsCliBin() bool {
    binPath, err := getBinPath()
    if err != nil {
        return false
    }
    file, err := os.Stat(filepath.Join(binPath, EcsCliCommand))
    return err == nil && !file.IsDir()
}

func getBinPath() (string, error) {
    pwd, err := os.Getwd()
    if err != nil {
        return "", err
    }

    return filepath.Join(pwd, BinPath), nil
}
