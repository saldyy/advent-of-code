package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var MAX_RED = 12 
var MAX_GREEN = 13
var MAX_BLUE = 14

type BallCount struct {
  R int
  G int
  B int
}

func getGameId(gameIdDescription string) int {
  gameIdString := strings.Split(gameIdDescription, " ")[1]
  gameId, _ := strconv.Atoi(gameIdString)

  return gameId
}

func isValidGame(drawDecription string) bool {
  for _, draw := range strings.Split(drawDecription, "; "){
    balls := strings.Split(draw, ", ")
    ballCountTotal := BallCount{R: 0, G: 0, B: 0}
    for _, ball := range balls {
      ballPart := strings.Split(ball, " ")
      ballCount, _ := strconv.Atoi(ballPart[0])
      ballColor := ballPart[1]
      switch ballColor {
        case "red": 
          ballCountTotal.R += ballCount
        case "green": 
          ballCountTotal.G += ballCount
        case "blue": 
          ballCountTotal.B += ballCount
      }
    }
    if (ballCountTotal.R > MAX_RED || ballCountTotal.G > MAX_GREEN || ballCountTotal.B > MAX_BLUE) {
      return false
    }
  }
  return true;
}

func getGameData(description string) (int, int, bool) {
  gameDescriptionPart := strings.Split(description, ": ")

  return getGameId(gameDescriptionPart[0]), getBallPower(gameDescriptionPart[1]), isValidGame(gameDescriptionPart[1])
}

func getBallPower(drawDescription string) int {
  maxBallCount := BallCount{R: -1, G: -1, B: -1}
  for _, draw := range strings.Split(drawDescription, "; "){
    balls := strings.Split(draw, ", ")
    for _, ball := range balls {
      ballPart := strings.Split(ball, " ")
      ballCount, _ := strconv.Atoi(ballPart[0])
      ballColor := ballPart[1]
      switch ballColor {
        case "red": 
          if (maxBallCount.R < ballCount) {
            maxBallCount.R = ballCount
          }
        case "green": 
          if (maxBallCount.G < ballCount) {
            maxBallCount.G = ballCount
          }
        case "blue": 
          if (maxBallCount.B < ballCount) {
            maxBallCount.B = ballCount
          }
      }
    }
  }

  return maxBallCount.R * maxBallCount.G * maxBallCount.B
}

func main() {
  currentDir, _ := os.Getwd()
  inputFilePath := ""
  if (len(os.Args) > 1 && os.Args[1] == "test") {
    inputFilePath = fmt.Sprintf("%s/2023/day-02/test-input.txt", currentDir)
  } else {
    inputFilePath = fmt.Sprintf("%s/2023/day-02/input.txt", currentDir)
  }
  
  inputString, _ := os.ReadFile(inputFilePath)

  fmt.Println(string(inputString))

  scanner := bufio.NewScanner(strings.NewReader(string(inputString[:])))

  sum := 0
  sumBallPower := 0

  for scanner.Scan() {
    game := scanner.Text()
    gameId, ballPower, isValid := getGameData(game)

    fmt.Printf("Game %d - is valid: %t - Ball Power: %d\n",gameId, isValid, ballPower)
    if isValid {
      sum += gameId
    }
    sumBallPower += ballPower
  }

  fmt.Println(sum)
  fmt.Println(sumBallPower)

}
