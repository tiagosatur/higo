package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func helloWorld() {
	fmt.Println("Hellow budy")
}

func serveHTMLFileThroughRequest() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Web services are easy with Go!")
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./home.html")
	})

	http.ListenAndServe(":3000", nil)
}

func createCliApplication() {
	level := flag.String("level", "CRITICAL", "log level to filter for") // 1: name of the parameter, 2: default value, 3: help message
	flag.Parse()                                                         // Look those line parameters and populate the variables

	file, err := os.Open("./log.txt") // Open log

	// Elevate the error management
	if err != nil {
		log.Fatal(err)
	}

	// release the file in the operating system so it can run efficiently
	// defers the execution until the main function exit
	defer file.Close()

	bufReader := bufio.NewReader(file)

	// line, err := bufReader.ReadString('\n'); - loop
	// err == nil - test making sure we don't have an error
	// line, err = bufReader.ReadString('\n') - Every time we go into the loop we gonna initialize the statement again
	for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
		// does that line contains that log level?
		// *level is a pointer
		if strings.Contains(line, *level) {
			fmt.Println((line))
		}
	}
}

func main() {
	createCliApplication()
}
