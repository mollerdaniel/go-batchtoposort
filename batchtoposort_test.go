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
		{
			"a", "b",
		},
		{
			"c", "d", "e",
		},
		{
			"f", "g", "h",
		},
		{
			"i", "j",
		},
	})
}

func TestIndividualGroups(t *testing.T) {
	x := make(map[string][]string)
	// Group A
	x["a"] = []string{"b"}
	x["b"] = []string{"c"}
	x["c"] = []string{"d"}
	x["d"] = []string{}

	// Group B
	x["aa"] = []string{"bb"}
	x["bb"] = []string{"cc"}
	x["cc"] = []string{"dd"}
	x["dd"] = []string{}

	// Edges to C
	x["c"] = []string{"f", "y"}
	x["f"] = []string{}
	x["y"] = []string{}

	// Individual Group in the DAG
	x["foo"] = []string{}

	res, err := FromMap(x)
	assert.NilError(t, err)
	fmt.Println(res)
	assert.DeepEqual(t, res, [][]string{
		{
			"a",
			"aa",
			"d",
			"foo",
		},
		{
			"b",
			"bb",
		},
		{
			"c",
			"cc",
		},
		{
			"dd",
			"f",
			"y",
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
