package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
)

type Client struct {
	IPAddress string
}

var Clients []Client
var Requests []string

func listenf() {
	ln, err := net.Listen("tcp", _server)
	if err != nil {
		fmt.Println(Fore["RED"]+"[-]", err,Fore["RESET"])
		os.Exit(0)
	}
	defer ln.Close()
	// Accept incoming connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		// Handle connection in a separate goroutine
		go handleConnection(conn)
	}
}

func rest() {
	// stop = false
	v_client = ""
}

var canexit bool = true
func handleConnection(conn net.Conn) {
	defer conn.Close()
	defer rest()

	var addr string = strings.Split(conn.RemoteAddr().String(), ":")[0]

	if !contains(addr, Requests) { 
		Clients = append(Clients, Client{IPAddress: addr})
		Requests = append(Requests, addr)
	}

	if v_client != "" {
		if strings.Split(conn.RemoteAddr().String(), ":")[0] == v_client || v_client=="*" {
			stop = true

			fmt.Println(Fore["GREEN"]+"[+] Connection from "+Fore["BLUE"]+addr+Fore["RESET"])

			canexit = false
			var wg sync.WaitGroup
			wg.Add(2)
			tcpconn := conn.(*net.TCPConn)
			go func() {
				io.Copy(tcpconn, os.Stdin)
				fmt.Fprintf(os.Stderr, Fore["RED"]+"[-] Connection Interrupt press ENTER to continue"+Fore["RESET"])
				tcpconn.CloseWrite()
				tcpconn.CloseRead()
				wg.Done()
			}()
			go func() {
				io.Copy(os.Stdout, tcpconn)
				fmt.Fprintf(os.Stderr, Fore["RED"]+"[-] Connection Interrupt press ENTER to continue"+Fore["RESET"])
				tcpconn.CloseRead()
				tcpconn.CloseWrite()
				wg.Done()
			}()
			wg.Wait()
			canexit = true
			fmt.Printf("\n")
		}
	}

	rest()
	conn.Close()
}
