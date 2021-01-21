package main

import (
	"bufio"
	"fmt"
	"os"
)

func read_file() {
	fPath := "read.txt"

	f, err := os.Open(fPath)
	if err != nil {
		fmt.Printf("file open failed, %v\n", err)
		return
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		fmt.Printf("file read failed, %v\n", err)
	}
}
