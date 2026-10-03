package Result

func PrintResult(lines []string, numbers [6]int) string {
	var result string
	for i := 0; i < 4; i++ {
		for count := 0; count < len(numbers); count++ {
			if count == len(numbers)-1 {
				if i == 0 {
					result += "   "
				} else {
					result += " | "
				}
			}
			result += lines[i+(count*4)]
		}
		result += "\n"
	}
	return result
}
