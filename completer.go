package completer

type Completer map[string]string

func NewCompleter() Completer {
	return Completer(make(map[string]string))
}

func (c Completer) Add(s string) {
	if _, found := c[s]; found {
		return
	}
	for i := 0; i < len(s); i++ {
		prefix := s[:i+1]
		if _, ok := c[prefix]; ok {
			delete(c, prefix)
		} else {
			c[prefix] = s
		}
	}
}

func (c Completer) Lookup(s string) (string, bool) {
	got, ok := c[s]
	return got, ok
}
