package main

import(
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
	"math"
	)
type three struct{
	destination int
	source int
	ranges int
}
func buildArr(seedArr []int, mp []three) []int {
	//fmt.Println(mp)
	resArr := make([]int, 0)
	for _, seed:= range seedArr{
		added := false
		for _, val:= range mp{
			if (val.source + val.ranges - 1) >= seed && val.source <= seed {
				newVal := (seed - val.source) + val.destination
				resArr = append(resArr, newVal)
				added = true
			}


		}
		if !added {
			resArr = append(resArr, seed)
		}
	}
	return resArr
}
func buildThree(re *regexp.Regexp, lines []string, currLine *int) []three {
	res := make([]three, 0)
	for lines[*currLine] != "" && *currLine != len(lines)-1 {
		findVals := re.FindAllString(lines[*currLine], -1)
		dest, err := strconv.Atoi(findVals[0])
		source, err := strconv.Atoi(findVals[1])
		ranges, err := strconv.Atoi(findVals[2])
		if err != nil {
			fmt.Println("Atoi fail")
		}
		var temp = three{dest, source, ranges}
		res = append(res, temp)
		*currLine++
	}
	*currLine+=2
	return res
}
func part1(lines []string){
	Patt := `\d+`
	re := regexp.MustCompile(Patt)

	findSeeds := re.FindAllString(lines[0], -1)
	seedArr := make([]int, 0)
	for _, val:= range findSeeds{
		seed, err := strconv.Atoi(val)
		if err != nil{
			fmt.Println("Atoi fail")
		}
		seedArr = append(seedArr, seed)
	}
	fmt.Println("Seeds")
	fmt.Println(seedArr)
	fmt.Println(len(seedArr))
	arrStruct := make([]three, 0)
	currLine := 0
	for !strings.Contains(lines[currLine], "seed-to-soil"){
		currLine++
	}

	fmt.Println("Build seed to soil")
	currLine++
	arrStruct = buildThree(re, lines, &currLine)
	seedToSoilArr := buildArr(seedArr, arrStruct)
	fmt.Println("SeedtoSoil")
	fmt.Println(seedToSoilArr)
	fmt.Println(len(seedToSoilArr))

	fmt.Println("Build soil to fert")
	arrStruct = buildThree(re, lines, &currLine)
	soilToFertArr := buildArr(seedToSoilArr, arrStruct)
	fmt.Println("SoiltoFert")
	fmt.Println(soilToFertArr)
	fmt.Println(len(soilToFertArr))

	fmt.Println("Build fert to water")
	arrStruct = buildThree(re, lines, &currLine)
	fertToWaterArr := buildArr(soilToFertArr, arrStruct)
	fmt.Println("fertToWaterArr")
	fmt.Println(fertToWaterArr)
	fmt.Println(len(fertToWaterArr))

	fmt.Println("Build water to light")
	arrStruct = buildThree(re, lines, &currLine)
	waterToLightArr := buildArr(fertToWaterArr, arrStruct)
	fmt.Println("waterToLightArr")
	fmt.Println(waterToLightArr)
	fmt.Println(len(waterToLightArr))

	fmt.Println("Build light to temp")
	arrStruct = buildThree(re, lines, &currLine)
	lightToTempArr := buildArr(waterToLightArr, arrStruct)
	fmt.Println("lightToTempArr")
	fmt.Println(lightToTempArr)
	fmt.Println(len(lightToTempArr))

	fmt.Println("Build temp to humidity")
	arrStruct = buildThree(re, lines, &currLine)
	tempToHumidArr := buildArr(lightToTempArr, arrStruct)
	fmt.Println("tempToHumidArr")
	fmt.Println(tempToHumidArr)
	fmt.Println(len(tempToHumidArr))

	fmt.Println("Build humid to location")
	arrStruct = buildThree(re, lines, &currLine)
	humidToLocArr := buildArr(tempToHumidArr, arrStruct)
	fmt.Println("humidToLocArr")
	fmt.Println(humidToLocArr)
	fmt.Println(len(humidToLocArr))

	min := math.MaxInt64 
	for i:=0; i < len(humidToLocArr); i++ {
		if humidToLocArr[i] < min{
			min = humidToLocArr[i]
		}
	}
	fmt.Println("MIN")
	fmt.Println(min)   
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
