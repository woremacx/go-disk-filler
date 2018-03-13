package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	name := "foo-bar.dat"
	file, err := os.Create(name)
	if err != nil {
		fmt.Printf("error on OpenFile : %s\n", err)
		return
	}

	defer func() {
		file.Close()
		if err := os.Remove(name); err != nil {
			fmt.Printf("error on Remove : %s\n", err)
			return
		}
	}()

	writer := bufio.NewWriterSize(file, 1048576)
	buf := make([]byte, 1048576)
	for {
		writer.Write(buf)
		err := writer.Flush()
		if err != nil {
			break
		}
	}
}
