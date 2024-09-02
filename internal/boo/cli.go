package boo

import (
	"fmt"
)

var silent bool

func Silent(b bool) {
	silent = b
}

func Intro(s string) {
	if silent {
		return
	}

	fmt.Printf("👻 bo(o) %s\n\n", s)
}

func Body(s string) {
	if silent {
		return
	}

	fmt.Printf("%s\n", s)
}

func Bodyf(s string, a ...any) {
	if silent {
		return
	}

	Body(fmt.Sprintf(s, a...))
}

func Outro(s string) {
	if silent {
		return
	}

	fmt.Printf("\n🍾 bo(o) %s\n", s)
}
