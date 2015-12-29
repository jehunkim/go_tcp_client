package main

import (
	"net"
	"log"
	"fmt"
	"os"
	"bufio"
	"bytes"
	"unicode/utf8"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	var buff = make([]byte, 1024)
	
	for {
		var strBuf bytes.Buffer // buffer for print
		fmt.Print("Enter message : ")
		text, err := reader.ReadString('\n')
		conn.Write([]byte(text))
		byteLength, err := conn.Read(buff)
		if err != nil {
			log.Fatal(err.Error())
		}
		for byteLength > 0 { // loop while all of characters are read
			// read rune(calculated byte for unicode)
			// returning rune and size of rune's bytes
			r, size := utf8.DecodeRune(buff)
			strBuf.WriteString(string(r))
			buff = buff[size:] // decrease buffer
			byteLength = byteLength - size // decrease remaining buffer length
		}
		log.Print(strBuf.String())
	}
}
