package gopek

import "testing"

func TestMap(t *testing.T) {
	input := []int{2, 3, 4}
	want := []int{4, 6, 8}
	got := Map(input, func(x int) int { return x * 2 })

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Map failed: expected %v, got %v", want, got)
		}
	}
}

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	want := []int{2, 4, 6}
	got := Filter(input, func(x int) bool { return x%2 == 0 })

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Filter failed: expected %v, got %v", want, got)
		}
	}
}

func TestReduce(t *testing.T) {
	input := []int{1, 2, 3, 4}
	want := 10
	got := Reduce(input, 0, func(acc, x int) int { return acc + x })

	if got != want {
		t.Errorf("Reduce failed: expected %v, got %v", want, got)
	}
}

func TestCompose(t *testing.T) {
	double := func(x int) int { return x * 2 }
	increment := func(x int) int { return x + 1 }

	composed := Compose(double, increment)
	result := composed(3)

	if result != 8 {
		t.Errorf("Compose failed: expected 8, got %v", result)
	}
}

func TestPipe(t *testing.T) {
	double := func(x int) int { return x * 2 }
	increment := func(x int) int { return x + 1 }

	piped := Pipe(double, increment)
	result := piped(3)

	if result != 7 {
		t.Errorf("Pipe failed: expected 7, got %v", result)
	}
}

func TestComposeMany(t *testing.T) {
	double := func(x int) int { return x * 2 }
	increment := func(x int) int { return x + 1 }
	square := func(x int) int { return x * x }

	composed := ComposeMany(square, double, increment)
	result := composed(2)

	if result != 36 {
		t.Errorf("ComposeMany failed: expected 36, got %v", result)
	}
}

func TestPipeMany(t *testing.T) {
	double := func(x int) int { return x * 2 }
	increment := func(x int) int { return x + 1 }
	square := func(x int) int { return x * x }

	piped := PipeMany(increment, double, square)
	result := piped(2)

	if result != 36 {
		t.Errorf("PipeMany failed: expected 36, got %v", result)
	}
}

func TestOptionSome(t *testing.T) {
	opt := Some(10)

	if opt.IsNone() || !opt.IsSome() {
		t.Error("Option: expected Some, got None")
	}

	val, err := opt.Get()
	if err != nil || val != 10 {
		t.Errorf("Option Get: expected 10, got %v (err: %v)", val, err)
	}
}

func TestOptionNone(t *testing.T) {
	opt := None[int]()

	if opt.IsSome() {
		t.Error("Option: expected None, got Some")
	}

	_, err := opt.Get()
	if err == nil {
		t.Error("Option Get: expected error on None")
	}
}

func TestOptionMap(t *testing.T) {
	opt := Some(5)
	mapped := MapOption(opt, func(x int) int { return x * 2 })

	val, _ := mapped.Get()
	if val != 10 {
		t.Errorf("Option Map failed: expected 10, got %v", val)
	}
}

func TestOptionFlatMap(t *testing.T) {
	opt := Some(5)
	flat := FlatMapOption(opt, func(x int) Option[int] {
		if x > 0 {
			return Some(x * 3)
		}
		return None[int]()
	})

	val, _ := flat.Get()
	if val != 15 {
		t.Errorf("Option FlatMap failed: expected 15, got %v", val)
	}
}

func TestEitherBasic(t *testing.T) {
	success := Right[string](42)
	fail := Left[string, int]("something went wrong")

	if !success.IsRight() || success.IsLeft() {
		t.Error("Expected Right")
	}

	if !fail.IsLeft() || fail.IsRight() {
		t.Error("Expected Left")
	}

	if v := success.GetOrElse(0); v != 42 {
		t.Errorf("Expected 42, got %v", v)
	}
}

func TestEitherFlatMap(t *testing.T) {
	right := Right[string](21)

	mapped := MapEither(right, func(x int) int { return x * 2 })
	flatMapped := FlatMapEither(mapped, func(x int) Either[string, int] {
		if x > 10 {
			return Right[string](x + 1)
		}
		return Left[string, int]("too small")
	})

	val, ok := flatMapped.GetRight()
	if !ok || val != 21 {
		t.Errorf("Expected Right(21), got %v", val)
	}
}
