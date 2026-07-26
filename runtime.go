package main

import "fmt"

type Runtime struct {
}

func NewRuntime() *Runtime {
	return &Runtime{}
}

func (r *Runtime) Run(){
	fmt.Println("Runtime started")
}
