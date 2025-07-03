package result

import (
	"fmt"
	"os"
	"testing"
)

func TestSimplestUnwrap(t *testing.T) {
	result := Ok("nopfx").Unwrap()
	if result != "nopfx" {
		t.Errorf("Expected Striong(nopfx) but got some shit with value: %v", result)
	}
}

func TestReadNonExistingFile(t *testing.T) {
	result := From(os.ReadFile("/tmp/nonexistent-file-hopefully"))

	if result.IsOk() {
		t.Errorf("Expected error but got success with value: %v", result.Unwrap())
	}
}

func TestUnwrapOrFallback(t *testing.T) {
	result := From(os.ReadFile("/tmp/definitely-missing"))

	val := result.UnwrapOr([]byte("default"))
	if string(val) != "default" {
		t.Errorf("Expected default fallback value, got: %s", val)
	}
}

func TestMapOnlyOnSuccess(t *testing.T) {
	result := Ok(2).Map(func(x int) int {
		return x * 10
	})

	if result.Unwrap() != 20 {
		t.Errorf("Expected mapped value 20, got: %d", result.Unwrap())
	}
}

func TestMapDoesNothingOnError(t *testing.T) {
	result := Err[int](fmt.Errorf("fail")).Map(func(x int) int {
		return x * 10
	})

	if result.IsOk() {
		t.Errorf("Map should not run on Err")
	}
}

func TestAndThenChains(t *testing.T) {
	double := func(x int) Result[int] {
		return Ok(x * 2)
	}

	result := Ok(3).AndThen(double).AndThen(double)

	if result.Unwrap() != 12 {
		t.Errorf("Expected chained result 12, got %d", result.Unwrap())
	}
}
