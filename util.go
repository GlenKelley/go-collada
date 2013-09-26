package collada

import (
	"regexp"
	"strconv"
	"strings"
)

func (uri *Uri) Id() (Id, bool) {
	idPattern, _ := regexp.Compile("#([\\w-]+)")
	matches := idPattern.FindStringSubmatch(string(*uri))
	if matches != nil {
		return Id(matches[1]), true
	}
	return Id(""), false
}

func (node *Node) HasGeometry() bool {
	return len(node.InstanceGeometry) > 0
}

func (values *Values) Components() []string {
    return strings.Split(values.V, " ")
}

func (ints *Ints) I() []int {
	ss := ints.Values.Components()
	vs := make([]int, len(ss))
	for i, value := range ss {
		vs[i], _ = strconv.Atoi(value)
	}
	return vs
}

func (floats *Floats) F() []float64 {
	ss := floats.Components()
	vs := make([]float64, len(ss))
	for i, value := range ss {
		vs[i], _ = strconv.ParseFloat(value, 64)
	}
	return vs
}

func (floats *Floats) F32() []float32 {
	ss := floats.Components()
	vs := make([]float32, len(ss))
	for i, value := range ss {
		f, _ := strconv.ParseFloat(value, 32)
        vs[i] = float32(f)
	}
	return vs
}

