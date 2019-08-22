package fileUtils

import (
	"employee-service/internal"
	"errors"
	"os"
)

//this functionality is mainly to teach testing with mock
func ReadFile(fileName string) (string, error) {
	file, err := os.Open(fileName) // readonly file
	defer file.Close()
	if err != nil {
		return "", errors.New("Error in opering file")
	}
	data := readUtils.ReadData(file) // *file return by os.Open implement the io reader interface
	return data, nil
}
