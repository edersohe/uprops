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

	strBase   = ""
	strCustom = ""
	strMerge  = ""
)

func dos2Unix(s string) string {
	return strings.Replace(strings.Replace(s, "\r\n", eol, -1), "\r", eol, -1)
}

func getLines(s string) []string {
	return strings.Split(s, eol)
}

func clean(s string) string {
	return strings.TrimSpace(strings.Trim(strings.TrimSpace(s), "\t"))
}

func getPropName(s string) string {
	return clean(strings.SplitN(s, separatorKV, 2)[0])
}

func replace(s string, l string) string {
	var re = regexp.MustCompile(`(?m:^` + regexp.QuoteMeta(getPropName(l)) + `[\t ]*=.*$)`)
	return re.ReplaceAllString(s, l)
}

func main() {
	args := os.Args

	if len(args) != 3 {
		os.Exit(1)
	}

	strBase = dos2Unix(args[1])
	strCustom = dos2Unix(args[2])
	strMerge = strBase

	for _, l := range getLines(strCustom) {
		if clean(l) != "" && !strings.HasPrefix(clean(l), commentChar) {
			strMerge = replace(strMerge, l)
		}
	}

	fmt.Println(strMerge)

	os.Exit(0)
}
