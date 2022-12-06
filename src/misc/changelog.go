// dtools : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/misc/changelog.go
// 2022-12-06 08:13:14

package misc

import "fmt"

func Changelog() {
	fmt.Printf("\x1bc")

	fmt.Print(`
VERSION         DATE                    COMMENT
-------         ----                    -------
0.100			2022.12.06              initial, non feature-complete, release
`)
	fmt.Println()
}
