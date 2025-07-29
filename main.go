package main

import (
	"fmt"
	"log"
	"net/http"
	"reddis/reddis"
	"strings"
)

var data = reddis.CreateDataStorage()

func parsePathAndCommand(path string) {
	commands := strings.Split(path, "/")[1:]
	fmt.Println(commands[0])
	switch commands[0] {
	case "set":
		if len(commands) < 3 {
			log.Fatal("Not enough commands")
		}
		reddis.AddIntoStorage(data, commands[1], commands[2])
	case "get":
		if len(commands) < 2 {
			log.Fatal("Not enough commands")
		}
		log.Println(reddis.GetFromStorage(data, commands[1]))
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		parsePathAndCommand(path)
	})

	http.ListenAndServe(":8080", nil)
}
