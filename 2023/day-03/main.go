package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var matrix [][]int
var inputArray []string;

type Point struct {
  X int
  Y int
}

var xRotation = []int{-1, 0, 1, 1, 1, 0, -1, -1}
var yRotation = []int{-1, -1, -1, 0, 1, 1, 1, 0}

func initMatrix(row int, col int) {
  matrix = make([][]int, row)
  for i := 0; i < row; i++{
    matrix[i] = make([]int, col)
  }
}

func printMatrix() {
  for _, row := range matrix {
    for _, colValue := range row {
      fmt.Printf("%d ", colValue)
    }
    fmt.Printf("\n")
  }
}

func isNumber(s byte) bool {
  return s >= 48 && s <= 57
}

func isValidCoordinate(point Point) bool {
  return point.X >= 0 && point.X < len(matrix) && point.Y >= 0 && point.Y < len(matrix[0])
}

func isValidNumberPoint(point Point) bool {
  for i := 0; i < 8; i++ {
    curX := point.X + xRotation[i]
    curY := point.Y + yRotation[i]
    curAroundPoint := Point{X: curX, Y: curY}

    if !isValidCoordinate(curAroundPoint) {
      continue
    }
    isValidPoint := string(inputArray[curX][curY]) == "." || isNumber(inputArray[curX][curY])
    if !isValidPoint {
      return false 
    }
  }
  return true
}

func findLastNumberIndex(str string, index int) int {
  i := index
  for i < len(str) && isNumber(str[i]) {
    i++
  }

  return i - 1
}

func findFirstNumberIndex(str string, index int) int {
  i := index
  for i >= 0 && isNumber(str[i]) {
    i--
  }

  return i + 1
}

func countNumber1(stringMap []string) int {
  colLength := len(stringMap[0])
  rowLength := len(stringMap)

  sum := 0

  for i := 0; i < rowLength; i++ {
    j := 0
    for j < colLength {
      if !isNumber(inputArray[i][j]) {
        j++;
        continue;
      }

      matrix[i][j] = 1
      lastNumberIndex := findLastNumberIndex(inputArray[i], j)
      curNum, _ := strconv.Atoi(inputArray[i][j:lastNumberIndex+1])
      sum += curNum


      isValid := true
      for z := j; z <= lastNumberIndex; z++ {
        isOk := isValidNumberPoint(Point{X: i, Y: z})
        if !isOk {
          isValid = false 
          break;
        }
      }

      if (isValid) {
        sum = sum - curNum
      }

      j = lastNumberIndex + 1

    }
  }

  return sum
}

func getAdjacentNumber(point Point) []int {
  adjacentNumbers := []int{}
  for i := 0; i < 8; i++ {
    curX := point.X + xRotation[i]
    curY := point.Y + yRotation[i]
    if isNumber(inputArray[curX][curY]) && matrix[curX][curY] == 0 {
      firstNumberIndex := findFirstNumberIndex(inputArray[curX], curY)
      lastNumberIndex := findLastNumberIndex(inputArray[curX], curY)
      for j := firstNumberIndex; j <= lastNumberIndex; j++ {
        matrix[curX][j] = 1
      }
      num, _ := strconv.Atoi(inputArray[curX][firstNumberIndex:lastNumberIndex + 1])
      adjacentNumbers = append(adjacentNumbers, num)
    }
  }

  return adjacentNumbers
}

func countNumber2(stringMap []string) int {
  colLength := len(stringMap[0])
  rowLength := len(stringMap)
  initMatrix(colLength, rowLength)

  sum := 0

  for i := 0; i < rowLength; i++ {
    for j := 0; j < colLength; j++ {
      if string(inputArray[i][j]) == "*" {
        numbers := getAdjacentNumber(Point {X: i, Y: j})
        if len(numbers) == 2 {
          sum += numbers[0] * numbers[1]
        }
      }
    }
  }

  return sum
}

func main() {
  currentDir, _ := os.Getwd()
  inputFilePath := ""
  if (len(os.Args) > 1 && os.Args[1] == "test") {
    inputFilePath = fmt.Sprintf("%s/2023/day-03/test-input.txt", currentDir)
  } else {
    inputFilePath = fmt.Sprintf("%s/2023/day-03/input.txt", currentDir)
  }
  
  inputString, _ := os.ReadFile(inputFilePath)

  inputArray = strings.Split(string(bytes.TrimSpace(inputString)), "\n")

  //result := countNumber1(inputArray)
  result := countNumber2(inputArray)

  fmt.Println(result)

}
