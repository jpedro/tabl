# tablelize

Prints out the values of a table (array of arrays) of strings aligned by
width. It also tries to align numbers to the right if **all** are
numeric.


### Usage

```go
package main

import (
    "github.com/jpedro/tablelize"
)

func main() {
    var data [][]any

    data = append(data, []any{"KEY", "VALUE", "NUMBER", "ALMOST_NUMBER"})
    data = append(data, []any{"char", "a", "1", "1"})
    data = append(data, []any{"longer-key-name", "Some text", "-2", "2"})
    data = append(data, []any{"key", "And now for something completely different", "3", "3a"})

    tablelize.Rows(data)
}
```

Output:
```
% go run example/main.go
KEY               VALUE                                        NUMBER   ALMOST_NUMBER
char              a                                                 1   1
longer-key-name   Some text                                        -2   2
key               And now for something completely different        3   3a
```

Check [example/main.go](example/main.go).
