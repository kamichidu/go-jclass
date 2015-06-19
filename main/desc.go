package main

import (
	"fmt"
	"github.com/kamichidu/go-jclass/parser"
)

func main() {
	descriptors := []string{
		"B",
		"C",
		"D",
		"F",
		"I",
		"J",
		"S",
		"Z",
		"Ljava/lang/Object;",
		"[B",
		"[C",
		"[D",
		"[F",
		"[I",
		"[J",
		"[S",
		"[Z",
		"[Ljava/lang/Object;",
		"[[B",
		"[[C",
		"[[D",
		"[[F",
		"[[I",
		"[[J",
		"[[S",
		"[[Z",
		"[[Ljava/lang/Object;",
	}
	for _, descriptor := range descriptors {
		fmt.Printf("Try to parse '%s'\n", descriptor)
		if ret, err := parser.ParseFieldDescriptor(descriptor); err == nil {
			fmt.Printf(" -> %s", ret.TypeName)
		} else {
			fmt.Printf(" -> error:\n%v", err)
		}
		fmt.Println()
	}
}
