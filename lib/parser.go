package lib

import "strings"

func GetRepoName(s string) string {
	split := strings.Split(s, "/")
	name := split[len(split)-1]
	return name[:len(name)-4]
}
