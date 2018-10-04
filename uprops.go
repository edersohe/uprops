package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	trimSpaces  = true
	separatorKV = "="
	commentChar = "#"
	eol         = "\n"
)

func dos2Unix(s string) string {
	var re = regexp.MustCompile(`(?m:\r\n?`)
	return re.ReplaceAllString(s, "\n")
}

func getLines(s string) []string {
	return strings.Split(s, eol)
}

func clean(s string) string {
	var re = regexp.MustCompile(`(?m:(^[ \t]+|[ \t]+$)`)
	return re.ReplaceAllString(s, "")
}

func getPropName(s string) string {
	return clean(strings.SplitN(s, separatorKV, 2)[0])
}

func replace(s string, l string) string {
	var re = regexp.MustCompile(`(?m:^[\t ]*` + regexp.QuoteMeta(getPropName(l)) + `[\t ]*=.*$)`)
	return re.ReplaceAllString(s, l)
}

func main() {
	args := os.Args

	if len(args) != 3 {
		os.Exit(1)
	}

	strBase := dos2Unix(args[1])
	strCustom := dos2Unix(args[2])
	strMerge := strBase

	for _, l := range getLines(strCustom) {
		l = clean(l)
		if l != "" && !strings.HasPrefix(l, commentChar) && strings.Contains(l, separatorKV) {
			strMerge = replace(strMerge, l)
		}
	}

	fmt.Println(strMerge)

	os.Exit(0)
}
