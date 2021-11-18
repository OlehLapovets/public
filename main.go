package main

import (
  "fmt"
  "bufio"
  "os"
  "log"
  "strings"
)
var max int = 100
var maxPointer *int = &max
var min int = 0
var minPointer *int = &min

func median(max, min int) int {
  if (max == 100) && (min == 99) {
    return 100
  } else if (max == 1) && (min == 0) {
    return 0
  }
  return (max + min) / 2
}

func answer() string {
  reader := bufio.NewReader(os.Stdin)
  input, err := reader.ReadString('\n')
  if err != nil {
    log.Fatal(err)
  }
  input = strings.TrimSpace(input)
  return input
}

func main() {
  fmt.Println("--------------------------")
  fmt.Printf("Go: Загадайте число від %v до %v:\n", min, max)
  for i := 9; i >= 0; i -- {
    guess := median(max, min)
    if guess == 100 {
      fmt.Println("Це число 100!")
      break
    } else if guess == 0 {
      fmt.Println("Це число 0!")
      break
    } else if max == min {
      fmt.Printf("це - %v", max)
      break
    }
    fmt.Printf("Go: Це число %v? Якщо ні, то більше чи менше?\n", guess)
    answer := answer()

    if answer == "більше" {
      *minPointer = guess
    } else if answer == "менше" {
      *maxPointer = guess
    } else if answer == "вгадав!"{
      fmt.Println("Go: Я вгадав!")
      break
    } else {
      fmt.Println("Введіть знову...")
      i += 1
    }
  fmt.Println("--------------------------")
  if i == 0 {
    fmt.Println("У мене закінчились спроби((")
    break
  }
  fmt.Println("Go: Блін! У мене залишилось", i, "спроб.")
  }
}
