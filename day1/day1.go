package main

import(
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	)
func main(){
	readFile, err := os.Open("advent1.txt")
	if err != nil  {
		fmt.Println(err)
	}

	scn := bufio.NewScanner(readFile)
	scn.Split(bufio.ScanLines)
	var line string
	var sum int = 0
	var re = regexp.MustCompile("\\d")
	for scn.Scan(){
		line = scn.Text()
		matches := re.FindAllString(line, -1)
		firstInt, err := strconv.Atoi(matches[0])
		secondInt, err := strconv.Atoi(matches[len(matches)-1])
		if err != nil{
			fmt.Println("yell")
		}
		sum+=firstInt*10
		sum+=secondInt
	}
	fmt.Println(sum)
	readFile.Close()

}
