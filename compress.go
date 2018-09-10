package core

import (
	"compress/zlib"
	"os"
)

const defaultPath = "yi.data"

func libCompress(b []byte) error {
	file, err := os.OpenFile(defaultPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	w := zlib.NewWriter(file)
	if _, err := w.Write(b); err != nil {
		return err
	}
	return nil
}

func libDecompress() ([]byte, error) {
	var b []byte
	file, err := os.OpenFile(defaultPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	r, err := zlib.NewReader(file)
	if err == nil {
		if _, err := r.Read(b); err == nil {
			return b, nil
		} else {

		}
	}
	return nil, err
}
