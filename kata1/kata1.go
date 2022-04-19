package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

func main() {
	fmt.Println("Starting kata 1")
}

func openReadFile(filename, appendStr string) error {
	i, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	fmt.Println(i)
	o := fmt.Sprintf("%s\n%s", i, appendStr)
	ioutil.WriteFile(filename, []byte(o), fs.ModeAppend)
	return nil
}
