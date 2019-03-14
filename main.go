package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
	"strings"
)

var Waitron string
var Listen string

/**
 * is the webpage handler and print the metrics of waitron
 */
func metrics(w http.ResponseWriter, r *http.Request){
	var message string

	resp := requestWaitron("list")
	var list []string
	err := json.Unmarshal(resp, &list)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(list); i++ {
		list[i] = strings.ReplaceAll(list[i], ".yaml", "")
	}

	resp = requestWaitron("health")
	type Health struct {
		State string
	}
	var health Health
	err = json.Unmarshal(resp, &health)
	if err != nil {
		panic(err)
	}
	message += "# TYPE waitron_health gauge\n"
	if (health.State) == "OK" {
		message += "waitron_health 1\n"
	} else {
		message += "waitron_health 0\n"
	}


	message += "# TYPE waitron_node_state gauge\n"
	for i := 0; i < len(list); i++ {
		if strings.EqualFold(string(requestWaitron("status/" + list[i])), "Installing") {
			message += "waitron_node_state{node=" + list[i] + "} 1\n"
		} else {
			message += "waitron_node_state{node=" + list[i] + "} 0\n"
		}
	}
	w.Write([]byte(message))
}

/**
 * request the waitron endpoint
 * @param path is the requested path
 * @return the body of the requested page
 */
func requestWaitron(path string) []byte {
	resp, err := http.Get(Waitron + "/" + path);
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}

/**
 * print the help text
 */
func help() {
	fmt.Println("Help dialog:\n")
	fmt.Println("Options:")
	fmt.Println("\thelp: Print this help dialog")
	fmt.Println("\tlisten: listen on this address")
	fmt.Println("\twaitron: address to the waitron server")
	fmt.Println("\nUsage:")
	fmt.Println("\twaitron_exporter listen=localhost:8080 waitron=localhost:8090")
	fmt.Println("\twaitron_exporter help")
}

func main() {
	args := os.Args[1:]
	if len(args) > 2 || len(args) == 0 {
		help()
		os.Exit(1)
	}

	for i := 0; i < len(args); i++ {
		if strings.EqualFold(args[i],"help") {
			help()
			os.Exit(0)
		} else if strings.HasPrefix(args[i], "listen="){
			com := strings.Split(args[i], "=")
			Listen = com[1]
		} else if strings.HasPrefix(args[i], "waitron="){
			com := strings.Split(args[i], "=")
			Waitron = com[1]
		} else {
			help()
			os.Exit(2)
		}
	}

	fmt.Println("started exporter on " + Listen)
	http.HandleFunc("/", metrics)
	err := http.ListenAndServe(Listen, nil)
	if err != nil {
		panic(err)
	}
}
