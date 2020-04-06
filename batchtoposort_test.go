package batchtoposort

import (
	"testing"

	"gotest.tools/assert"
)

func TestToposortCycle(t *testing.T) {
	x := make(map[string][]string)
	x["a"] = []string{"b"}
	x["b"] = []string{"a"}
	_, err := BatchToposort(x)
	assert.ErrorContains(t, err, "Cycle")
}

func TestToposort(t *testing.T) {
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

	res, err := BatchToposort(x)
	assert.NilError(t, err)
	assert.DeepEqual(t, res, [][]string{
		[]string{
			"b", "a",
		},
		[]string{
			"d", "e", "c",
		},
		[]string{
			"g", "h", "f",
		},
		[]string{
			"j", "i",
		},
	})
}
