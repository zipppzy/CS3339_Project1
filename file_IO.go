package main

import (
	"bufio"
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
	for scanner.Scan() {
		InstructionList = append(InstructionList, Instruction{rawInstruction: scanner.Text()})
	}
}
