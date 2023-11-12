package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		//Ini di skip
		if i == 2 {
			continue
		}

		//Ini di stop
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

}
