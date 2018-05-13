package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	f,err:= os.Create("task.txt")
	if(err != nil){
		return
	}
	defer f.Close()

	fmt.Println("Tasks for today:")
	f.WriteString("Tasks for today:\r\n")
	for i := 0; i < 50; i++ {
		num:= rand.Uint32()
		fmt.Println(num)
		str := fmt.Sprint(num) + "\r\n"
		f.WriteString(str)
	}
	for i := 0; i < 50; i++ {
		num:= rand.Uint64()
		fmt.Println(num)
		str := fmt.Sprint(num)
		f.WriteString(str + "\r\n")
	}
}
