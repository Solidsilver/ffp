package ffp

import (
	"errors"
	"slices"
)

// Result is a generic struct that holds a value or an error.x
// A result can be checked using the `IsOk()` and `IsErr()` functions.
// A result can be unwrapped with `Error()` and `Get()` functions.
type Result[T any] struct {
	value *T
	err   error
}

func NewResult[T any](val T, err error) Result[T] {
	return Result[T]{
		value: &val,
		err:   err,
	}
}

// Ok creates a new Result of type T with the given value
func Ok[T any](val T) Result[T] {
	return Result[T]{
		value: &val,
		err:   nil,
	}
}

// Err creates a new Result of type T with the given error
func Err[T any](err error) Result[T] {
	return Result[T]{
		value: nil,
		err:   err,
	}
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

func (r Result[T]) IsErr() bool {
	return r.err != nil
}

func (r Result[T]) ErrorIs(otherErr error) bool {
	return errors.Is(r.err, otherErr)
}

func (r Result[T]) Get() (T, error) {
	return r.OrEmpty(), r.err
}

func (r Result[T]) Error() error {
	return r.err
}

// OrElse returns the Result value if the result is Ok, otherwise
// it will return the passed in value
func (r Result[T]) OrElse(val T) T {
	if r.IsOk() {
		return *r.value
	}
	return val
}

// OrElse returns the Result value if the result is Ok, otherwise
// it will return the default value for the given type
func (r Result[T]) OrEmpty() T {
	if r.IsOk() {
		return *r.value
	}
	var val T
	return val
}

// Try runs the given function that possibly returns an error.
// It then returns a Result of the returned values
func Try[T any](f func() (T, error)) Result[T] {
	val, err := f()
	if err != nil {
		return Err[T](err)
	}
	return Ok(val)
}

// ThenTry returns an error result if the previous result was an error.
// Otherwise, it returns the the output of the function as a Result.
func (r Result[T]) ThenTry(f func() (T, error)) Result[T] {
	if r.IsErr() {
		return r
	}
	return Try(f)
}

// Calls the given function with the given argument
// and returns a Result
func Call[T any, V any](f func(T) (V, error), val T) Result[V] {
	value, err := f(val)
	if err != nil {
		return Err[V](err)
	}
	return Ok(value)
}

// ThenCall returns an error result if the previous result was an error.
// Otherwise, it returns the the output of the function as a Result.
func (r Result[T]) ThenCall(f func(T) (T, error)) Result[T] {
	if r.IsErr() {
		return r
	}
	return Try(func() (T, error) { return f(*r.value) })
}

func (r Result[T]) ThenCallResult(f func(T) Result[T]) Result[T] {
	if r.IsErr() {
		return r
	}
	return f(*r.value)
}

// OnlyErr changes the return type of the function to only return
// an error
func OnlyErr[T any, V any](f func(T) (V, error)) func(T) error {
	return func(val T) error {
		_, err := f(val)
		return err
	}
}

type MapperFn[T any, V any] func(T) (V, error)

type Mapper[T any, V any] Result[T]

func (m Mapper[T, V]) Map(fMap MapperFn[T, V]) Result[V] {
	rslt := Result[T](m)
	if rslt.IsErr() {
		return Err[V](rslt.Error())
	}
	mapVal, err := fMap(rslt.OrEmpty())
	if err != nil {
		return Err[V](err)
	}
	return Ok(mapVal)
}


// IfNotError returns an error result if the previous result was an error.
// Otherwise, it returns a Result with the input value.
// IfNotError also takes in an optional array of errors
// to ignore.
func (r Result[T]) IfNotError(f func(T) error, ignore ...error) Result[T] {
	if r.IsErr() {
		return r
	}
	err := f(*r.value)
	if slices.ContainsFunc(ignore, func(err error) bool {
		return errors.Is(err, r.Error())
	}) {
		err = nil
	}
	return Result[T]{
		value: r.value,
		err:   err,
	}
}

