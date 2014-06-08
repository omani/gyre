package main

import (
	"github.com/armen/gyre"

	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	input = make(chan string)
	name  = flag.String("name", "Gyreman", "Your name or nick name in the chat session")
)

func chat() {
	node, err := gyre.New()
	if err != nil {
		log.Fatalln(err)
	}
	//defer node.Disconnect()
	node.Start()
	node.Join("CHAT")

	for {
		select {
		case e := <-node.Events():
			switch e.Type() {
			case gyre.EventShout:
				fmt.Printf("\r%s\n%s> ", string(e.Msg()), *name)
			}
		case msg := <-input:
			node.Shout("CHAT", []byte(msg))
		}
	}
}

func main() {

	flag.Parse()

	go chat()

	fmt.Printf("%s> ", *name)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input <- fmt.Sprintf("%s: %s", *name, scanner.Text())
		fmt.Printf("%s> ", *name)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln("reading standard input:", err)
	}
}