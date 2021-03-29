package title

import "fmt"

func SetTitle(text string) {
	fmt.Printf("\033]0;" + text + "\007")
}
