package main

/*
#include "gettid.h"
*/
import "C"

import "fmt"

func main() {
	ch := make(chan bool)
	for i:=0; i<10; i++ {
	go func() {
		f()
		ch <- true
	}()
	}
	f()
	for i:=0; i<10; i++ {
	<-ch
	}
}

func f() {
	for i := 0; i < 10; i++ {
		fmt.Println(C.cgettid())
	}
}
