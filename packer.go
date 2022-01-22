package bingo

type (
	Packer interface {
		Pack(v interface{}) Packer
		Done()
	}
)
