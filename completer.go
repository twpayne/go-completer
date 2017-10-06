// Package completer helps implement autocompletion.
package completer

import (
	"fmt"
	"strings"
)

// An errDuplicate is returned when a string is added more than once.
type errDuplicate string

func (e errDuplicate) Error() string {
	return fmt.Sprintf("duplicate key: %q", string(e))
}

// A Completer is a set of strings that can be addressed by their unique
// prefixes.
type Completer struct {
	aliases   map[string]string
	originals map[string]struct{}
}

// NewCompleter returns an empty Completer.
func NewCompleter() Completer {
	return Completer{
		aliases:   make(map[string]string),
		originals: make(map[string]struct{}),
	}
}

// Add adds s to the set of possible completions.
func (c *Completer) Add(s string) error {
	if _, ok := c.aliases[s]; ok {
		if _, ok := c.originals[s]; ok {
			return errDuplicate(s)
		}
	}
	for i := 0; i < len(s); i++ {
		prefix := s[:i+1]
		if _, ok := c.originals[prefix]; ok {
			continue
		}
		if _, ok := c.aliases[prefix]; ok {
			delete(c.aliases, prefix)
		} else {
			c.aliases[prefix] = s
		}
	}
	c.originals[s] = struct{}{}
	return nil
}

// Lookup returns the unique completion of s, or the empty string and false if
// there is no unique completion.
func (c *Completer) Lookup(s string) (string, bool) {
	got, ok := c.aliases[s]
	return got, ok
}

// Complete returns all possible completions of s.
func (c *Completer) Complete(s string) []string {
	// This is O(N*M) where N is the number of originals and M is their length.
	// FIXME Find a more efficient implementation.
	out := []string{}
	for v := range c.originals {
		if strings.HasPrefix(v, s) {
			out = append(out, v)
		}
	}
	return out
}
