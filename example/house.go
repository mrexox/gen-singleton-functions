package test

import (
	"fmt"

	"github.com/mrexox/gen-singleton-functions/example/dog"
	"github.com/mrexox/gen-singleton-functions/example/signal"
)

//go:generate go run ../main.go github.com/mrexox/gen-singleton-functions/example *House
type House struct {
	signals []signal.Signal
	dog     *dog.Dog
}

func (h *House) Dog() *dog.Dog {
	return h.dog
}

func (h *House) EnableSignals(signals []signal.Signal) {
	h.signals = signals
}

func (h *House) Protect(dog *dog.Dog) error {
	for _, s := range h.signals {
		s.Beep()
	}

	if dog == nil {
		return fmt.Errorf("can't protect the house without a dog :(")
	}

	h.dog = dog
	return h.dog.Bark()
}

func (h *House) Protected() (bool, error) {
	if h.dog == nil {
		return false, fmt.Errorf("dog has ran away")
	}

	return true, h.dog.Bark()
}
