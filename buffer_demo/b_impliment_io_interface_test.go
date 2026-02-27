package buffer_demo

import "testing"

func TestIoWrite(t *testing.T) {
	asIoWriter()
}

func Test_asIoReader(t *testing.T) {
	asIoReader()
}

func Test_asReadFrom(t *testing.T) {
	asReadFrom()
}

func Test_asFileWriter(t *testing.T) {
	asFileWriter()
}

func Test_asReaderAndWriter(t *testing.T) {
	asReaderAndWriter()
}
