package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func isNumber(s byte) bool {
  return s >= 48 && s <= 57
}

func execution(s string) int {
  stringNumbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
  first, last := 0, len(s) - 1
  foundFirst := false;
  foundLast := false;
  var firstNumber int
  var lastNumber int

  for first <= last {
    if (foundFirst && foundLast) {
      break;
    }

    for i, val := range(stringNumbers) {
      if (!foundFirst && first+len(val) <= len(s) && s[first:first+len(val)] == val) {
        foundFirst = true        
        firstNumber = i+1
      }
      if (!foundLast && last-len(val)+1 >= 0 && s[last-len(val)+1:last+1] == val) {
        foundLast = true        
        lastNumber = i+1
      }
    }

    if (!foundFirst) {
      if (!isNumber(s[first])) {
        first += 1
      } else {
        foundFirst = true
        firstNumber,_ = strconv.Atoi(string(s[first]))
      }
    }

    if (!foundLast) {
      if (!isNumber(s[last])) {
        last -= 1
      } else {
        foundLast = true
        lastNumber,_ = strconv.Atoi(string(s[last]))
      }
    }
  }

  return firstNumber*10 + lastNumber
}

func main () {
  current, err := os.Getwd()
  if err != nil {
    panic("cannot read current directory")
  }

  input, err := os.ReadFile(fmt.Sprintf("%s/2023/day-01/input.txt", current))

  if err != nil {
    panic("cannot read input")
  }

  scanner := bufio.NewScanner(strings.NewReader(string(input[:])))

  sum := 0

  for scanner.Scan(){
    s := scanner.Text()
    num := execution(s)
    sum += num
  }
  fmt.Println(sum)
}
