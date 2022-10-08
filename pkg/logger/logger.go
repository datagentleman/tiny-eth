package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct{}

var (
	info  = log.New(os.Stdout, fmt.Sprint(pencil, "  ", green_on_green, " INFO  ", reset, " "), log.Ldate|log.Ltime)
	err   = log.New(os.Stdout, fmt.Sprint(fire, " ", white_on_yellow, " ERROR ", reset, " "), log.Ldate|log.Ltime)
	panic = log.New(os.Stdout, fmt.Sprint(exploding_head, " ", white_on_red, " PANIC ", reset, " "), log.Ldate|log.Ltime)
)

// emoji
var (
	pencil         = "\u270f\ufe0f"
	fire           = "\U0001f525"
	exploding_head = "\U0001f92f"
)

// colours
var (
	green           = "\033[32m"
	green_on_green  = "\033[32m\033[40m"
	red             = "\033[31m"
	white_on_yellow = "\033[20m\033[43m"
	white_on_red    = "\033[41m"
	reset           = "\033[0m"
)

func (l *Logger) Info(msg any) {
	info.Println(green, msg, reset)
}

func (l *Logger) Error(msg any) {
	err.Println(red, msg, reset)
}

func (l *Logger) Panic(msg any) {
	panic.Println(white_on_red, msg, reset)
}

func Info(msg any) {
	info.Println(green, msg, reset)
}

func Error(msg any) {
	err.Println(red, msg, reset)
}

func Panic(msg any) {
	panic.Println(white_on_red, msg, reset)
}
