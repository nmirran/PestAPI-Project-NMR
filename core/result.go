package core

type Result[T any] struct {
	Value T
	Err   error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{Value: value, Err: nil}
}

func ErrResult[T any](err error) Result[T] {
	var empty T
	return Result[T]{Value: empty, Err: err}
}

func (r Result[T]) IsOk() bool {
	return r.Err == nil
}

func (r Result[T]) IsErr() bool {
	return r.Err != nil
}

func (r Result[T]) Unwrap0r(def T) T {
	if r.IsErr() {
		return def
	}

	return r.Value
}

func (r Result[T]) Map (fn func(T) T) Result[T] {
	if r.IsErr() {
		return r
	}
	return Ok(fn(r.Value))
}
