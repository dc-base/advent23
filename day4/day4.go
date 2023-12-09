package main

import(
	"bufio"
	"fmt"
	"os"
	"unicode"
	"regexp"
	//"strconv"
	//"strings"
	)

func isSymbol(val rune) bool {
	if(val != '.' && unicode.IsSymbol(val)){
		return true
	}
	return false
}
func makeLineMap(line string) map[int]bool {
	symbolPatt := `[*+@$\-%&#=\/]`
	//symbolPatt := `[^\w\s.:w]`
	reSym := regexp.MustCompile(symbolPatt)
	symbols := reSym.FindAllStringIndex(line, -1)
	symMap := make(map[int]bool)
	for _, symbol := range symbols{
		symMap[symbol[0]] = true
	}
	return symMap
}
func part1(lines []string){
	var sum int = 0;
	winPatt := `\d+`
	havePatt := `\d+`
	reWin := regexp.MustCompile(winPatt)
	reHave := regexp.MustCompile(havePatt)
	for _, line := range lines{
		var cardSum int = 0
		winNums := reWin.FindAllString(line, -1)
		haveNums := reHave.FindAllString(line, -1)
		winNums = winNums[1:11]
		haveNums = haveNums[11:len(haveNums)]
		fmt.Println(winNums)
		fmt.Println(haveNums)
		
		winMap := make(map[string]bool)
		for _, val :=range winNums {
			winMap[val] = true
		}

		for _, val:=range haveNums{
			if winMap[val]{
				if cardSum == 0{
					cardSum = 1;
				} else {
					cardSum *=2
				}
			}
		}
		sum+=cardSum


	}
	fmt.Println(sum)
	
}
func main(){
	//readFile, err := os.Open("small.txt")
	file, err := os.Open("data.txt")
	if err != nil  {
		fmt.Println(err)
	}
	defer file.Close()
    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
	//part1(lines[0:3])
	//part1(lines[1:6])
	part1(lines)
}