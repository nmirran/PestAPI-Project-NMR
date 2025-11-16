package core

type Result[T any] struct {
	Value T
	Err   error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{Value: value, Err: nil}
}

func Err[T any](err error) Result[T] {
	var empty T
	return Result[T]{Value: empty, Err: err}
}

func (r Result[T]) IsOk() bool {
	return r.Err == nil
}

func (r Result[T]) IsErr() bool {
	return r.Err != nil
}
