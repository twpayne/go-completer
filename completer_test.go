package completer

import "testing"

func TestCompleter(t *testing.T) {
	c := NewCompleter()
	for _, tc := range []struct {
		add  string
		want map[string]string
	}{
		{
			add: "foo",
			want: map[string]string{
				"":    "",
				"f":   "foo",
				"fo":  "foo",
				"foo": "foo",
			},
		},
		{
			add: "bar",
			want: map[string]string{
				"":    "",
				"b":   "bar",
				"ba":  "bar",
				"bar": "bar",
				"f":   "foo",
				"fo":  "foo",
				"foo": "foo",
			},
		},
		{
			add: "baz",
			want: map[string]string{
				"":    "",
				"b":   "",
				"ba":  "",
				"bar": "bar",
				"baz": "baz",
				"f":   "foo",
				"fo":  "foo",
				"foo": "foo",
			},
		},
		{
			add: "fux",
			want: map[string]string{
				"":    "",
				"b":   "",
				"ba":  "",
				"bar": "bar",
				"baz": "baz",
				"f":   "",
				"fo":  "foo",
				"foo": "foo",
				"fu":  "fux",
				"fux": "fux",
			},
		},
	} {
		c.Add(tc.add)
		for prefix, want := range tc.want {
			if got, ok := c.Lookup(prefix); got != want || (got == "" && ok) {
				t.Errorf("%+v.Lookup(%q) == %q, %t, want %q", c, prefix, got, ok, want)
			}
		}
	}
}

func TestDuplicates(t *testing.T) {

	c := NewCompleter()
	c.Add("foo")
	c.Add("foo")

	prefix := "fo"
	if got, ok := c.Lookup(prefix); !ok {
		t.Errorf("not handling duplicates. %+v.Lookup(%q) == %q, want \"foo\"", c, prefix, got)
	}
}
