package main

import (
	"fmt"
	"time"
)

// Define the server type with quit & msg channel
type Server struct {
	quitch chan struct{}
	msgch  chan string
}

// Return a Pointer of the Server with func init the variables
func newServer() *Server {
	return &Server{
    // define the quit channel with empty struct to have 0 mem
    // since we just close the channel on quitting
		quitch: make(chan struct{}),
		msgch:  make(chan string),
	}
}

func (s *Server) start() {
	fmt.Println("Server is starting...")
	s.loop()
}

func (s *Server) loop() {
	fmt.Println("looping started...")

  mainloop:
	for {
		select {
		case <-s.quitch:
			fmt.Println("The quitch triggered...")
      break mainloop

		case msg := <-s.msgch:
			fmt.Println("The msg channel is triggered...", msg)

		default:
		}
	}
}

func main() {
	server := newServer()

	// an async self invoked lambda func that triggers after the 5 secs
	go func() {
		time.Sleep(time.Second * 4)
		server.msgch <- "poda pota is the new trend"

		time.Sleep(time.Second * 3)
		close(server.quitch)
	}()

	// make this async by running it as goroutine
	// go server.start()

	// the below is the blocking code
	// blocks the server and runs infinite loop on the server
	server.start()
  fmt.Println("the server is gracefully shutdonw")
}
