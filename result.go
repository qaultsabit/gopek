package gopek

type Result[T any] = Either[error, T]

func Ok[T any](value T) Result[T] {
	return Right[error](value)
}

func Err[T any](err error) Result[T] {
	return Left[error, T](err)
}

func MapResult[T any, R any](res Result[T], f func(T) R) Result[R] {
	return MapEither(res, f)
}

func FlatMapResult[T any, R any](res Result[T], f func(T) Result[R]) Result[R] {
	return FlatMapEither(res, f)
}

func Try[T any](f func() (T, error)) Result[T] {
	val, err := f()
	if err != nil {
		return Err[T](err)
	}
	return Ok(val)
}
