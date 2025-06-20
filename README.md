# Tabl

[![Test](https://github.com/jpedro/tabl/actions/workflows/test.yaml/badge.svg)](https://github.com/jpedro/tabl/actions/workflows/test.yaml)

Prints out the values of a table (array of arrays) of strings aligned by
width. It also tries to align numbers to the right if **all** are
numeric.


## Usage

```go
package main

import (
	"github.com/jpedro/tabl"
)

func main() {
	table := tabl.New()

	table.Add("KEY", "DESCRIPTION", "COLORED_NUMBERS", "ALMOST_A_NUMBER")
	table.Add("Some metric", "", "1", 1)
	table.Add("Nobody expects the Spanish Inquisition", "Ah!", "1", 1.23)
	table.Add("Uh", "Integer with color codes", "\033[32;1m333\033[0m", "3a")

	table.Print()
}
```

Output:
```
% go run example/main.go
KEY                                      DESCRIPTION                COLORED_NUMBERS   ALMOST_A_NUMBER
Some metric                                                                       1   1
Nobody expects the Spanish Inquisition   Ah!                                      1   1.23
Uh                                       Integer with color codes               333   3a
```

Check [example/main.go](example/main.go).


## Todos

- [ ] Fix the goddamned version to make the go index use the right code

- [ ] Runes, nils and prob others are not printed correctly

- [ ] Columns and Values should be their own structs (avoids duplicated calls)

- [ ] Columns can have unit and separate colors

- [ ] Support alternate row and column colors

- [ ] Highlighted values at row or cell level
