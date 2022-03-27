package types

type MyTypeReturnsInt interface {
	GetInt() int
}

type MyTypeReturnsIntOrString[V int | string] interface {
	GetIntOrString() V
}

type MyTypeReturnsAny[V any] interface {
	GetAny() V
}

type First struct {
	Val int
}

func (f First) GetIntOrString() int {
	return f.Val
}
func (f First) GetAny() int {
	return f.Val
}

type Second struct {
	Val string
}

func (s Second) GetIntOrString() string {
	return s.Val
}

func (s Second) GetAny() string {
	return s.Val
}
