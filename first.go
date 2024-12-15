package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    fmt.Print("Введите ваше имя: ")

    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    name := scanner.Text()

    fmt.Printf("Привет, %s!\n", name)
}