package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	fmt.Println("Starting kata 1")
}

func openReadFile(filename, appendStr string) error {
	i, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	fmt.Println(i)
	o := fmt.Sprintf("%s\n%s", i, appendStr)
	os.WriteFile(filename, []byte(o), fs.ModeAppend)
	return nil
}
