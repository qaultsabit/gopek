package gopek

import "errors"

type Option[T any] struct {
	value   T
	present bool
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: value, present: true}
}

func None[T any]() Option[T] {
	var zero T
	return Option[T]{value: zero, present: false}
}

func (o Option[T]) IsSome() bool {
	return o.present
}

func (o Option[T]) IsNone() bool {
	return !o.present
}

func (o Option[T]) Get() (T, error) {
	if o.present {
		return o.value, nil
	}
	var zero T
	return zero, errors.New("no value present")
}

func (o Option[T]) GetOrElse(defaultValue T) T {
	if o.present {
		return o.value
	}
	return defaultValue
}

func MapOption[T any, R any](opt Option[T], f func(T) R) Option[R] {
	if opt.present {
		return Some(f(opt.value))
	}
	return None[R]()
}

func FlatMapOption[T any, R any](opt Option[T], f func(T) Option[R]) Option[R] {
	if opt.present {
		return f(opt.value)
	}
	return None[R]()
}
