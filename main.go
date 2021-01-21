package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func produce(srcFilePath string, data chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	f, err := os.Open(srcFilePath)
	if err != nil {
		fmt.Printf("file open failed, %v\n", err)
		return
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		data <- s.Text()
	}
	err = s.Err()
	if err != nil {
		fmt.Printf("file read failed, %v\n", err)
	}
}

func consume(dstFilePath string, data chan string, done chan bool) {
	f, err := os.OpenFile(dstFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			done <- false
			return
		}
	}

	done <- true
}

func main() {
	srcFiles := []string{"001.txt", "002.txt", "003.txt"}
	dscFile := "merge_result.txt"

	ch := make(chan string, 64)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(len(srcFiles))
	for i := 0; i < len(srcFiles); i++ {
		go produce(srcFiles[i], ch, &wg)
	}

	go consume(dscFile, ch, done)

	go func() {
		wg.Wait()
		close(ch)
	}()

	d := <-done
	if d == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}
}
