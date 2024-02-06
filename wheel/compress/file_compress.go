package compress

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

const (
	basePath  = "testdata"
	zipSuffix = "zip"
)

func Compress2ZipFile(dstPath, srcPath string) {
	arcFile, err := os.Create(getFileName(dstPath, true))
	defer arcFile.Close()
	if err != nil {
		panic(err)
	}

	arcWriter := zip.NewWriter(arcFile)
	defer arcWriter.Close()

	helloc, err := arcWriter.Create(getFileName(srcPath))
	if err != nil {
		panic(err)
	}

	srcFile, err := os.Open(getFileName(srcPath))
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()
	if _, err = io.Copy(helloc, srcFile); err != nil {
		panic(err)
	}
}

func UnCompressZipFile(dstPath, srcPath string) {
	zippedFile, err := zip.OpenReader(getFileName(srcPath, true))
	if err != nil {
		panic(err)
	}
	defer zippedFile.Close()
	for i := range zippedFile.File {
		zf := zippedFile.File[i]
		fileName := zf.Name
		if zf.FileInfo().IsDir() {
			_ = os.MkdirAll(fileName, os.ModePerm)
			continue
		}
		err = os.MkdirAll(filepath.Dir(fileName), os.ModePerm)
		if err != nil {
			panic(err)
		}

		destFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zf.Mode())
		if err != nil {
			panic(err)
		}
		defer destFile.Close()

		srcFile, err := zf.Open()
		if err != nil {
			panic(err)
		}
		defer srcFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			panic(err)
		}

	}
}

func getFileName(dst string, suffix ...bool) string {
	if len(suffix) > 0 && suffix[0] {
		return basePath + "/" + dst + "." + zipSuffix
	}
	return basePath + "/" + dst
}
