package Format

import (
	"strconv"
	"strings"
)

func Formated(numbers [6]int) []string {
	var result string
	for _, nb := range numbers {
		if nb < 10 {
			result += "   _   \n" + " /   \\ \n" + "   " + strconv.Itoa(nb) + "   \n" + " \\ _ / \n"
		} else {
			result += "   __   \n" + " /    \\ \n" + "   " + strconv.Itoa(nb) + "   \n" + " \\ __ / \n"
		}
	}
	str := strings.Split(result, "\n")
	return str
}
