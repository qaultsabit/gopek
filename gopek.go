package gopek

func Map[T any, R any](items []T, fn func(T) R) []R {
	result := make([]R, len(items))
	for i, item := range items {
		result[i] = fn(item)
	}
	return result
}

func Filter[T any](items []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func Reduce[T any, R any](items []T, initial R, reducer func(R, T) R) R {
	acc := initial
	for _, item := range items {
		acc = reducer(acc, item)
	}
	return acc
}

func Compose[A any, B any, C any](f func(B) C, g func(A) B) func(A) C {
	return func(a A) C {
		return f(g(a))
	}
}

func Pipe[A any, B any, C any](f func(A) B, g func(B) C) func(A) C {
	return func(a A) C {
		return g(f(a))
	}
}

func ComposeMany[T any](fns ...func(T) T) func(T) T {
	return func(x T) T {
		for i := len(fns) - 1; i >= 0; i-- {
			x = fns[i](x)
		}
		return x
	}
}

func PipeMany[T any](fns ...func(T) T) func(T) T {
	return func(x T) T {
		for _, fn := range fns {
			x = fn(x)
		}
		return x
	}
}
