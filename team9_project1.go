package main

import "fmt"

type Instruction struct {
	rawInstruction string
	lineValue      uint64
	opcode         uint64
}

var InstructionList []Instruction

func main() {
	ReadBinary("addtest1_bin.txt")
	//fmt.Println(InstructionList)
	ProcessInstructionList(InstructionList)
	fmt.Println(InstructionList)
}
