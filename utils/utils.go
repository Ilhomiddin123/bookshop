package utils

import "strings"

func GetParamFromPath(path string) (param string) {
	segments := strings.Split(path, "/")
	param = segments[len(segments)-1]
	return param
}
