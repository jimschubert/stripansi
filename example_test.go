package stripansi_test

import (
	"fmt"

	"github.com/jimschubert/stripansi"
)

func ExampleString() {
	result := stripansi.String("\x1b[30mfoo\x1b[0m \x1b[30mbar\x1b[0m \x1b[30mbaz\x1b[0m")
	fmt.Println(result)

	// Output: foo bar baz
}
