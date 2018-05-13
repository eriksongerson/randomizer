package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	f,err:= os.Create("task.txt")
	if(err != nil){
		return
	}
	defer f.Close()

	fmt.Println("Задание на сегодня:")
	f.WriteString("Задание на сегодня:\n")
	for i := 0; i < 50; i++ {
		num:= rand.Uint32()
		fmt.Println(num)
		str := fmt.Sprint(num)
		f.WriteString(str + "\n")
	}
	for i := 0; i < 50; i++ {
		num:= rand.Uint64()
		fmt.Println(num)
		str := fmt.Sprint(num)
		f.WriteString(str + "\n")
	}
}
