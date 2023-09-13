package ffp

type Option[T any] struct {
	value *T
}

func Some[T any](val T) Option[T] {
	return Option[T]{
		value: &val,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		value: nil,
	}
}

func (o Option[T]) IsSome() bool {
	return o.value != nil
}

func (o Option[T]) IsNone() bool {
	return o.value == nil
}

func (o Option[T]) OrElse(val T) T {
	if o.IsNone() {
		return val
	}
	return *o.value
}

func (o Option[T]) OrEmpty() T {
	var val T
	if o.IsSome() {
		val = *o.value
	}
	return val
}

func (o Option[T]) Get() (T, bool) {
	return o.OrEmpty(), o.IsSome()
}


// func (o Option[T]) Then(f func() (T, bool)) Option[T] {
// 	if o.IsNone() {
// 		return o
// 	}
	
// }