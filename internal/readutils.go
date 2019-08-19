package fileUtils

import (
	"bufio"
	"io"
	"strings"
)

type ReadDataService interface {
	ReadData(reader io.Reader)
}

//read the data from the reader and send the data in form of string
//can read data from anywhere not only from file
func ReadData(reader io.Reader) string {
	var message strings.Builder
	scan := bufio.NewScanner(reader)
	for scan.Scan() {
		message.WriteString(scan.Text())
	}
	return message.String()
}
