package utils

import (
	"os"

	"github.com/i582/cfmt/cmd/cfmt"
)

func CheckErrors(err error) {
	if err != nil {
		cfmt.Printf("Davo 🥬! -> {{ERROR}}::red|bold\n{{%s}}::red\n", err.Error())
		os.Exit(2)
	}
}

func Error(err string) {
	cfmt.Printf("Davo 🥬! -> {{ERROR}}::red|bold\n{{%s}}::red\n", err)
	os.Exit(2)
}

func Info(msg string) {
	cfmt.Printf("Davo 🥬! -> {{INFO:}}::blue|bold {{%s}}::cyan\n", msg)
}
