package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

const exitCode  = "@#quit\n"


func initializeClient (clientSocket net.Conn, clientsConnections * [] net.Conn){
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go listeningClient(&clientSocket, clientsConnections, &waitGroup)
	waitGroup.Wait()
}

func listeningClient(clientSocket * net.Conn,  clientsConnections * [] net.Conn, waitGroup * sync.WaitGroup) {
	cadena, _ := bufio.NewReader(*clientSocket).ReadString('\n')
	for strings.Compare(cadena, exitCode) != 0 {
		cadena, _ = bufio.NewReader(*clientSocket).ReadString('\n')
		sendMessage(clientSocket,clientsConnections,cadena)
	}
	fmt.Println("Cierro el socket")
	err := (*clientSocket).Close()
	if err != nil {
		fmt.Println("Error al cerrar el socket del cliente")
	}
	waitGroup.Done()
}

func sendMessage(clientSocket *net.Conn, clientsConnections *[]net.Conn, cadena string) {
	for _, v := range *clientsConnections  {
		
	}

}

func main() {
	var clientsConnections []net.Conn

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
		clientsConnections = append(clientsConnections,conn)
		go initializeClient(conn, &clientsConnections)
	}

}
