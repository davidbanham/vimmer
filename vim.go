package main

import (
  "os"
  "os/exec"
  "log"
)

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
