package main

import (
  "C"
  "fmt"
  "time"
)

func main() {
  timer := time.After(10 * time.Second)
  fmt.Println("timer has started!")
  t := <-timer
  fmt.Println("timer has finished!")
  fmt.Printf("%v", t)
}
