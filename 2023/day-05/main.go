package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
  currentDir, _ := os.Getwd()
  inputFilePath := ""
  if (len(os.Args) > 1 && os.Args[1] == "test") {
    inputFilePath = fmt.Sprintf("%s/2023/day-05/test-input.txt", currentDir)
  } else {
    inputFilePath = fmt.Sprintf("%s/2023/day-05/input.txt", currentDir)
  }
  
  inputString, _ := os.ReadFile(inputFilePath)

  mapParts := strings.Split(string(inputString), "\n\n")

  seeds :=  strings.Split(strings.Split(mapParts[0], ": ")[1], " ")

  fmt.Println(seeds)
  minNum := math.MaxInt

  for _, seedStr := range seeds {
    seed, _ := strconv.Atoi(seedStr)
    locationNumber := seed
    for _, mapDesc := range mapParts[1:] {
      mapRanges := strings.Split(mapDesc, "\n")[1:]
      fmt.Println("map: ", strings.Split(mapDesc, "\n")[0])
      for _, mapRange := range mapRanges {
        if len(mapRange) == 0 {
          continue
        }
        numbers := strings.Split(mapRange, " ")
        fmt.Println("numbers", numbers)
        destination, _ := strconv.Atoi(numbers[0])
        source, _ := strconv.Atoi(numbers[1])
        rangeVal, _ := strconv.Atoi(numbers[2])

        if locationNumber > source && locationNumber - source <= rangeVal {
          fmt.Println("in here")
          locationNumber = destination + (locationNumber - source) 
        }
        fmt.Println("location number: ", locationNumber)
        fmt.Println("seed: ", seed)
        fmt.Println("destination: ", destination)
        fmt.Println("source: ", source)
        fmt.Println("rangeVal: ", rangeVal)
        fmt.Println("dif: ", seed - source)
      }
      fmt.Println("----------------")
    }

    if (minNum > locationNumber) {
      minNum = locationNumber
    }
  }

  fmt.Println("Min number: ", minNum)

}
