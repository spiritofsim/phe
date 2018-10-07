package pha

import (
	"compress/gzip"
	"io/ioutil"
	"os"
)

func LoadRanks(name string) ([]byte, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return data, nil
}
