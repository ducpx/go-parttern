package main

import "fmt"

func main() {
	h := NewHouse(WithFloors(4), WithMaterial("concrete"), WithHasFireplace(true))

	fmt.Printf("House %#v\n", h)

}

type House struct {
	Material     string
	HasFireplace bool
	Floors       int
}

type Option func(*House)

func NewHouse(opts ...Option) *House {

	h := &House{}
	for _, opt := range opts {
		opt(h)
	}

	return h
}

func WithMaterial(meterial string) Option {
	return func(h *House) {
		h.Material = meterial
	}
}

func WithHasFireplace(has bool) Option {
	return func(h *House) {
		h.HasFireplace = has
	}
}

func WithFloors(floor int) Option {
	return func(h *House) {
		h.Floors = floor
	}
}
