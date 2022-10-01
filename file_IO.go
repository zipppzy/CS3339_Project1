package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// reads text file and makes Instructions and adds them to the InstructionList
func ReadBinary(filePath string) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var pc uint64
	pc = 96
	for scanner.Scan() {
		InstructionList = append(InstructionList, Instruction{rawInstruction: scanner.Text(), memLoc: pc})
		pc += 4
	}
}

func WriteInstructions(filePath string, list []Instruction) {
	f, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for i := 1; i < len(list); i++ {
		switch list[i].instructionType {
		case "B":

		case "I":

		case "CB":

		case "IM":

		case "D":
			//write binary with spaces
			_, err := fmt.Fprintf(f, "%s %s %s %s %s\t", list[i].rawInstruction[0:11], list[i].rawInstruction[11:20], list[i].rawInstruction[20:22], list[i].rawInstruction[22:27], list[i].rawInstruction[27:32])
			//write memLoc and opcode
			_, err = fmt.Fprintf(f, "%d\t%s\t", list[i].memLoc, list[i].op)
			//write operands
			_, err = fmt.Fprintf(f, "R%d, [R%d, #%d]\n", list[i].rt, list[i].rn, list[i].address)
			if err != nil {
				log.Fatal(err)
			}
		case "R":
			//write binary with spaces
			_, err := fmt.Fprintf(f, "%s %s %s %s %s\t", list[i].rawInstruction[0:11], list[i].rawInstruction[11:16], list[i].rawInstruction[16:22], list[i].rawInstruction[22:27], list[i].rawInstruction[27:32])
			//write memLoc and opcode
			_, err = fmt.Fprintf(f, "%d\t%s\t", list[i].memLoc, list[i].op)
			//write operands
			_, err = fmt.Fprintf(f, "R%d, R%d, ", list[i].rd, list[i].rn)
			if list[i].op == "LSL" || list[i].op == "ASR" || list[i].op == "LSR" {
				_, err = fmt.Fprintf(f, "#%d\n", list[i].shamt)
			} else {
				_, err = fmt.Fprintf(f, "R%d\n", list[i].rm)
			}
			if err != nil {
				log.Fatal(err)
			}

		}
	}
}
