package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	httpRequest := "GET / HTTP/1.1\n" +
		"Host: tut.by\n\n"
	conn, err := net.Dial("tcp", "golang.org:80")
	fmt.Println(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	if _, err = conn.Write([]byte(httpRequest)); err != nil {
		fmt.Println("Error: ")
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, conn)
	fmt.Println("Done")
}
