package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/context"
	"os"
	"time"
)

var sc = bufio.NewScanner(os.Stdin)
var cM map[string]chan string

func main() {
	// init channels
	cCmd := make(chan string, 1)
	cCmdR := make(chan string, 1)
	cM = make(map[string]chan string)
	cM["A"] = make(chan string, 1)
	cM["Aa"] = make(chan string, 1)
	cM["Ab"] = make(chan string, 1)
	cM["B"] = make(chan string, 1)
	cM["Ba"] = make(chan string, 1)
	cM["Bb"] = make(chan string, 1)

	go CommandUI(cCmd, cCmdR)
	go A(cM["A"])
	go Aa(cM["Aa"])
	go Ab(cM["Ab"])
	go B(cM["B"])
	go Ba(cM["Ba"])
	go Bb(cM["Bb"])

	for {
		for key, value := range cM {
			str := fmt.Sprintf("Alive < %s > thread", key)
			value <- str
		}
		time.Sleep(500 * time.Millisecond)
		cCmdR <- "Next"
		fmt.Println("# What process do stop")
		output := <-cCmd
		switch output {
		case "A":
			_, ok := cM[output]
			if ok {
				close(cM["A"])
			}
		case "Aa":
			_, ok := cM[output]
			if ok {
				close(cM["Aa"])
			}
		case "Ab":
			_, ok := cM[output]
			if ok {
				close(cM["Ab"])
			}
		case "B":
			_, ok := cM[output]
			if ok {
				close(cM["B"])
			}
		case "Ba":
			_, ok := cM[output]
			if ok {
				close(cM["Ba"])
			}
		case "Bb":
			_, ok := cM[output]
			if ok {
				close(cM["Bb"])
			}
		case "exit":
			fmt.Println("process finished !!!")
			os.Exit(0)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func CommandUI(cCmd, cCmdR chan string) {
	for {
		<-cCmdR
		fmt.Print("> ")
		if sc.Scan() {
			cCmd <- sc.Text()
		}
	}
}

func A(cA chan string) {
	for {
		str, ok := <-cA
		if ok {
			fmt.Println(str)
		} else {
			delete(cM, "A")
			return
		}
	}
}
func Aa(cAa chan string) {
	for {
		str, ok := <-cAa
		if ok {
			fmt.Println(str)
		} else {
			delete(cM, "Aa")
			return
		}
	}
}
func Ab(cAb chan string) {
	for {
		str, ok := <-cAb
		if ok {
			fmt.Println(str)
		} else {
			delete(cM, "Ab")
			return
		}
	}
}
func B(cB chan string) {
	for {
		str, ok := <-cB
		if ok {
			fmt.Println(str)
		} else {
			delete(cM, "B")
			return
		}
	}
}
func Ba(cBa chan string) {
	for {
		str, ok := <-cBa
		if ok {
			fmt.Println(str)
		} else {
			delete(cM, "Ba")
			return
		}
	}
}
func Bb(cBb chan string) {
	for {
		str, ok := <-cBb
		if ok {
			fmt.Println(str)
		} else {
			delete(cM, "Bb")
			return
		}
	}
}
