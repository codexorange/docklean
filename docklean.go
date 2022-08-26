package main

import (
    "fmt"
    "log"
    "os/exec"
    "strings"
    "bufio"
    "time"
)

func main() {
    fmt.Println("Scanning containers...")
    time.Sleep(2 * time.Second)

    // Remove unused containers
    fmt.Println("Deleting unused containers...")
    time.Sleep(1 * time.Second)
    outputContainers, err := exec.Command("docker", strings.Split("ps -a"," ")...).Output()
    if err != nil {
        log.Fatal(err)
    }

    scannerContainers := bufio.NewScanner(strings.NewReader(string(outputContainers)))
    for scannerContainers.Scan() {
        textLine := scannerContainers.Text()

        if strings.Contains(textLine, "Exited") {
            containerID := textLine[0:12]
            deleted, err := exec.Command("docker", strings.Split(fmt.Sprintf("rm %s", containerID)," ")...).Output()
            if err != nil {
                log.Fatal(err)
            }

            fmt.Println("Container deleted. ID: ", string(deleted))
        }
    }

    fmt.Println("Scanning images...")
    time.Sleep(2 * time.Second)

    // Remove unused images
    fmt.Println("Deleting unused images...")
    time.Sleep(1 * time.Second)
    outputImages, err := exec.Command("docker", strings.Split("images"," ")...).Output()
    if err != nil {
        log.Fatal(err)
    }

    scannerImages := bufio.NewScanner(strings.NewReader(string(outputImages)))
    for scannerImages.Scan() {
        textLine := scannerImages.Text()

        if strings.Contains(textLine, "<none>") {
            imageID := textLine[65:77]
            fmt.Println("Image deleted. ID: ", string(imageID))
            _, err := exec.Command("docker", strings.Split(fmt.Sprintf("rmi %s", imageID), " ")...).Output()
            if err != nil {
                log.Fatal(err)
            }

            fmt.Println("Image deleted. ID: ", string(imageID))
        }
    }
}

