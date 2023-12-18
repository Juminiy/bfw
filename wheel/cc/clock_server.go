package cc

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func NewClockServer(port string) {
	serv, err := net.Listen("tcp", "127.0.0.1:"+port)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := serv.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleClockConn(conn)
	}
}

func handleClockConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("1001-01-01-11:01:01\n"))
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func NewClockClientV0(port string) {
	conn, err := net.Dial("tcp", "127.0.0.1:"+port)
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	writeStdout(os.Stdout, conn)
}

func writeStdout(dest io.Writer, src io.Reader) {
	_, err := io.Copy(dest, src)
	if err != nil {
		panic(err)
	}
}

func ServerClient() {
	go NewClockServer("1202")
	for i := 0; i < 3; i++ {
		go NewClockClientV0("1202")
	}
	NewClockClientV0("1202")
}
