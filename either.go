package gopek

type Either[L any, R any] struct {
	left    *L
	right   *R
	isRight bool
}

func Left[L any, R any](value L) Either[L, R] {
	return Either[L, R]{left: &value, isRight: false}
}

func Right[L any, R any](value R) Either[L, R] {
	return Either[L, R]{right: &value, isRight: true}
}

func (e Either[L, R]) IsRight() bool {
	return e.isRight
}

func (e Either[L, R]) IsLeft() bool {
	return !e.isRight
}

func (e Either[L, R]) GetOrElse(defaultValue R) R {
	if e.IsRight() {
		return *e.right
	}
	return defaultValue
}

func (e Either[L, R]) GetRight() (R, bool) {
	if e.IsRight() {
		return *e.right, true
	}
	var zero R
	return zero, false
}

func (e Either[L, R]) GetLeft() (L, bool) {
	if e.IsLeft() {
		return *e.left, true
	}
	var zero L
	return zero, false
}

func MapEither[L any, R any, R2 any](e Either[L, R], f func(R) R2) Either[L, R2] {
	if e.IsRight() {
		return Right[L](f(*e.right))
	}
	return Left[L, R2](*e.left)
}

func FlatMapEither[L any, R any, R2 any](e Either[L, R], f func(R) Either[L, R2]) Either[L, R2] {
	if e.IsRight() {
		return f(*e.right)
	}
	return Left[L, R2](*e.left)
}
