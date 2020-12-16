package hello

import "testing"

func TestHello(t *testing.T) {
	expect := "Hello, world."
	actual := Hello()
	if expect != actual {
		t.Errorf("Hello() = %q, expect = %q", actual, expect)
	}

}

func TestProverb(t *testing.T) {
	expect := "Concurrency is not parallelism."
	actual := Proverb()
	if expect != actual {
		t.Errorf("Proverb() = %q, expect = %q", actual, expect)
	}
}
