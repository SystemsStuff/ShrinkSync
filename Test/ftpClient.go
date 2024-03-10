package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

var conn *ftp.ServerConn

func create_dir(output_dir string) {
	err := conn.MakeDir(output_dir)
	if err != nil {
		panic(err)
	}
	log.Printf("Successfully created directory: %s\n", output_dir)
}

func upload_file(input_file string, output_file string) {
	file_open, err := os.Open(input_file)
	if err != nil {
		panic(err)
	}
	log.Printf("Successfully opened the input file: %s\n", input_file)

	reader := bufio.NewReader(file_open)
	err = conn.Stor(output_file, reader)
	if err != nil {
		panic(err)
	}
	log.Printf("Successfully uploaded the file: %s to the file %s in FTP server\n", input_file, output_file)
}

func upload_data(data string, output_file string) {
	byte_data := bytes.NewBufferString(data)
	err := conn.Stor(output_file, byte_data)
	if err != nil {
		panic(err)
	}
	log.Printf("Successfully uploaded the data: %s to the file %s in FTP server\n", data, output_file)
}

func download_file(file_path string) string {
	response, err := conn.Retr(file_path)
	if err != nil {
		panic(err)
	}
	defer response.Close()
	log.Printf("Successfully retrieved the file: %s\n", file_path)
	buf, err := io.ReadAll(response)
	return string(buf)
}

func connect(server string, network string, port int) {
	var err error
	conn, err = ftp.Dial(fmt.Sprintf("%s.%s:%d", server, network, port), ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully connected to the server %s in the network %s on port %d\n", server, network, port)

	err = conn.Login("foo", "bar")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully logged in as user foo\n")
}

func main() {
	connect("datagrid", "shrink-sync-network", 21)
	create_dir("TestInput")
	upload_file("test_input", "./TestInput/test_input")
	upload_data("Hello World!", "./TestInput/test_input_1")
	file_content := download_file("./TestInput/test_input")
	fmt.Println("File Content: %s", file_content)
}
