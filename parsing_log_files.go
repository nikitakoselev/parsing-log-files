package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`(^\[TRC\]|^\[DBG\]|^\[INF\]|^[WRN]|^\[ERR\]|^\[FTL\])`)

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`"(.*)(?i)password(.*)"`)
	var sum int

	for _, val := range lines {
		res := re.FindAllString(val, -1)
		sum += len(res)
	}
	return sum
}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line\d+`).ReplaceAllString(text, "")
}
func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	for n, line := range lines {
		if match := re.FindStringSubmatch(line); match != nil {
			lines[n] = "[USR] " + match[1] + " " + line
		}
	}
	return lines
}
