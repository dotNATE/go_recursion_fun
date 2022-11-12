package main

type Holding struct {
	Name   string  `json:"name"`
	Weight float32 `json:"weight"`
}

func findHoldingInSlice(s []Holding, x string) int {
	for i, n := range s {
		if n.Name == x {
			return i
		}
	}
	return -1
}

func newHolding(n string, w float32) Holding {
	return Holding{
		Name:   n,
		Weight: w,
	}
}
