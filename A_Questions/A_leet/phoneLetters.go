package main

import (
	"fmt"
)

func main() {

	lettersMap := make(map[string][]string , 0)
	lettersMap["2"] = []string{"a","b","c"}
	lettersMap["3"] = []string{"d","e","f"}
	lettersMap["4"] = []string{"g","h","i"}
	lettersMap["5"] = []string{"j","k","l"}
	lettersMap["6"] = []string{"m","n","o"}
	lettersMap["7"] = []string{"p","q","r","s"}
	lettersMap["8"] = []string{"t","u","v","w"}
	lettersMap["9"] = []string{"x","y","z"}

	input := "23"
	//inputArr := make([]int,0)
	//inp , _ := strconv.Atoi(input)
	//
	//for inp != 0 {
	//	inputArr = append(inputArr , inp%10)
	//	inp
	//}
	ans := make([]string , 0)

	//strasa := strings.Split(input , ans)
	fmt.Println(input)

	findCombo("" , input , &ans , lettersMap)

	fmt.Println(ans)

}

func findCombo(combination ,nextDigit string , ans *[]string , letterMap map[string][]string) {

	if len(nextDigit) == 0 {
		*ans = append(*ans , combination)
	} else {
		digit := nextDigit[0:1]
		letters := letterMap[digit]
		for _,v:=range letters {
			findCombo(combination+v , nextDigit[1:],ans,letterMap)
		}
	}

}