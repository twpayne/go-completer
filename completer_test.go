package completer

import (
	"reflect"
	"sort"
	"testing"
)

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

func TestSubstrings(t *testing.T) {
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
			add: "foobar",
			want: map[string]string{
				"":       "",
				"f":      "",
				"fo":     "",
				"foo":    "foo",
				"foob":   "foobar",
				"fooba":  "foobar",
				"foobar": "foobar",
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

func TestSubstringLookup(t *testing.T) {
	c := NewCompleter()
	c.Add("foor")
	err := c.Add("fo")
	if err != nil {
		t.Errorf("%+v.Add(\"fo\") == %v, want 'nil'", c, err)
	}

	prefix := "f"
	if got, ok := c.Lookup(prefix); ok {
		t.Errorf("%+v.Lookup(%q) == %q, %t, want \"\", false", c, prefix, got, ok)
	}
}

func TestDuplicateKey(t *testing.T) {
	c := NewCompleter()
	c.Add("foo")
	err := c.Add("foo")
	if err == nil {
		t.Errorf("%+v.Add(\"foo\") == nil, want error", c)
	}
}

func TestCompleteFunction(t *testing.T) {
	c := NewCompleter()
	words := []string{"foobar", "foobaz"}
	wantWords := map[string][]string{
		"f":      []string{"foobar", "foobaz"},
		"fo":     []string{"foobar", "foobaz"},
		"foo":    []string{"foobar", "foobaz"},
		"foob":   []string{"foobar", "foobaz"},
		"fooba":  []string{"foobar", "foobaz"},
		"foobar": []string{"foobar"},
	}
	for _, w := range words {
		c.Add(w)
	}
	for prefix, want := range wantWords {
		got := c.Complete(prefix)
		sort.Strings(got)
		sort.Strings(want)
		if !reflect.DeepEqual(got, want) {
			t.Errorf(".Complete(%q) == %q, want %q", prefix, got, want)
		}
	}
}
