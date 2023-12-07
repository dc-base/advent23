package main

import(
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	)
//var values = make(map[string]int)
//func loadValues(){
//	var arr = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
//	var arr2  = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
//	//insert text mappings
//	for i:=0; i<9; i++{
//		values[arr[i]] = i+1;
//	}
//	//insert integer mappings
//	for i:=0; i<9; i++{
//		values[arr2[i]] = i+1;
//	}
//	fmt.Println(values)
//}
func main(){
	//readFile, err := os.Open("small.txt")
	readFile, err := os.Open("data.txt")
	if err != nil  {
		fmt.Println(err)
	}

	scn := bufio.NewScanner(readFile)
	scn.Split(bufio.ScanLines)
	var line string
	var sum int = 0
	gamePatt := `Game\s*(\d+):`
	numPatt := `\b\d+\b`
	wordPatt := `\b([a-zA-Z]+)\b`
	var gameNum = regexp.MustCompile(gamePatt)
	var rgxnumber = regexp.MustCompile(numPatt)
	var rgxword = regexp.MustCompile(wordPatt)
	for scn.Scan(){
		line = scn.Text()
		fmt.Println(line)
		
		

	}
	fmt.Println(sum)
	readFile.Close()
}