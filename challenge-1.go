package main

import "fmt"

func main()  {
	var (
		i int = 21
		persen string = "%"
		j bool = true
		rusiancode string = "Я"
		base8  uint = 031
		base10  uint = 21
		base16 uint = 15
		float float32 = 123.456
	)

	fmt.Println(i)
	fmt.Printf("%T \n", i)
	fmt.Println(persen)
	fmt.Println(j)
	fmt.Println(rusiancode)
	fmt.Printf("%d\n", base10)
	fmt.Println(base8)
	fmt.Printf("%x\n", base16)
	fmt.Printf("%X\n", base16)
	fmt.Printf("%+q\n", rusiancode)
	fmt.Println(float)
	fmt.Printf("%e\n", float)
}