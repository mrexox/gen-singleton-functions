package signal

import "fmt"

//go:generate go run ../../main.go github.com/mrexox/gen-singleton-functions/example/signal Signal
type Signal struct{}

func (s Signal) Beep() {
	fmt.Println("beep!")
}
