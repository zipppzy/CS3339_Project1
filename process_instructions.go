package main

import (
	"fmt"
	"strconv"
)

func ProcessInstructionList(list []Instruction) {
	for i := 0; i < len(list); i++ {
		//use processing functions on instructions here
		translateToInt(&list[i])
		opcodeMasking(&list[i])
	}
}

func translateToInt(ins *Instruction) {
	i, err := strconv.ParseUint(ins.rawInstruction, 2, 64)
	if err == nil {
		ins.lineValue = i
	} else {
		fmt.Println(err)
	}
}

// 4292870144 mask for opcode(first 11 bits)
func opcodeMasking(ins *Instruction) {
	ins.opcode = (ins.lineValue & 4292870144) >> 21
}
