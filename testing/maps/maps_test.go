package maps

import "testing"

func TestMap(t *testing.T) {
	dic := Dictionary{"test": "test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dic.Search("test")
		want := "test"
		if got != want {
			t.Errorf("%q != %q ", got, want)
		}

	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dic.Search("unknown")
		assertError(t, got, ErrNotFound)

	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "test1")
	want := "test1"
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word", err)
	}
	if got != want {
		t.Errorf("%q != %q", got, want)
	}
}
func TestDelete(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "test1")
	dictionary.Delete("test")
	_, err := dictionary.Search("test")
	if err != ErrNotFound {
		t.Errorf("Expectes %q to be deleted", "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
