package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"../../pkg/filedump"
	"../../pkg/httproxy"
)

const (
	addr     = "127.0.0.1:8080"
	filename = "wstalker.csv"
)

func main() {
	// Create Proxy
	log.Println("Creating HTTP Proxy")
	h, e1 := httproxy.NewHttProxy()
	if e1 != nil {
		log.Fatal(e1)
	}

	// Start Proxy
	log.Println("Starting in " + addr)
	e1 = h.StartBackground(addr)
	if e1 != nil {
		log.Fatal(e1)
	}

	// Manage CTRL+C
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		fmt.Print("\r")
		e := h.StopBackground()
		if e != nil {
			log.Fatal(e)
		}
		log.Println("Stopped HTTP Proxy")
	}()

	// Prepare FileDump
	log.Println("Saving Request in " + filename)
	f, e2 := filedump.NewFileDump(filename)
	if e2 != nil {
		log.Fatal(e2)
	}
	defer f.Close()

	// Read until Stop
	log.Println("Stalking Connections...")
	for {
		method, url, request, response, e := h.Read()
		if e != nil {
			break
		}

		// Print Method and URL
		log.Println(method + " - " + url)

		// Write into file
		e = f.Write(method, url, request, response)
		if e != nil {
			log.Fatal(e)
		}
	}
	log.Println("Closing wstalker")
}
