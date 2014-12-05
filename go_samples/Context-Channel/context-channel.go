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
	cM["A1"] = make(chan string, 1)
	cM["A11"] = make(chan string, 1)
	cM["A2"] = make(chan string, 1)
	cM["B"] = make(chan string, 1)
	cM["B1"] = make(chan string, 1)
	cM["B2"] = make(chan string, 1)
	ctxA, cancelA := context.WithCancel(context.Background())
	ctxA1, cancelA1 := context.WithCancel(ctxA)
	ctxA11, cancelA11 := context.WithCancel(ctxA1)
	ctxA2, cancelA2 := context.WithCancel(ctxA)
	ctxB, cancelB := context.WithCancel(context.Background())
	ctxB1, cancelB1 := context.WithCancel(ctxB)
	ctxB2, cancelB2 := context.WithCancel(ctxB)

	go CommandUI(cCmd, cCmdR)
	go A(cM["A"], ctxA)
	go A1(cM["A1"], ctxA1)
	go A11(cM["A11"], ctxA11)
	go A2(cM["A2"], ctxA2)
	go B(cM["B"], ctxB)
	go B1(cM["B1"], ctxB1)
	go B2(cM["B2"], ctxB2)

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
				cancelA()
			}
		case "A1":
			_, ok := cM[output]
			if ok {
				cancelA1()
			}
		case "A11":
			_, ok := cM[output]
			if ok {
				cancelA11()
			}
		case "A2":
			_, ok := cM[output]
			if ok {
				cancelA2()
			}
		case "B":
			_, ok := cM[output]
			if ok {
				cancelB()
			}
		case "B1":
			_, ok := cM[output]
			if ok {
				cancelB1()
			}
		case "B2":
			_, ok := cM[output]
			if ok {
				cancelB2()
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

func A(cA chan string, ctxA context.Context) {
	for {
		select {
		case <-ctxA.Done():
			close(cM["A"])
			delete(cM, "A")
			return
		default:
			str, _ := <-cA
			fmt.Println(str)
		}
	}
}
func A1(cA1 chan string, ctxA1 context.Context) {
	for {
		select {
		case <-ctxA1.Done():
			close(cM["A1"])
			delete(cM, "A1")
			return
		default:
			str, _ := <-cA1
			fmt.Println(str)
		}
	}
}
func A11(A11 chan string, ctxA11 context.Context) {
	for {
		select {
		case <-ctxA11.Done():
			close(cM["A11"])
			delete(cM, "A11")
			return
		default:
			str, _ := <-A11
			fmt.Println(str)
		}
	}
}
func A2(cA2 chan string, ctxA2 context.Context) {
	for {
		select {
		case <-ctxA2.Done():
			close(cM["A2"])
			delete(cM, "A2")
			return
		default:
			str, _ := <-cA2
			fmt.Println(str)
		}
	}
}
func B(cB chan string, ctxB context.Context) {
	for {
		select {
		case <-ctxB.Done():
			close(cM["B"])
			delete(cM, "B")
			return
		default:
			str, _ := <-cB
			fmt.Println(str)
		}
	}
}
func B1(cB1 chan string, ctxB1 context.Context) {
	for {
		select {
		case <-ctxB1.Done():
			close(cM["B1"])
			delete(cM, "B1")
			return
		default:
			str, _ := <-cB1
			fmt.Println(str)
		}
	}
}
func B2(cB2 chan string, ctxB2 context.Context) {
	for {
		select {
		case <-ctxB2.Done():
			close(cM["B2"])
			delete(cM, "B2")
			return
		default:
			str, _ := <-cB2
			fmt.Println(str)
		}
	}
}
