package completer

import (
	"reflect"
	"sort"
	"testing"
)

func TestCompleterAddAndLookup(t *testing.T) {
	for _, tc := range []struct {
		add          []string
		wantLookup   map[string]string
		wantComplete map[string][]string
	}{
		{
			add: []string{
				"foo",
			},
			wantComplete: map[string][]string{
				"f":   {"foo"},
				"fo":  {"foo"},
				"foo": {"foo"},
			},
			wantLookup: map[string]string{
				"":    "",
				"f":   "foo",
				"fo":  "foo",
				"foo": "foo",
			},
		},
		{
			add: []string{
				"bar",
				"foo",
			},
			wantComplete: map[string][]string{
				"b":   {"bar"},
				"ba":  {"bar"},
				"bar": {"bar"},
				"f":   {"foo"},
				"fo":  {"foo"},
				"foo": {"foo"},
			},
			wantLookup: map[string]string{
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
			add: []string{
				"bar",
				"baz",
				"foo",
			},
			wantComplete: map[string][]string{
				"b":   {"bar", "baz"},
				"ba":  {"bar", "baz"},
				"bar": {"bar"},
				"baz": {"baz"},
				"f":   {"foo"},
				"fo":  {"foo"},
				"foo": {"foo"},
			},
			wantLookup: map[string]string{
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
			add: []string{
				"bar",
				"baz",
				"foo",
				"fux",
			},
			wantComplete: map[string][]string{
				"b":   {"bar", "baz"},
				"ba":  {"bar", "baz"},
				"bar": {"bar"},
				"baz": {"baz"},
				"f":   {"foo", "fux"},
				"fo":  {"foo"},
				"foo": {"foo"},
				"fu":  {"fux"},
				"fux": {"fux"},
			},
			wantLookup: map[string]string{
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
			add: []string{
				"foo",
				"foobar",
			},
			wantComplete: map[string][]string{
				"f":      {"foo", "foobar"},
				"fo":     {"foo", "foobar"},
				"foo":    {"foo", "foobar"},
				"foob":   {"foobar"},
				"fooba":  {"foobar"},
				"foobar": {"foobar"},
			},
			wantLookup: map[string]string{
				"":       "",
				"f":      "",
				"fo":     "",
				"foo":    "foo",
				"foob":   "foobar",
				"fooba":  "foobar",
				"foobar": "foobar",
			},
		},
		{
			add: []string{
				"foo",
				"foobar",
				"foobaz",
			},
			wantComplete: map[string][]string{
				"f":      {"foo", "foobar", "foobaz"},
				"fo":     {"foo", "foobar", "foobaz"},
				"foo":    {"foo", "foobar", "foobaz"},
				"foob":   {"foobar", "foobaz"},
				"fooba":  {"foobar", "foobaz"},
				"foobar": {"foobar"},
			},
			wantLookup: map[string]string{
				"":       "",
				"f":      "",
				"fo":     "",
				"foo":    "foo",
				"foob":   "",
				"fooba":  "",
				"foobar": "foobar",
				"foobaz": "foobaz",
			},
		},
	} {
		c := NewCompleter()
		for _, s := range tc.add {
			if err := c.Add(s); err != nil {
				t.Errorf("%+v.Add(%q) == %s, want <nil>", c, s, err)
			}
		}
		for prefix, wantLookup := range tc.wantLookup {
			if got, ok := c.Lookup(prefix); got != wantLookup || (got == "" && ok) {
				t.Errorf("%+v.Lookup(%q) == %q, %t, want %q", c, prefix, got, ok, wantLookup)
			}
		}
		for prefix, wantComplete := range tc.wantComplete {
			gotComplete := c.Complete(prefix)
			sort.Strings(gotComplete)
			if !reflect.DeepEqual(gotComplete, wantComplete) {
				t.Errorf("%+v.Complete(%q) == %v, want %v", c, prefix, gotComplete, wantComplete)
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
