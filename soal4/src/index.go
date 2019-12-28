package main

import (
	"strings"
	"fmt"
)

func main()  {
	print("48172a94")
}

func print(a string){
	angka := [10]string{"0","1","2","3","4","5","6","7","8","9"}

	b := strings.Split(a,"")
	count := len(b)
	for i:= 0; i < count; i++{
		for j := 0; j < 10; j++ {
			if (b[i] == angka[j]) {
				fmt.Print(b[i])
			}	
		}	
	}
}
