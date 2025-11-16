package core
// ini dipakai di funcutils.go

func Map[T any, R any](arr []T, fn func(T) R) []R {
	result := make([]R, len(arr))
	for i, v := range arr {
		result[i] = fn(v)
	}
	return result
}

func Reduce[T any, R any](arr []T, init R, fn func(R, T) R) R {
	acc := init
	for _, v := range arr {
		acc = fn(acc, v)
	}
	return acc
}
