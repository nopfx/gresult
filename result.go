package result

import "fmt"

type Result[T any] struct {
	val T
	err error
}

func From[T any](val T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Ok(val)
}

func Ok[T any](val T) Result[T] {
	return Result[T]{val: val, err: nil}
}

func Err[T any](err error) Result[T] {
	var none T
	return Result[T]{val: none, err: err}
}
func (r Result[T]) IsErr() bool {

	return r.err != nil
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic(fmt.Sprintf("called Unwrap on Err: %v", r.err))
	}
	return r.val
}

func (r Result[T]) UnwrapOr(defaultVal T) T {
	if r.err != nil {
		return defaultVal
	}
	return r.val
}

func (r Result[T]) Map(f func(T) T) Result[T] {
	if r.err != nil {
		return r
	}
	return Ok(f(r.val))
}

func (r Result[T]) AndThen(f func(T) Result[T]) Result[T] {
	if r.err != nil {
		return r
	}
	return f(r.val)
}
