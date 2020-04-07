package batchtoposort

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestFromMapCycle(t *testing.T) {
	x := make(map[string][]string)
	x["a"] = []string{"b"}
	x["b"] = []string{"a"}
	_, err := FromMap(x)
	assert.ErrorContains(t, err, "Cycle")
}

func TestFromMap(t *testing.T) {
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

	res, err := FromMap(x)
	assert.NilError(t, err)
	assert.DeepEqual(t, res, [][]string{
		[]string{
			"a", "b",
		},
		[]string{
			"c", "d", "e",
		},
		[]string{
			"f", "g", "h",
		},
		[]string{
			"i", "j",
		},
	})
}

func ExampleFromMap() {
	x := make(map[string][]string)
	x["a"] = []string{"b"}
	x["b"] = []string{"c"}
	x["c"] = []string{}
	x["d"] = []string{"b"}
	x["e"] = []string{"b"}

	r, _ := FromMap(x)
	fmt.Println(r)
	// Output: [[a d e] [b] [c]]
}
