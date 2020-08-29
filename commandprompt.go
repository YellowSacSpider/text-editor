package main

// #include<conio.h>
import "C"

import (
	"fmt"
	//"bufio"
	"os/exec"
)


func main() {
	fmt.Printf("Welcome to the TEXT EDITOR CMD \n")
	fmt.Printf("Commands: \n")
	fmt.Printf("Click 'i' for text editor \n")
	fmt.Printf("Click 'r' for text reader \n")
	for	{
    	c := C.getch()
		switch char := c; char {
			case 'i':
				cmd := exec.Command("cmd", "/C start texteditor.exe")
				err := cmd.Start()
				if err != nil {
					fmt.Printf("There is someting wrong! I cant find texteditor.exe in main folder! \n")
				}
				fmt.Printf("The text editor is open. \n")
		}
	}
}