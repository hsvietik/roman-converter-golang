package main

import (
	"fmt"
	"regexp"
)

func main(){
	fmt.Println("Let's convert Roman numerals (e.g. VII) to integer value")
	for{
		var result uint = 0;
	var romanInt string;
	var romanArr []string;
	romanToInt := map[string]uint{"M": 1000,"D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1}
	romanValidation := regexp.MustCompile(`^(M{0,3})(D?C{0,3}|C[DM])(L?X{0,3}|X[LC])(V?I{0,3}|I[VX])$`)
	
	fmt.Println("Enter the Roman numeral:")
	fmt.Scan(&romanInt)

	if romanValidation.MatchString(romanInt){
		for _, letter := range romanInt{romanArr = append(romanArr, string(letter))}
		for i:=0; i<len(romanArr)-1;i++{
			curr:=romanToInt[romanArr[i]]
			next:= romanToInt[romanArr[i+1]]
			if curr < next {
				result -= curr;
			} else {
				result += curr;
			}
		}
		result+=romanToInt[romanArr[len(romanArr)-1]]
		fmt.Printf("You entered %v Roman numeral that is %v \n", romanInt, result)
		} else{fmt.Printf("The numeral %v is not a Roman numeral. \n", romanInt)}	
	}
}