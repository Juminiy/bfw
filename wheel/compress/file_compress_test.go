package compress

import "testing"

func TestCompress2ZipFile(t *testing.T) {
	Compress2ZipFile("arc", "hello.c")
}

func TestUnCompressZipFile(t *testing.T) {
	UnCompressZipFile("", "arc")
}
