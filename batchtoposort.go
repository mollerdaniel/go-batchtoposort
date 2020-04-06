package batchtoposort

import (
	"errors"
)

// BatchToposort get task order.
func BatchToposort(m map[string][]string) ([][]string, error) {
	in := countin(m)
	srt := [][]string{}

	rs := getrs(in)

	for len(rs) > 0 {
		srt = append(srt, rs)
		nr := []string{}
		for _, r := range rs {
			for _, dependent := range m[r] {
				in[dependent] = in[dependent] - 1
				if in[dependent] == 0 {
					nr = append(nr, dependent)
				}
			}
		}

		rs = nr
	}
	if len(getNonrs(in)) > 0 {
		return srt, errors.New("Cycle detected")
	}

	return srt, nil
}

func getrs(m map[string]int) []string {
	r := []string{}
	for k, n := range m {
		if n == 0 {
			r = append(r, k)
		}
	}
	return r
}

func getNonrs(m map[string]int) []string {
	r := []string{}
	for k, n := range m {
		if n != 0 {
			r = append(r, k)
		}
	}
	return r
}

func countin(m map[string][]string) map[string]int {
	counts := make(map[string]int)
	for k, va := range m {
		if _, ok := counts[k]; !ok {
			counts[k] = 0
		}
		x := va
		for _, dep := range x {
			if _, ok := counts[dep]; !ok {
				counts[dep] = 0
			}
			counts[dep] = counts[dep] + 1
		}
	}
	return counts
}
