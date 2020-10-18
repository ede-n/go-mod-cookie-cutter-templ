package pkg

import "testing"

func TestHello(t *testing.T) {
  want := "Hello"
  if got := hello(); got != want {
    t.Errorf("hello() = %q, want = %q", got, want)
  }
} 
