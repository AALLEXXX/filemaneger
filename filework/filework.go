package filework

import (
	"fmt"
	"os"
	"bufio"

)

type File_maneger interface {
	Cat(string) error
	Tac(string) error
	Tial(string) error
}


func NewFileManager() *MyFileManager {
	return &MyFileManager{}
}


type MyFileManager struct{}

func (f MyFileManager) Cat(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			fmt.Print(string(buffer[:n]))
		}
		if err != nil {
			break
		}
	}
	return nil
}


func (f MyFileManager) Tac(filename string) error {
	stack, _ := get_datafile_stack(filename)
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(top)
	}
	return nil
}

func (f MyFileManager) Tail(filename string) error {
	stack, _ := get_datafile_stack(filename)
	counter := 0 
	for len(stack) > 0 {
		if counter == 10{
			break
		}
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(top)
		counter++ 
	}
	return nil
}


func get_datafile_stack(filename string) ([]string, error){
	file, err := os.Open(filename)
	stack := []string{} 
	if err != nil {
		return stack, err 
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		stack = append(stack, scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		return stack,err
	}

	return stack,err
}