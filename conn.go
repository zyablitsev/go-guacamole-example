package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"time"
)

// OpenGuacdConn try connect to guacd service and return *net.TCPConn when success
// also return an error if resolve tcp addr failed
func OpenGuacdConn(host string, port uint16) (*net.TCPConn, error) {
	proto := "tcp4"
	tcpAddr := (*net.TCPAddr)(nil)
	err := (error)(nil)
	if tcpAddr, err = net.ResolveTCPAddr(
		proto,
		fmt.Sprintf("%s:%d", host, port),
	); err != nil {
		return nil, err
	}

	conn := (*net.TCPConn)(nil)
	ticker := time.NewTicker(time.Second * 1)
	for {
		log.Println("trying connect to guacd…")
		if conn, err = net.DialTCP("tcp", nil, tcpAddr); err != nil {
			log.Println(err)
			<-ticker.C
			continue
		}
		log.Println("connection to guacd established: ok")
		ticker.Stop()
		break
	}

	return conn, nil
}

// CloseGuacdConn returns function that send 'disconnect' instruction
// to the guacd and close connection when invoked
func CloseGuacdConn(conn *net.TCPConn) func() {
	data := EncodeInstructions([]string{
		"disconnect",
	})
	return func() {
		defer conn.Close()
		log.Println("close guacd connection")
		if _, err := conn.Write(data); err != nil {
			log.Fatal(err)
		}
	}
}

// ConnReadWholeInstruction read data from connection till ';' instruction delimiter
// and return instruction as []byte slice
func ConnReadWholeInstruction(conn *net.TCPConn) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (error)(nil)
	for {
		b := [64]byte{}
		n := 0
		if n, err = conn.Read(b[:]); err != nil {
			return nil, err
		}
		buf.Write(b[:n])
		if b[n-1] == ';' {
			break
		}
	}
	return buf.Bytes(), nil
}
