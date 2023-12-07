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
	//loadValues()
	//readFile, err := os.Open("small.txt")
	if err != nil  {
		fmt.Println(err)
	}

	scn := bufio.NewScanner(readFile)
	scn.Split(bufio.ScanLines)
	var line string
	var sum int = 0
	//var reS = regexp.MustCompile(".*(\\d|one|two|three|four|five|six|seven|eight|nine).*?$")
	//\s* matches 0-infinite whitespaces
	gamePatt := `Game\s*(\d+):`
	numPatt := `\b\d+\b`
	wordPatt := `\b([a-zA-Z]+)\b`
	var gameNum = regexp.MustCompile(gamePatt)
	var rgxnumber = regexp.MustCompile(numPatt)
	var rgxword = regexp.MustCompile(wordPatt)

	var reds = 12
	var greens = 13
	var blues = 14
	var colors = make(map[string]int)
	colors["red"] = reds
	colors["green"] = greens
	colors["blue"] = blues
	out:
	for scn.Scan(){
		line = scn.Text()
		fmt.Println(line)
		gameN := gameNum.FindStringSubmatch(line)
		gamenum, err := strconv.Atoi(gameN[1])
		if err != nil{
			fmt.Println("mald")
		}
		//fmt.Println(gameN)
		line = gameNum.ReplaceAllString(line, "")
		gamesets := strings.Split(line, ";")
		//fmt.Println("gameset")
		//fmt.Println(gamesets)
		//fmt.Println(len(gamesets))
		for _, game := range gamesets {
			pulls := strings.Split(game, ",")
			for _, indiv := range pulls{
				//fmt.Println(indiv)
				nums := rgxnumber.FindAllString(indiv, -1)
				words := rgxword.FindAllString(indiv, -1)
				//fmt.Println(nums)
				num, err := strconv.Atoi(nums[0])
				if err != nil{
					fmt.Println("mald")
				}
				//fmt.Println(words)
				if num > colors[words[0]]{
					fmt.Println("fail")
					continue out
				}
			}
		}
		sum+=gamenum

	}
	fmt.Println(sum)
	readFile.Close()

}
