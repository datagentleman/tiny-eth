package logger

import (
	"log"
)

type Logger struct{}

// emoji
var (
	greenHeart = "\U0001f49a"
	fire       = "\U0001f525"
)

// colours
var (
	pink  = "\033[35m"
	red   = "\033[41m"
	reset = "\033[0m"
)

func (l *Logger) Info(msg string) {
	log.Println(greenHeart, pink, msg, reset)
}

func (l *Logger) Error(msg string) {
	log.Println(fire, red, msg, reset)
}
