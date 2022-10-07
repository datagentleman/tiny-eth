package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct{}

var (
	info  = log.New(os.Stdout, fmt.Sprint(greenCircle, green, " INFO  ", reset), log.Ldate|log.Ltime)
	err   = log.New(os.Stdout, fmt.Sprint(fire, red, " ERROR ", reset), log.Ldate|log.Ltime)
	panic = log.New(os.Stdout, fmt.Sprint(explodingHead, white_on_red, " PANIC ", reset), log.Ldate|log.Ltime)
)

// emoji
var (
	greenHeart    = "\U0001f49a"
	fire          = "\U0001f525"
	explodingHead = "\U0001f92f"
	pushpin       = "\U0001f4cc"
	greenCircle   = "\U0001f7e2"
)

// colours
var (
	green        = "\033[32m"
	red          = "\033[31m"
	white_on_red = "\033[41m"
	reset        = "\033[0m"
)

func (l *Logger) Info(msg string) {
	info.Println(green, msg, reset)
}

func (l *Logger) Error(msg string) {
	err.Println(red, msg, reset)
}

func (l *Logger) Panic(msg string) {
	panic.Println(white_on_red, msg, reset)
}

func Info(msg string) {
	info.Println(green, msg, reset)
}

func Error(msg string) {
	err.Println(red, msg, reset)
}

func Panic(msg string) {
	panic.Println(white_on_red, msg, reset)
}
