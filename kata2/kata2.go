package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	fmt.Println("kata2")
}

func openReadUpdateJSON(filename, key, value string) error {

	j, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	r := bytes.NewReader(j)
	d := json.NewDecoder(r)
	fmt.Println(d)

	var m map[string]interface{}
	err = d.Decode(&m)
	if err != nil {
		return err
	}

	m[key] = value
	jOut, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jOut, fs.ModeExclusive)
	if err != nil {
		return err
	}
	return nil
}
