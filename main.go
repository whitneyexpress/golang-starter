package main

import (
    "fmt"
    "os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("ERROR!")
		os.Exit(1)
	}
	
	var timesToPrint int
	var err error
	timesToPrint, err = strconv.Atoi(os.Args[1])
	if err != nil || timesToPrint < 0 || timesToPrint > 10	{
		fmt.Println("ERROR!")
		os.Exit(1)
	}
	for i:=0; i < timesToPrint; i++{
		fmt.Println("careful-kangaroo")
	}
}
