package nzym

import (
	"fmt"
	"sort"
	"strings"
)

// Ls lists all aliases.
func Ls(args []string) {
	nza, spaces := GetConfig(), 0
	keys := make([]string, 0, len(nza))
	for k := range nza {
		keys = append(keys, k)
		if spaces < len(k) {
			spaces = len(k)
		}
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k+strings.Repeat(" ", spaces-len(k)), "->", nza[k])
	}
}
