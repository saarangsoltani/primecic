package fillers

import "github.com/go-vgo/robotgo"

func PressTab(times int) {
	for i := 0; i < times; i++ {
		robotgo.KeyTap("tab")
	}
}

func SelectSex(gender string) {
	if gender == "male" {
		robotgo.KeyTap("down")
		robotgo.KeyTap("down")
		robotgo.KeyTap("down")
		robotgo.KeyTap("down")
		robotgo.KeyTap("enter")
	} else if gender == "female" {
		robotgo.KeyTap("down")
		robotgo.KeyTap("down")
		robotgo.KeyTap("down")
		robotgo.KeyTap("enter")
	}
}

func Type(content string, args ...int) {
	robotgo.TypeStr(content)
}

func PressKey(key string) {
	robotgo.KeyTap("space")
}
