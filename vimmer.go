package main

import (
  "os"
  "os/exec"
  "fmt"
  "log"
)
func main() {
  files := os.Args[1:]
  if len(files) == 0 {
    fmt.Println("Provide your files as a list of arguments.")
    fmt.Println("eg: ls ./* | xargs vimmer")
    return
  }
  for i := 0 ; i < len(files) ; i++ {
    cmd := exec.Command("vim", files[i])
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
      log.Fatal(err)
    }
  }
}
