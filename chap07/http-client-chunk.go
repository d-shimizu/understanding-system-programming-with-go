package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err := nil {
		panic(err)
	}

	request, err := http.NewRequest(
		"GET",
		"http://localhost:8888",
		nil)
	if err != nil {
		panic(err)
	}

	err := request.Write(conn)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(conn)
	response, err := http.DumpResponse(reader, request)
	if err != nil {
		panic(err)
	}

	dump, err := http
