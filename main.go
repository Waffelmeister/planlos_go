package main

import (
	"log"
	"net"
	"net/http"
)

func createListener() (l net.Listener, close func()) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	return l, func() {
		_ = l.Close()
	}
}
func main() {
	l, close := createListener()
	defer close()
	http.Handle("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// handle like normal
	}))
	log.Println("listening at", l.Addr().(*net.TCPAddr).Port)
	http.Serve(l, nil)
}
