package lang

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func QRead() {
	quickRead()
}

func quickRead() {

}

func QWrite() {
	quickWrite()
}

func quickWrite() {

}

//inline int read() {
//int s = 0, w = 1;
//char ch = getchar();
//while (ch < '0' || ch > '9') {
//if (ch == '-') w = -1;
//ch = getchar();
//}
//while (ch >= '0' && ch <= '9') {
//s = s * 10 + ch - '0';
//ch = getchar();
//}
//return s * w;
//}

// QReadInt32
// invalid in Golang
func QReadInt32() int32 {
	var (
		ch   rune
		nump int32
		sign bool = true
	)
	_, err := fmt.Scan(&ch)
	if err != nil {
		panic(err)
	}
	for ch < '0' || ch > '9' {
		if ch == '-' {
			sign = false
		}
		_, err = fmt.Scan(&ch)
		if err != nil {
			panic(err)
		}
	}
	for ch >= '0' && ch <= '9' {
		nump = nump<<1 + nump<<3 + ch - '0'
		_, err = fmt.Scan(&ch)
		if err != nil {
			panic(err)
		}
	}
	if !sign {
		nump = -nump
	}
	return nump
}

func ReadFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		return
	}
}

func WriteFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("This is some information to write to file.")
	if err != nil {
		panic(err)
	}
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
	err = file.Close()
	if err != nil {
		panic(err)
	}
}

func LogFile(fileName string) *log.Logger {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return log.New(file, "", log.LstdFlags)

}
