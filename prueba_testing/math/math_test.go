package math

import (
	"fmt"
	"testing"
)

type addTest struct {
	arg1, expected int
}

var addTests = []addTest{
	{3, 6},
	{4, 24},
	{5, 120},
	{0, 1},
	{1, 1},
}

func TestAdd(t *testing.T) {

	for _, test := range addTests {
		if output := Factorial(test.arg1); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(4)
	}
}

func ExampleAdd() {
	fmt.Println(Factorial(4))
	// Output: 24
}
