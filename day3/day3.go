package main

import(
	"bufio"
	"fmt"
	"os"
	"unicode"
	"regexp"
	"strconv"
	//"strings"
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
// symbols := make(map[byte]bool{
//     "*": true,
//     "+":   true,
//     "@": true,
//     "$": true,
// 	"/": true,
// 	"-": true,
// 	"%": true,
// 	"&": true,
// 	"#": true,
// 	"=": true,
// })
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
	decPatt := `\d+`
	symbolPatt := `[*+@$\-%&#=\/]`
	reDec := regexp.MustCompile(decPatt)
	reSym := regexp.MustCompile(symbolPatt)
	var sum int = 0
	for i, line := range lines {
		values := reDec.FindAllStringIndex(line, -1)
		symbols := reSym.FindAllStringIndex(line, -1)
		symMap := make(map[int]bool)
		for _, symbol := range symbols{
			symMap[symbol[0]] = true
			//fmt.Println("sym")
			//fmt.Println(symbol)

		}
		fmt.Println(symMap)
		fmt.Println(values)
		for _, value := range values {
			dec := line[value[0]:value[1]]
			//fmt.Println(dec)
			_ = i

			if value[0] != 0 {
				if symMap[value[0]-1] {
					fmt.Println(dec)
					fmt.Println("Left")
					temp, err := strconv.Atoi(dec)
					if err != nil {
						fmt.Println(err)
					}
					sum+=temp
					continue
				}
			}
			if(value[1] != len(line)-1){
				if symMap[value[1]] {
					fmt.Println(dec)
					fmt.Println("Right")
					temp, err := strconv.Atoi(dec)
					if err != nil {
						fmt.Println(err)
					}
					sum+=temp
					continue
				}
			}
			//first line of file
			if i == 0{
				start := value[0] - 1
				if start < 0 {
					start = 0
				}
				end := value[1] + 1
				if end >= len(line) {
					end = len(line) - 1
				}
				nextLine := makeLineMap(lines[i+1])
				for start < end {
					if nextLine[start]{
						fmt.Println(dec)
						fmt.Println("FL adj")
						temp, err := strconv.Atoi(dec)
						if err != nil {
							fmt.Println(err)
						}
						sum+=temp
					}
					start++
				}
				continue
			} else if i == len(lines) - 1{ //handle last line
				start := value[0] - 1
				if start < 0 {
					start = 0
				}
				end := value[1] + 1
				if end >= len(line) {
					end = len(line) - 1
				}
				prevLine := makeLineMap(lines[len(lines)-2])
				for start < end {
					if prevLine[start]{
						fmt.Println(dec)
						fmt.Println("LL adj")
						temp, err := strconv.Atoi(dec)
						if err != nil {
							fmt.Println(err)
						}
						sum+=temp
					}
					start++
				}
				continue
			} else { //default case 
				start := value[0] - 1
				if start < 0 {
					start = 0
				}
				end := value[1] + 1
				if end >= len(line) {
					end = len(line) - 1
				}
				nextLine := makeLineMap(lines[i+1])
				prevLine := makeLineMap(lines[i-1]) 
				for start < end {
					if nextLine[start] || prevLine[start]{
						fmt.Println(dec)
						fmt.Println("MID adj")
						temp, err := strconv.Atoi(dec)
						if err != nil {
							fmt.Println(err)
						}
						sum+=temp
					}
					start++
				}
				continue
			}

			//default case

			
			//start := value[0]
			//for start < value[1] {
				
			//}
		}


	}
	fmt.Println("Sum=%d", sum)

	
}
func main(){
	//readFile, err := os.Open("small.txt")
	file, err := os.Open("data2.txt")
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