package tool

import (
	"regexp"
	"strconv"
)

type FileNameList []string

func (s FileNameList) Len() int {
	return len(s)
}
func (s FileNameList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s FileNameList) Less(i, j int) bool {
	var iValue, jValue int

	exp := regexp.MustCompile("\\d+")
	if result := exp.FindAllString(s[i], -1); result != nil {
		if value, err := strconv.Atoi(result[0]); err == nil {
			iValue = value
		}
	}

	if result := exp.FindAllString(s[j], -1); result != nil {
		if value, err := strconv.Atoi(result[0]); err == nil {
			jValue = value
		}
	}

	return iValue < jValue
}
