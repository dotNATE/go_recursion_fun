package main

type holding struct {
	Name   string  `json:"name"`
	Weight float32 `json:"weight"`
}

func newHolding(n string, w float32) holding {
	return holding{
		Name:   n,
		Weight: w,
	}
}
