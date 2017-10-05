package completer

import "fmt"

type Completer struct {
	aliases   map[string]string
	originals map[string]struct{}
}

func NewCompleter() Completer {
	return Completer{aliases: make(map[string]string), originals: make(map[string]struct{})}
}

func (c Completer) Add(s string) error {
	if _, ok := c.aliases[s]; ok {
		if _, ok := c.originals[s]; ok {
			return fmt.Errorf("unable to add duplicate key %q", s)
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

func (c Completer) Lookup(s string) (string, bool) {
	got, ok := c.aliases[s]
	return got, ok
}
