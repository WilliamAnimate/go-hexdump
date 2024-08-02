package main

import (
  "fmt"
  "io/ioutil"
  "os"
)

var file string

func main() {
  args := os.Args
  for _, arg := range args[1:] {
    if arg == "--help" {
      fmt.Printf("No.\n")
      return
    }
    file = arg
  }

  hex, err := ioutil.ReadFile(file)
  if err != nil {
    fmt.Println("i cant read:\n", err)
    return
  }

  for _, b := range hex {
    fmt.Printf("%02x ", b)
  }
  fmt.Printf("\n")
}

