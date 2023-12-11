package main

import(
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	)
type race struct {
	time int 
	record int
}
func findNumWays(val race) int {
	res := 0
	for i:=val.time; i > 0; i-- {
		speed:=i
		availTime := val.time - i
		distance:=availTime*speed
		if distance > val.record {
			res++
		}
	}
	return res
}
func part1(lines []string){
	Patt := `\d+`
	re := regexp.MustCompile(Patt)
	times := re.FindAllString(lines[0], -1)
	distances := re.FindAllString(lines[1], -1)
	races := make([]race, 0)
	for i, val:= range times {
		time, err := strconv.Atoi(val)
		distance, err := strconv.Atoi(distances[i])
		if err != nil {
			fmt.Println("atoi failure")
		}
		temp := race{time, distance}
		races = append(races, temp)
	}
	fmt.Println(races)
	mul := 1
	for _, val:= range races {
		mul*=findNumWays(val)
	}
	fmt.Println(mul)

	
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
	part1(lines)
}
