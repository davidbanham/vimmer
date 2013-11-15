package main

import (
  "os"
  "fmt"
  "bufio"
  "os/exec"
  "log"
)
func main() {
  args := os.Args[1:]
  if len(args) == 0 {
    fmt.Println("Provide your files as a list of arguments.")
    fmt.Println("eg: ls ./* | xargs vimmer")
    fmt.Println("There is an optional -m flag to run mvim instead of vim")
    return
  }
  if args[0] == "-m" {
    mvim(args[1:])
  } else {
    vim(args[0:])
  }

}
func check_if_continue() {
  fmt.Println("Enter to continue, anything else to bail")
  reader := bufio.NewReader(os.Stdin)
  resp, err := reader.ReadString('\n')
  if resp != "\n" {
    log.Fatal("Got something other than enter: ", resp)
  }
  if err != nil {
    log.Fatal(err)
  }
}

func vim(files []string) {
  for i := 0 ; i < len(files) ; i++ {
    cmd := exec.Command("vim", files[i])
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
      log.Fatal(err)
    }
    if i != len(files) - 1 {
      check_if_continue()
    }
  }
}

func mvim(files []string) {
  for i := 0 ; i < len(files) ; i++ {
    cmd := exec.Command("mvim", files[i])
    err := cmd.Run()
    if err != nil {
      log.Fatal(err)
    }
  }
}

