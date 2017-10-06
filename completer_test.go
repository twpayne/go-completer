package completer

import (
	"reflect"
	"sort"
	"testing"
)

func TestCompleterAddAndLookup(t *testing.T) {
	for _, tc := range []struct {
		add  []string
		want map[string]string
	}{
		{
			add: []string{"foo"},
			want: map[string]string{
				"":    "",
				"f":   "foo",
				"fo":  "foo",
				"foo": "foo",
			},
		},
		{
			add: []string{"foo", "bar"},
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
			add: []string{"foo", "bar", "baz"},
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
			add: []string{"foo", "bar", "baz", "fux"},
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
		{
			add: []string{"foo", "foobar"},
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
		c := NewCompleter()
		for _, s := range tc.add {
			if err := c.Add(s); err != nil {
				t.Errorf("%+v.Add(%q) == %s, want <nil>", c, s)
			}
		}
		for prefix, want := range tc.want {
			if got, ok := c.Lookup(prefix); got != want || (got == "" && ok) {
				t.Errorf("%+v.Lookup(%q) == %q, %t, want %q", c, prefix, got, ok, want)
			}
		}
	}
}

func TestCompleterSubstringLookup(t *testing.T) {
	c := NewCompleter()
	c.Add("foo")
	if err := c.Add("fo"); err != nil {
		t.Errorf("%+v.Add(\"fo\") == %v, want <nil>", c, err)
	}
	prefix := "f"
	if got, ok := c.Lookup(prefix); ok {
		t.Errorf("%+v.Lookup(%q) == %q, %t, want \"\", false", c, prefix, got, ok)
	}
}

func TestCompleterAddDuplicate(t *testing.T) {
	c := NewCompleter()
	c.Add("foo")
	if err := c.Add("foo"); err == nil {
		t.Errorf("%+v.Add(\"foo\") == <nil>, want !<nil>", c)
	}
}

func TestCompleterComplete(t *testing.T) {
	c := NewCompleter()
	for _, w := range []string{
		"foo",
		"foobar",
		"foobaz",
	} {
		c.Add(w)
	}
	for _, tc := range []struct {
		prefix string
		want   []string
	}{
		{"f", []string{"foo", "foobar", "foobaz"}},
		{"fo", []string{"foo", "foobar", "foobaz"}},
		{"foo", []string{"foo", "foobar", "foobaz"}},
		{"foob", []string{"foobar", "foobaz"}},
		{"fooba", []string{"foobar", "foobaz"}},
		{"foobar", []string{"foobar"}},
	} {
		got := c.Complete(tc.prefix)
		sort.Strings(got)
		sort.Strings(tc.want)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%+v.Complete(%q) == %q, want %q", c, tc.prefix, got, tc.want)
		}
	}
}
