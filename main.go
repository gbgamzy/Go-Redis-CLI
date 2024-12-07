package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	fmt.Println("Listening on port :6379")

	// Create a new server
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer l.Close()

	// Loop to accept multiple connections
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("Connection received")

		// Handle the connection in a separate goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("Connection received")
	defer conn.Close()

	for {
		resp := newResp(conn)
		value, err := resp.Read()
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected")
				return // Exit the loop when the client disconnects
			}
			fmt.Println("Read error:", err)
			return
		}

		if value.typ != "array" {
			fmt.Println("Invalid request, expected array")
			continue
		}

		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		writer := NewWriter(conn)
		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("Invalid Command:", command)
			writer.Write(Value{typ: "string", str: "Invalid Command"})
			continue
		}

		result := handler(args)
		writer.Write(result)
	}
}

type Writer struct {
	writer io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{writer: w}
}

func (w *Writer) Write(v Value) error {
	var bytes = v.Marshal()

	_, err := w.writer.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
