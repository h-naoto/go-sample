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
	cM = make(map[string]chan string)
	cM["A"] = make(chan string, 1)
	cM["A1"] = make(chan string, 1)
	cM["A11"] = make(chan string, 1)
	cM["A2"] = make(chan string, 1)
	cM["B"] = make(chan string, 1)
	cM["B1"] = make(chan string, 1)
	cM["B2"] = make(chan string, 1)
	ctxA, cancelA := context.WithCancel(context.Background())
	ctxA = context.WithValue(ctxA, "keyA", "valA")
	ctxA1, cancelA1 := context.WithCancel(ctxA)
	ctxA1 = context.WithValue(ctxA1, "keyA1", "valA1")
	ctxA11, cancelA11 := context.WithCancel(ctxA1)
	ctxA11 = context.WithValue(ctxA11, "keyA11", "valA11")
	ctxA2, cancelA2 := context.WithCancel(ctxA)
	ctxA2 = context.WithValue(ctxA2, "keyA2", "valA2")
	ctxB, cancelB := context.WithCancel(context.Background())
	ctxB = context.WithValue(ctxB, "keyB", "valB")
	ctxB1, cancelB1 := context.WithCancel(ctxB)
	ctxB1 = context.WithValue(ctxB1, "keyB1", "valB1")
	ctxB2, cancelB2 := context.WithCancel(ctxB)
	ctxB2 = context.WithValue(ctxB2, "keyB2", "valB2")
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
		fmt.Println("# What process do stop")
		fmt.Print("> ")
		sc.Scan()
		output := sc.Text()
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
		//time.Sleep(1000 * time.Millisecond)
	}
}

func A(cA chan string, ctxA context.Context) {
	for {
		select {
		case <-ctxA.Done():
			close(cM["A"])
			delete(cM, "A")
			//fmt.Println("dead")
			return
		case str, _ := <-cA:
			str1 := ctxA.Value("keyA")
			fmt.Println(str, str1)
		}
	}
}
func A1(cA1 chan string, ctxA1 context.Context) {
	for {
		select {
		case <-ctxA1.Done():
			close(cM["A1"])
			delete(cM, "A1")
			//fmt.Println("dead")
			return
		case str, _ := <-cA1:
			str1 := ctxA1.Value("keyA1")
			fmt.Println(str, str1)
		}
	}
}
func A11(cA11 chan string, ctxA11 context.Context) {
	for {
		select {
		case <-ctxA11.Done():
			close(cM["A11"])
			delete(cM, "A11")
			//fmt.Println("dead")
			return
		case str, _ := <-cA11:
			str1 := ctxA11.Value("keyA11")
			fmt.Println(str, str1)
		}
	}
}
func A2(cA2 chan string, ctxA2 context.Context) {
	for {
		select {
		case <-ctxA2.Done():
			close(cM["A2"])
			delete(cM, "A2")
			//fmt.Println("dead")
			return
		case str, _ := <-cA2:
			str1 := ctxA2.Value("keyA2")
			fmt.Println(str, str1)
		}
	}
}
func B(cB chan string, ctxB context.Context) {
	for {
		select {
		case <-ctxB.Done():
			close(cM["B"])
			delete(cM, "B")
			//fmt.Println("dead")
			return
		case str, _ := <-cB:
			str1 := ctxB.Value("keyB")
			fmt.Println(str, str1)
		}
	}
}
func B1(cB1 chan string, ctxB1 context.Context) {
	for {
		select {
		case <-ctxB1.Done():
			close(cM["B1"])
			delete(cM, "B1")
			//fmt.Println("dead")
			return
		case str, _ := <-cB1:
			str1 := ctxB1.Value("keyB1")
			fmt.Println(str, str1)
		}
	}
}
func B2(cB2 chan string, ctxB2 context.Context) {
	for {
		select {
		case <-ctxB2.Done():
			close(cM["B2"])
			delete(cM, "B2")
			//fmt.Println("dead")
			return
		case str, _ := <-cB2:
			str1 := ctxB2.Value("keyB2")
			fmt.Println(str, str1)
		}
	}
}
