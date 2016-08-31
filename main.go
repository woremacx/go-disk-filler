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
	defer file.Close()
	writer := bufio.NewWriterSize(file, 1048576)
	buf := make([]byte, 1048576)

	for {
		writer.Write(buf)
		err := writer.Flush()
		if err != nil {
			file.Close()
			return
		}
	}
}
