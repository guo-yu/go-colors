package main

import "fmt"
import ".."

func main() {
  fmt.Printf("\n")
  fmt.Printf("Colors list: \n")
  fmt.Printf("\n")
  fmt.Printf("  In %s \n", colors.Print("blue", "blue"))
  fmt.Printf("  In %s \n", colors.Print("yellow", "yellow"))
  fmt.Printf("  In %s \n", colors.Print("green", "green"))
  fmt.Printf("\n")
}