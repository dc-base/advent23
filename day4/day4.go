package main

import(
	"bufio"
	"fmt"
	"os"
	"regexp"
	)
func part1(lines []string){
	var sum int = 0;
	Patt := `\d+`
	re := regexp.MustCompile(Patt)
	for _, line := range lines{
		var cardSum int = 0
		nums := re.FindAllString(line, -1)
		winNums := nums[1:11]
		haveNums := nums[11:len(nums)]
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
