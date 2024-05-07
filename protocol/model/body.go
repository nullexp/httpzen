package model

type Body struct {
	Data   []byte
	Parent uint
}

const (
	Send = iota
	Receive
)
