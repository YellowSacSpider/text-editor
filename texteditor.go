package main

// #include<conio.h>
import "C"

import (
	"fmt"
	"bufio"
	"os"
)

type StructArray struct{
	text [999]string
	ArrayIndex int
}

func TypeText(sa *StructArray){
		sa.ArrayIndex = 0

		
		file, err1 := os.Create("text.txt")
		if err1 != nil {
			return
		}
		
		scanner := bufio.NewScanner(os.Stdin)

		for fmt.Printf("%d| ", sa.ArrayIndex); scanner.Scan(); fmt.Printf("%d| ", sa.ArrayIndex) {
			sa.text[sa.ArrayIndex] = scanner.Text()
			file.WriteString(sa.text[sa.ArrayIndex] + "\n")
			sa.ArrayIndex++
		}
	
		if err := scanner.Err(); err != nil {
    		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		
}

func main() {
	sa := StructArray{}
	TypeText(&sa)
}