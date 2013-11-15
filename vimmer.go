package main

import (
  "os"
  "fmt"
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
