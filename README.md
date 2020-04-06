# go-batchtoposort
Go implementation of batch toposort

```
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

	res, err := batchtoposort.BatchToposort(x)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	for _, r := range res {
		fmt.Println(r)
	}
}
```

```
Output:
[a b]
[c d e]
[f g h]
[i j]
```