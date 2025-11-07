package tee_test

import (
	"strings"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/tee"
)

func ExampleTee_basic() {
	// echo "Hello World" | tee
	// Note: tee writes to files and stdout; this example just shows stdout
	yup.MustRun(
		Tee(strings.NewReader("Hello World")),
	)
	// Output:
	// Hello World
}

