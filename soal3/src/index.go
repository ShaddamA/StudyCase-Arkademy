package main

import (
	"fmt"
	"strings"
)

func main()  {
	vocalKonsonant("ARKADEMY")
}

func vocalKonsonant(teks string){
	teks1 := strings.ToUpper(teks)// gunanya apabila input berupa huruf kecil mempermudah perkondisian
	teksArray := strings.Split(teks1, "")
	fmt.Println(teksArray)		
	for _, r := range teksArray{
		if (r == `A` || r == `I` || r == `U` || r == `E` || r == `O`){
			fmt.Println(r)
		}
	}
	for _, r := range teksArray{
		if (r == `A` || r == `I` || r == `U` || r == `E` || r == `O`){
			fmt.Print("")
		}else{
			fmt.Println(r)
		}
	}
}
