package main

import (
  "os"
  "fmt"
  "bufio"
  "log"
)

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

