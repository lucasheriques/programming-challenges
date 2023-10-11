package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	pattern := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return pattern.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[\~\*\=\-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) (numOfStrings int) {
	re := regexp.MustCompile(`"[^"\n]*[pP][aA][sS][sS][wW][oO][rR][dD][^"\n]*"`)

	for _, str := range lines {
		if re.MatchString(str) {
			numOfStrings++
		}
	}

	return
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`\bend-of-line\d+\b`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`(User\s+)(\S+)`)
	taggedLines := make([]string, len(lines))

	for i, line := range lines {
		if re.MatchString(line) {
			username := re.FindStringSubmatch(line)[2]
			taggedLines[i] = fmt.Sprintf("[USR] %s %s", username, line)
		} else {
			taggedLines[i] = line
		}
	}

	return taggedLines
}
