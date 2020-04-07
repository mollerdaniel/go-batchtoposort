# Go implementation of batch toposort

[![GoDoc](https://godoc.org/github.com/mollerdaniel/go-batchtoposort?status.svg)](https://godoc.org/github.com/mollerdaniel/go-batchtoposort)



```  go
package main

import (
	"fmt"

	"github.com/mollerdaniel/go-batchtoposort"
)

func main() {
	x := make(map[string][]string)
	x["a"] = []string{"c", "f"}
	x["b"] = []string{"d", "e"}
	x["c"] = []string{"f"}
	x["d"] = []string{"f", "g"}
	x["e"] = []string{"h"}
	x["f"] = []string{"i"}
	x["g"] = []string{"j"}
	x["h"] = []string{"j"}
	x["i"] = []string{}
	x["j"] = []string{}

	res, err := batchtoposort.FromMap(x)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	for _, r := range res {
		fmt.Println(r)
	}
}
```

```  go
Output:
[a b]
[c d e]
[f g h]
[i j]
```
