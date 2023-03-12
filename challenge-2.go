package main

import "fmt"

func main()  {
	i := 5
	j := 10

	for x := 0; x < i; x++ {
		fmt.Println("Nilai i = ", x)
	}

	for x := 0; x < j; x++ {
		if x < 5 {
			fmt.Println("Nilai j = ", x)
		}
	}

	for i, char := range "САШАРВО" {
		fmt.Printf("character %#U starts at byte position %d\n", char, i)
	}

	for x := 0; x <= j; x++ {
		if x > 5 {
			fmt.Println("Nilai j = ", x)
		}
	}
}