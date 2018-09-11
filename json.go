package core

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

func transFromJson(output interface{}) error {
	file, err := os.OpenFile("yi.json", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	buff := bufio.NewReader(file)

	b := bytes.Buffer{}
	if _, err := io.Copy(&b, buff); err != nil {
		return err
	}

	output = b.Bytes()

	return nil
}
