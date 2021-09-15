package main

import (
	"crypto/sha1"
	"fmt"
	"hal9000"
	"hal9000/types"
	"net"
	"os"
	"strings"
)

type InterfaceTypeSocket struct {
	Connection net.Conn
	Open       bool
}

func (i InterfaceTypeSocket) Type() string {
	return "socket"
}

func (i InterfaceTypeSocket) ID() string {
	h := sha1.New()
	h.Write([]byte(i.Connection.RemoteAddr().String()))
	bs := h.Sum(nil)
	return fmt.Sprintf("s-%x", bs)
}

func (i InterfaceTypeSocket) IsStillValid() bool {
	return i.Open
}

func (i InterfaceTypeSocket) SupportsVisuals() bool {
	return false
}

func (i InterfaceTypeSocket) SendMessage(m types.ResponseMessage) error {
	_, err := i.Connection.Write([]byte(m.Text))
	return err
}

func handleConnection(runtime types.Runtime, conn net.Conn) {
	defer conn.Close()

	readChannel := make(chan string)
	go (func() {
		lastBuff := []byte{}
		for {
			buff := make([]byte, 1024)
			n, err := conn.Read(buff)
			if err != nil {
				runtime.Logger().LogError(err)
			} else {
				newLine := -1
				for i, b := range buff[:n] {
					if b == '\n' || b == '\r' {
						newLine = i
						break
					}
				}
				if newLine < 0 {
					lastBuff = append(lastBuff, buff[:n]...)
				} else {
					readChannel <- strings.Trim(string(append(lastBuff, buff[:newLine]...)), " \n\r\t")
					lastBuff = buff[newLine:n]
				}
			}
		}
	})()

	userId := <-readChannel
	person, err := runtime.People().GetPersonByID(userId)
	if err != nil {
		runtime.Logger().LogError(err)
		return
	}

	iface := InterfaceTypeSocket{conn, true}
	runtime.InterfaceStore().Register(person, iface)
	ses := hal9000.NewSession(runtime, person, iface)

	for {
		input := <-readChannel
		halReq := types.RequestMessage{Message: input}

		response, err := hal9000.ProcessIncomingMessage(runtime, &ses, halReq)
		if err != nil {
			fmt.Println(err)
			iface.Open = false
			return
		}

		err = ses.Interface.SendMessage(response)
		if err != nil {
			fmt.Println(err)
			iface.Open = false
			return
		}

		runtime.SessionStore().SaveSession(ses)
		if err != nil {
			fmt.Println(err)
			iface.Open = false
			return
		}
	}
}

func startSocketServer(runtime types.Runtime) {
	ln, err := net.Listen("tcp", os.Getenv("SOCKET_SERVER"))
	if err != nil {
		runtime.Logger().LogError(err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			runtime.Logger().LogError(err)
		} else {
			go handleConnection(runtime, conn)
		}
	}
}