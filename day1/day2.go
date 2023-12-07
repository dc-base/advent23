package main

import(
	"bufio"
	"fmt"
	"os"
	"regexp"
	//"strconv"
	)
var values = make(map[string]int)
func loadValues(){
	var arr = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var arr2  = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	//insert text mappings
	for i:=0; i<9; i++{
		values[arr[i]] = i+1;
	}
	//insert integer mappings
	for i:=0; i<9; i++{
		values[arr2[i]] = i+1;
	}
	fmt.Println(values)
}

func main(){
	readFile, err := os.Open("advent1.txt")
	loadValues()
	//readFile, err := os.Open("small.txt")
	if err != nil  {
		fmt.Println(err)
	}

	scn := bufio.NewScanner(readFile)
	scn.Split(bufio.ScanLines)
	var line string
	var sum int = 0
	//var re = regexp.MustCompile("\\d|one|two|three|four|five|six|seven|eight|nine|oneight|twone|threeight|fiveight|eightwo|eighthree|nineight")
	var reF = regexp.MustCompile(".*?(\\d|one|two|three|four|five|six|seven|eight|nine).*")
	var reS = regexp.MustCompile(".*(\\d|one|two|three|four|five|six|seven|eight|nine).*?$")
	for scn.Scan(){
		line = scn.Text()
		//fmt.Println(line)
		//matches := re.FindAllString(line, -1)
		//fmt.Println(matches)
		firstInt := values[reF.ReplaceAllString(line, "$1")]
		//values[matches[0]]

		secondInt := values[reS.ReplaceAllString(line, "$1")]
		//fmt.Println(firstInt)
		//fmt.Println(secondInt)
		
		//values[matches[len(matches) - 1]]
		total := firstInt*10 + secondInt
		//fmt.Println(total)
		sum+=total
		//firstInt, err := strconv.Atoi(matches[0])
		//secondInt, err := strconv.Atoi(matches[len(matches)-1])
		//if err != nil{
		//	fmt.Println("yell")
		//}
		//sum+=firstInt*10
		//sum+=secondInt
	}
	fmt.Println(sum)
	readFile.Close()

}
