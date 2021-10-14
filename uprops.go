package main

import (
	"fmt"
	"os"
	"regexp"
)

func getProperties(s string) [][]string {
	var re = regexp.MustCompile(`(?m)^[\t ]*([^#\r\n].+[^\t ])[\t ]*=((?:.*\\\s*\n)*.*)$`)
	return re.FindAllStringSubmatch(s, -1)
}

func replace(s string, k string, v string) string {
	var re = regexp.MustCompile(`(?m)^[\t ]*` + regexp.QuoteMeta(k) + `[\t ]*=(?:.*\\\s*\n)*.*$`)
	return re.ReplaceAllString(s, v)
}

func main() {
	args := os.Args

	if len(args) != 3 {
		os.Exit(1)
	}

	strBase := args[1]
	strCustom := args[2]
	strMerge := strBase

    properties := getProperties(strCustom)
	for i := range properties {
        //fmt.Printf("property: %s\n\nkey: %s\n\nvalue: %s\n\n\n", properties[i][0], properties[i][1], properties[i][2])
        strMerge = replace(strMerge, properties[i][1], properties[i][0])
	}

	fmt.Println(strMerge)

	os.Exit(0)
}
