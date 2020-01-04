package nzym

import "fmt"

// Mv can be used to give an alias a different
// name. It can also be used to swap two existing
// aliases if destination alias already exists.
func Mv(args []string) {
	fmt.Println(args)
}
