package main

import (
	"file_maneger/filework"
	"fmt"
	"os"
)
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: ./program <function>")
		os.Exit(1)
	}

	arg := os.Args[1]
	file_name := os.Args[2]
	f := filework.NewFileManager()
	switch arg {
	case "cat":
		f.Cat(file_name)
	case "tac":
		f.Tac(file_name)
	case "tail":
		f.Tail(file_name)
	default:
		fmt.Println("Неверная функция")
	}
}
