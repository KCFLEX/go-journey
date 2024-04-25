package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Olga")
	if s != "welcome Olga" {
		t.Error("got:", s, "expected", "welcome Olga")
	}
}

func ExampleGreet() {
	d := Greet("Olga")
	fmt.Println(d)
	//Output:
	//welcome Olga
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Olga")
	}
}
