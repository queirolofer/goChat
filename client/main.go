package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const exitCode = "@#quit\n"

func main() {
	fmt.Println("Inicio del cliente")

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Error al conectarse con el servidor")
		return
	}
	go readFromSerer(conn)
	reader := bufio.NewReader(os.Stdin)
	cadena, _ := reader.ReadString('\n')

	for strings.Compare(cadena, exitCode) != 0 {
		_, _ = conn.Write(([]byte)(fmt.Sprintf("%s\n", cadena)))
		cadena, _ = reader.ReadString('\n')
	}
	_, _ = conn.Write(([]byte)(exitCode))
}

func readFromSerer(conn net.Conn) {
	var cadena string
	for true {
		cadena, _ = bufio.NewReader(conn).ReadString('\n')
		if len(cadena) > 0 {
			fmt.Print(cadena)
		}
	}
}
