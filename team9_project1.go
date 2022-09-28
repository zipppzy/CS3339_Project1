package main

import "fmt"

type Instruction struct {
	rawInstruction string
	lineValue      uint64
}

var InstructionList []Instruction

func main() {
	fmt.Println("Hello World")
	ReadBinary("addtest1_bin.txt")
	fmt.Println(InstructionList)
	ProcessInstructionList(InstructionList)
	fmt.Println(InstructionList)
}
