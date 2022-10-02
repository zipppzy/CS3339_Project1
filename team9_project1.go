package main

import (
	"os"
)

type Instruction struct {
	rawInstruction  string
	lineValue       uint64
	memLoc          uint64
	opcode          uint64
	op              string
	instructionType string
	rm              uint8
	shamt           uint8
	rn              uint8
	rd              uint8
	rt              uint8
	op2             uint8
	address         uint16
	immediate       int16
	offset          int32
	conditional     uint8
	shiftCode       uint8
	field           uint32
	breakIns        bool
}

var InstructionList []Instruction

func main() {
	//Checks if addtest1_bin.txt exists
	if _, err := os.Stat("addtest1_bin.txt"); err == nil {
		os.Args[0] = "addtest1_bin.txt"
	}

	//Inputs Command-Line
	ReadBinary(os.Args[0])

	ProcessInstructionList(InstructionList)
	//fmt.Println(InstructionList)
	WriteInstructions("out.txt", InstructionList)

}
