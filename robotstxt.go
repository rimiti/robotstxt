package main

import (
  "github.com/rimiti/robotstxt/parser"
)

var result string

func Setup() {
  result = parser.Smile()
}

func GetResult() string {
  return result
}
