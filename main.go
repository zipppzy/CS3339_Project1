package main

import "fmt"

type Instruction struct {
	rawInstruction string
}

var InstructionList []Instruction

func main() {
	fmt.Println("Hello World")
	ReadBinary("addtest1_bin.txt")
	fmt.Println(InstructionList)
}
