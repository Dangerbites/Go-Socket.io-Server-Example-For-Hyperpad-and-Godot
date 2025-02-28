package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/zishang520/socket.io/v2/socket"
)

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "localhost"
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "localhost"
}

func main() {
	ip := getLocalIP()
	fmt.Printf("Socket.IO server is running at %s:3000\n", ip)

	io := socket.NewServer(nil, nil)

	// Handle Socket.IO connections
	io.On("connection", func(clients ...any) {
		client := clients[0].(*socket.Socket)
		fmt.Println("A client connected:", client.Id())
		client.Send("chat message", "Welcome to the chat room!")

		// Handle custom events
		client.On("chat message", func(data ...any) {
			fmt.Println("Received message:", data[0])
			io.Emit("chat message", data[0]) // Broadcast the message to all clients
		})

		// Handle disconnection
		client.On("disconnect", func(...any) {
			fmt.Println("A client disconnected:", client.Id())
		})
	})

	// Serve the Socket.IO handler on an HTTP server
	http.Handle("/socket.io/", io.ServeHandler(nil))

	// Serve the HTML file at the root URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	go http.ListenAndServe(":3000", nil)

	exit := make(chan struct{})
	SignalC := make(chan os.Signal, 1)
	signal.Notify(SignalC, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-SignalC
		close(exit)
	}()

	<-exit
	io.Close(nil)
	os.Exit(0)
}
