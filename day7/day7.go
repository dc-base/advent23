package main

import(
	"bufio"
	"fmt"
	"os"
	//"regexp"
	"strconv"
	"strings"
	)
const (
	HighCard int = iota
	OnePair     
	TwoPair     
	ThreeOfAKind 
	FullHouse
	FourOfAKind
	FiveOfAKind    
	)
type hand struct {
	cards string
	bid int
}
var valMap = map[rune]int {
	'2' : 2,
	'3' : 3,
	'4' : 4,
	'5' : 5,
	'6' : 6,
	'7' : 7,
	'8' : 8,
	'9' : 9,
	'T' : 10,
	'J' : 11,
	'Q' : 12,
	'K' : 13,
	'A' : 14,
}
func getHandStatus(s string) int{

}
func part1(lines []string){
	data := make([]hand, 0)
	for _, line := range lines{
		arr := strings.Split(line, " ")
		bid, err := strconv.Atoi(arr[1])
		if err != nil {
			fmt.Println("atoi fail")
		}
		temp := hand{arr[0], bid}
		data = append(data, temp)
	}
	fmt.Println(data)

	
}
func main(){
	file, err := os.Open("input.txt")
	if err != nil  {
		fmt.Println(err)
	}
	defer file.Close()
    	var lines []string
    	scanner := bufio.NewScanner(file)
    	for scanner.Scan() {
    	    lines = append(lines, scanner.Text())
    	}
	part1(lines[0:5])
	//part2()
}
