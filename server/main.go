package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

const exitCode = "@#quit\n"


type cliente struct {
	clientSocket net.Conn
	idUsuario    int
}

func initializeClient(clientSocket net.Conn, clientsConnections *[] cliente) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go listeningClient(&clientSocket, clientsConnections, &waitGroup)
	waitGroup.Wait()
}

func listeningClient(clientSocket *net.Conn, clientsConnections *[] cliente, waitGroup *sync.WaitGroup) {
	cadena, _ := bufio.NewReader(*clientSocket).ReadString('\n')

	for strings.Compare(cadena, exitCode) != 0 {
		fmt.Print("Usuario -> ", (*clientSocket).RemoteAddr(), "	MSG -> ")
		fmt.Print(cadena)
		sendMessage(clientSocket, clientsConnections, cadena)
		cadena, _ = bufio.NewReader(*clientSocket).ReadString('\n')
	}
	fmt.Println("Cierro el socket")
	err := (*clientSocket).Close()
	if err != nil {
		fmt.Println("Error al cerrar el socket del cliente")
	}
	waitGroup.Done()
}

func sendMessage(clientSocket *net.Conn, clientsConnections *[]cliente, cadena string) {
	for _, v := range *clientsConnections {
		if (v.clientSocket.RemoteAddr() != (*clientSocket).RemoteAddr()) {
			v.clientSocket.Write(([]byte) (fmt.Sprintf("[%d] : %s",v.idUsuario,cadena)))
		}
	}
}

func main() {
	var clientsConnections []cliente
	i := 1

	fmt.Println("Iniciando el servidor")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error al iniciar el servidor", err)
		return
	}
	fmt.Println("- - - - - - - - - ")
	fmt.Println("Servidor iniciado al espera de conexiones")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error al establecer la conexion con el nuevo cliente", err)
		}
		clientsConnections = append(clientsConnections, cliente{conn,i})
		i++
		fmt.Println("Nueva Conexion", conn.RemoteAddr())
		go initializeClient(conn, &clientsConnections)
	}

}
