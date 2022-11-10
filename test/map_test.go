package test

import (
	"gosampler/util"
	"testing"
)

func TestMap(t *testing.T) {
	m := make(map[int]struct{})
	m[1] = struct{}{}
	m[2] = struct{}{}
	m[3] = struct{}{}
	m[4] = struct{}{}
	m[5] = struct{}{}
	m[6] = struct{}{}

	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
		util.AVOIDUNUSED(keys)
	}

	for k := range m {
		delete(m, k)
	}
}
