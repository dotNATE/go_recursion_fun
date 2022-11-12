package main

type holding struct {
	Name   string  `json:"name"`
	Weight float32 `json:"weight"`
}

func findHoldingInSlice(s []holding, x string) int {
	for i, n := range s {
		if n.Name == x {
			return i
		}
	}
	return -1
}

func newHolding(n string, w float32) holding {
	return holding{
		Name:   n,
		Weight: w,
	}
}
