package main

import (
  "os"
  "os/exec"
  "fmt"
  "log"
  "bufio"
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
    if i != len(files) - 1 {
      check_if_continue()
    }
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
