package dog

//go:generate go run ../../main.go github.com/mrexox/gen-singleton-functions/example/dog *Dog
type Dog struct {
	barkingSound string
}

func (d *Dog) Bark() error {
	return nil
}
