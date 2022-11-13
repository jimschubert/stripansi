# stripansi

This is a port of Sindre Sorhus's excellent [chalk/strip-ansi](https://github.com/chalk/strip-ansi) to Go.

## Install

```
$ go get -u github.com/jimschubert/stripansi
```

## Usage

```
import (
	"fmt"
	"github.com/jimschubert/stripansi"
)

func main() {
	result := stripansi.String("\x1b[30mfoo\x1b[0m \x1b[30mbar\x1b[0m \x1b[30mbaz\x1b[0m")
	fmt.Println(result)
}
```

## License

This is [licensed](./LICENSE) under MIT.
