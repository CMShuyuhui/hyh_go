package mypath

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
)

type Book struct {
	Title     string
	Authors   string
	Publisher string
	Price     float32
}

func Chap5() {
	if len(os.Args) != 2 {
		fmt.Println(os.Stderr, "Usage: ", os.Args[0], "host:port")

		fmt.Println(errors.New("start main error").Error())
		os.Exit(1)
	}

	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	fmt.Println(string(result))

	fmt.Println("-------------------------------------------")
	gobooktest := Book{"go语言编程", "huyuhui", "工业出版社", 99.99}

	b, err := json.Marshal(gobooktest)
	if err != nil {
		fmt.Println(errors.New("json.Marshal(gobooktest) error").Error())
		os.Exit(1)
	} else {
		fmt.Println("json marshal success")
	}

	var book Book
	unmarshal_err := json.Unmarshal(b, book)
	if unmarshal_err != nil {
		fmt.Println(errors.New("json.Unmarshal(b, book) error").Error())
		os.Exit(1)
	} else {
		fmt.Println("json Unmarshal success")
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
