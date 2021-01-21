package main

import (
	"fmt"
	"os"
)

func write_file() {
	f, err := os.OpenFile("write.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	newLine := "File handling is easy."
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Printf("file appended failed, %v\n", err)
		return
	}
	fmt.Println("file appended successfully")
}
