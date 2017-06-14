package game

import (
	"fmt"
	"strconv"
	"time"
)

func WriteSlow(msg string, dur time.Duration) {
	for i := 0; i < len(msg); i++ {
		fmt.Print(string(msg[i]))
		time.Sleep(dur * time.Millisecond)
	}
}

func WritelnSlow(msg string, dur time.Duration) {
	WriteSlow(msg, dur)
	fmt.Print("\n")
}

func WriteS(msg string) {
	WritelnSlow(msg, 80)
}

func WriteWall() {
	WriteS("-------------")
}

func Write(msg string) {
	fmt.Print(msg)
}

func Writeln(msg string) {
	fmt.Println(msg)
}

func GetInputWithOptionsV(options []string, question string) int {
	WriteWall()
	max := len(options)
	for i := 0; i <= max; i++ {
		fmt.Print("[" + strconv.Itoa(i) + "] - " + options[i])
	}
	fmt.Print(question + " ")
	var answer int
	fmt.Scanf("%d", &answer)
	if answer > max || answer < 0 {
		WriteS("Incorrect Input!\nTry Again!")
		return GetInputWithOptionsV(options, question)
	}
	return answer
}

func GetInputWithOptionsH(options []string, question string) int {
	WriteWall()
	max := len(options)
	for i := 0; i <= max; i++ {
		fmt.Print("|" + options[i] + "|   ")
	}
	fmt.Println("")
	for i := 0; i <= max; i++ {
		for c := 0; c < len(options[i]); c++ {
			if c == len(options[i])/2-1 {
				fmt.Print("[")
			} else if c == len(options[i])/2 {
				fmt.Print(strconv.Itoa(i))
			} else if c == len(options[i])/2+1 {
				fmt.Print("]")
			}
			fmt.Print(" ")
		}
		fmt.Print("  ")
	}
	fmt.Print(question + " ")
	var answer int
	fmt.Scanf("%d", &answer)
	if answer > max || answer < 0 {
		WriteS("Incorrect Input!\nTry Again!")
		return GetInputWithOptionsV(options, question)
	}
	return answer
}

func GetInputWithQuestionYesNo(question string) bool {
	WriteWall()
	WriteS(question + "[y/n]: ")
	var answer string
	fmt.Scanf("%s", answer)
	if answer == "y" {
		return true
	} else if answer == "n" {
		return false
	}
	return GetInputWithQuestionYesNo(question)
}

func GetInputWithQuestion(question string) string {
	WriteWall()
	WriteS(question + " ")
	var answer string
	fmt.Scanf("%s", answer)
	return answer
}

func GetEnter() {
	WriteS("Press any key to continiue...")
	fmt.Scanf("")
}
