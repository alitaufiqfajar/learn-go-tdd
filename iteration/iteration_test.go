package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("expected %q but got %q", want, got)
		}
	}

	t.Run("return 2 times character", func(t *testing.T) {
		got := Repeat("a", 2)
		want := "aa"
		assertCorrectMessage(t, got, want)
	})

	t.Run("return 5 times character", func(t *testing.T) {
		got := Repeat("d", 5)
		want := "ddddd"
		assertCorrectMessage(t, got, want)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}
