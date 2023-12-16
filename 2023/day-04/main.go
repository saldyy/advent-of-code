package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getTotalWinningCount(drawDescription string) int {
  total := 1
  drawDescriptions := strings.Split(drawDescription, "|")

  winningDraws := strings.Split(strings.Trim(drawDescriptions[0], " "), " ")
  draws := strings.Split(strings.Trim(drawDescriptions[1], " "), " ")

  // fmt.Println(winningDraws, "winningDraws")
  countingWinningDraw := make(map[int]int) 

  for _, winningDraw := range winningDraws {
    if len(winningDraw) == 0 {
      continue
    }
    num, _ := strconv.Atoi(winningDraw)
    countingWinningDraw[num] = 1
  }

  // fmt.Println(countingWinningDraw, "countingWinningDraw")

  for _, draw := range draws {
    drawNum, _ := strconv.Atoi(draw)
    if countingWinningDraw[drawNum] == 1 {
      total++
      countingWinningDraw[drawNum] = 0
    }
  }

  return total - 1
}

func main() {
  currentDir, _ := os.Getwd()
  inputFilePath := ""
  if (len(os.Args) > 1 && os.Args[1] == "test") {
    inputFilePath = fmt.Sprintf("%s/2023/day-04/test-input.txt", currentDir)
  } else {
    inputFilePath = fmt.Sprintf("%s/2023/day-04/input.txt", currentDir)
  }
  
  inputString, _ := os.ReadFile(inputFilePath)
  fmt.Println(string(inputString))

  scanner := bufio.NewScanner(strings.NewReader(string(inputString[:])))

  sum := 0
  scratchCardCount := 0

  cardCount := make(map[int]int)

  for scanner.Scan() {
    game := scanner.Text()
    descriptionParts := strings.Split(game, ":")
    cardDetail := descriptionParts[0]
    cardNumber, _ := strconv.Atoi(strings.Trim(cardDetail[len(cardDetail)-3:], " "))
    winningCount := getTotalWinningCount(descriptionParts[1])
    sum += int(math.Pow(2, float64(winningCount - 1)))
    fmt.Println("cardNumber: ", cardNumber)
    fmt.Println(winningCount, "winningCount")
    
    cardCount[cardNumber] += 1
    for j := cardNumber+1; j <= cardNumber + winningCount; j++ {
      cardCount[j] += cardCount[cardNumber]
    }
  }
  
  fmt.Println("Winning count", sum)

  for _, scratchCard := range cardCount {
    scratchCardCount += scratchCard
  }

  fmt.Println("Scratch card count", scratchCardCount)
}


