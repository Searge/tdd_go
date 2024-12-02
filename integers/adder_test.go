package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Log("Testing Add function with inputs 2 and 2")
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	} else {
		t.Logf("Success: expected '%d' and got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
