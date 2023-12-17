package main

import(
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	)
func part1(lines []string) {
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
	var sum int = 0
	out:
	for _, line := range lines{
		//fmt.Println(line)
		gameN := gameNum.FindStringSubmatch(line)
		gamenum, err := strconv.Atoi(gameN[1])
		if err != nil{
			fmt.Println("atoi fail")
		}
		line = gameNum.ReplaceAllString(line, "")
		gamesets := strings.Split(line, ";")
		for _, game := range gamesets {
			pulls := strings.Split(game, ",")
			for _, indiv := range pulls{
				nums := rgxnumber.FindAllString(indiv, -1)
				words := rgxword.FindAllString(indiv, -1)
				num, err := strconv.Atoi(nums[0])
				if err != nil{
					fmt.Println("atoi fail")
				}
				if num > colors[words[0]]{
					fmt.Println("fail")
					continue out
				}
			}
		}
		sum+=gamenum
	}
	fmt.Println(sum)
}
func part2(lines []string) {
	gamePatt := `Game\s*(\d+):`
	numPatt := `\b\d+\b`
	wordPatt := `\b([a-zA-Z]+)\b`
	var reID = regexp.MustCompile(gamePatt)
	var rgxnumber = regexp.MustCompile(numPatt)
	var rgxword = regexp.MustCompile(wordPatt)

	var reds = 12
	var greens = 13
	var blues = 14
	var colors = make(map[string]int)
	colors["red"] = reds
	colors["green"] = greens
	colors["blue"] = blues
	var minColors = make(map[string]int)
	minColors["red"] = 1
	minColors["green"] = 1
	minColors["blue"] = 1

	var addCubes int = 0
	for _, line := range lines{
		line = reID.ReplaceAllString(line, "")
		gamesets := strings.Split(line, ";")
		for _, game := range gamesets {
			pulls := strings.Split(game, ",")
			for _, indiv := range pulls{
				nums := rgxnumber.FindAllString(indiv, -1)
				words := rgxword.FindAllString(indiv, -1)
				num, err := strconv.Atoi(nums[0])
				if err != nil{
					fmt.Println("atoi fail")
				}
				if num > minColors[words[0]]{
					minColors[words[0]] = num
				}
			}			
		}
		temp := minColors["red"] * minColors["green"] * minColors["blue"]
			fmt.Println(temp)
			addCubes+=temp
			minColors["red"] = 1
			minColors["green"] = 1
			minColors["blue"] = 1
	}
	fmt.Println(addCubes)
}

func main(){
	readFile, err := os.Open("data.txt")
	defer readFile.Close()
	if err != nil  {
		fmt.Println(err)
	}

	scn := bufio.NewScanner(readFile)
	scn.Split(bufio.ScanLines)
	var lines []string


	for scn.Scan(){
		lines = append(lines, scn.Text())
	}
	//part1(lines)
	part2(lines)

}
