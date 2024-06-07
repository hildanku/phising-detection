package main

import (
  "fmt"
  "strings"
)

func main() {
  var link string
  fmt.Println("Masukkan link yang ingin dianalisis:")
  fmt.Scanln(&link)
  fmt.Println("analyze:", link)

  isSafe := strings.Contains(link, "dana.id")

  if (!isSafe) {
    fmt.Println("Not Safe")
  }
  fmt.Println("Safety")
}