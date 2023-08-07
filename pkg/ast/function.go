package myast

type Function struct {
	Inputs  []Field
	Outputs []Field
	Body    *Block
}
