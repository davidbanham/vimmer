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
    cmd := exec.Command("mvim", files[i])
    err := cmd.Run()
    if err != nil {
      log.Fatal(err)
    }
  }
}
