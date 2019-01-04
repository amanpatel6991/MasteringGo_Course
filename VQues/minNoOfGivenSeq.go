package main

import "fmt"

func main() {

	str := "IIIIII"
	//str := "DDIDDIID"
	//str := "IIDDIIDD"
	str1 := []rune(str)
	//fmt.Println(len(str1))
	iterator := 0
	num := 0

	if str1[0] == 'I' {
		num = 1
		iterator ++
	}

	for iterator < len(str) {

		if str1[iterator] == 'I' {
			countI := 1
			i := iterator
			for {
				fmt.Println(i)
				if str1[i] == 'I' {
					countI++
				} else {
					iterator = i
					fmt.Println(iterator)
					break
				}
				i++
			}

			for j := 0; j < countI + 1; j++ {
				num = num * 10 + j
			}

		} else if str1[iterator] == 'D' {

			countD := 1
			i := iterator + 1
			for {
				if str1[i] == 'D' {
					countD++
				} else {
					iterator = i
					fmt.Println(iterator)
					break
				}
				i++
			}

			for j := countD + 1; j > 0; j-- {
				num = num * 10 + j
			}

		}

	}

	fmt.Println(num, iterator)

}
