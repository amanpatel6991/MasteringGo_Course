package main
import "fmt"
func main() {

	//var n int
	//fmt.Scanln(&n)
	//
	//var str string
	//fmt.Scanln(&str)

	 str := "011101011101"

	count0:=0
	count1:=0
	tempLen:=0
	largestLen :=0

	for _ , v:=range str {

		if v == 48 {
			count0++
		}
		if v==49{
			count1++
		}

		if count0>=count1{
			tempLen++
		}else{
			count0=0
			count1=0
			tempLen=0
		}

		if tempLen>=largestLen && count0!=count1{
			largestLen = tempLen

		}

	}
	// if count1 == 0 {
	//             largestLen++
	// }

	fmt.Print(largestLen)

}