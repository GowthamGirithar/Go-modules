package fileUtils

import "testing"

var pathToFile = "" // pass the file path

func TestReadFile(t *testing.T) {
	data, _ := ReadFile(pathToFile)
	println(data)
}

//demo to show examples in go doc
func ExampleReadFile() {
	ReadFile(pathToFile)
}

//for benchmark
func BenchmarkReadFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadFile(pathToFile)
	}
}
