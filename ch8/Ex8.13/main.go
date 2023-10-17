package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type client chan<- string // an outgoing message channel
type chatMessage struct {
	message string
	ch      client
}

var (
	entering  = make(chan client)
	leaving   = make(chan client)
	usernames = make(map[string]bool)
	messages  = make(chan chatMessage) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]string) // all connected clients
	for {
		select {
		case msg := <-messages:
			if strings.HasPrefix(msg.message, "[register]:") {
				username := strings.TrimPrefix(msg.message, "[register]:")
				fmt.Printf("[%s]", username)
				if _, ok := usernames[username]; ok {
					// inform the user that the username has already been taken
					msg.ch <- "[sign_up_error]"
					continue
				}
				usernames[username] = true
				users := "[sign_up_success]:Current users: "
				for username_db := range usernames {
					if username_db == username {
						users += username_db + " (you), "
					} else {
						users += username_db + ", "
					}
				}
				users = strings.TrimSuffix(users, ", ")
				// sending current users to the new user
				msg.ch <- users
				// registering the username to the clients map
				clients[msg.ch] = username
				fmt.Printf("A new user %q has registered\n", username)
				continue
			}
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg.message
			}
		case cli := <-entering:
			clients[cli] = ""
			fmt.Printf("A new client has connected\n")
			fmt.Printf("Current clients: %v\n", len(clients))
			cli <- "Please enter your name:"
		case cli := <-leaving:
			for client := range clients {
				if client == cli {
					continue
				}
				client <- clients[cli] + " has left"
			}
			fmt.Printf("%s\n", clients[cli])
			// first you delete the username
			delete(usernames, clients[cli])
			// and then its channel
			delete(clients, cli)
			close(cli)
			fmt.Printf("A client has disconnected\n")
			fmt.Printf("Current clients: %v\n", len(clients))
		}
	}
}

func clientWriter(conn net.Conn, registered *bool, ch <-chan string) {
	for msg := range ch {
		if strings.HasPrefix(msg, "[sign_up_success]:") {
			*registered = true
			fmt.Fprintln(conn, strings.TrimPrefix(msg, "[sign_up_success]:"))
			continue
		}
		if msg == "[sign_up_error]" {
			*registered = false
			fmt.Fprintln(conn, "This username has already been taken, please register other")
			continue
		}
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	registered := false
	go clientWriter(conn, &registered, ch)
	who := conn.RemoteAddr().String()

	messages <- chatMessage{
		message: who + " has arrived",
		ch:      ch,
	}
	entering <- ch

	timeout := time.NewTimer(5 * time.Minute)
	go func() {
		<-timeout.C
		timeout.Stop()
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		timeout.Reset(5 * time.Minute)
		if !registered {
			messages <- chatMessage{
				message: "[register]:" + input.Text(),
				ch:      ch,
			}
			continue
		}
		messages <- chatMessage{
			message: who + ": " + input.Text(),
			ch:      ch,
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	leaving <- ch
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
