package strset

import (
	"encoding/json"
	"sort"
)

type Set struct{ items []string }

// New creates a set with a given cap
func New(size int) *Set {
	return &Set{items: make([]string, 0, size)}
}

// Use turns a slice into a set, re-using the underlying slice
// WARNING: this function is destructive and will mutate the passed slice
func Use(vv ...string) *Set {
	sort.Strings(vv)
	return &Set{items: vv}
}

// Len returns the set length
func (s *Set) Len() int { return len(s.items) }

// Add adds an item to the set
func (s *Set) Add(v string) bool {
	if pos := sort.SearchStrings(s.items, v); pos < len(s.items) {
		if s.items[pos] == v {
			return false
		}

		s.items = append(s.items, "")
		copy(s.items[pos+1:], s.items[pos:])
		s.items[pos] = v
	} else {
		s.items = append(s.items, v)
	}
	return true
}

// Remove removes an item from the set
func (s *Set) Remove(v string) bool {
	if pos := sort.SearchStrings(s.items, v); pos < len(s.items) && s.items[pos] == v {
		s.items = s.items[:pos+copy(s.items[pos:], s.items[pos+1:])]
		return true
	}
	return false
}

// Exists checks the existence
func (s *Set) Exists(v string) bool {
	pos := sort.SearchStrings(s.items, v)
	return pos < len(s.items) && s.items[pos] == v
}

// Intersects checks if intersectable
func (s *Set) Intersects(t *Set) bool {
	ls, lt := len(s.items), len(t.items)
	if lt < ls {
		ls, lt = lt, ls
		s, t = t, s
	}
	if ls == 0 || s.items[0] > t.items[lt-1] || t.items[0] > s.items[ls-1] {
		return false
	}

	offset := 0
	for _, v := range s.items {
		pos, ok := index(t.items, v, offset)
		if ok {
			return true
		} else if pos >= lt {
			return false
		}
		offset = pos
	}
	return false
}

// Slice returns the string slice
func (s *Set) Slice() []string { return s.items }

// MarshalJSON encodes the set as JSON
func (s *Set) MarshalJSON() ([]byte, error) { return json.Marshal(s.items) }

// UnmarshalJSON decodes JSON into a set
func (s *Set) UnmarshalJSON(data []byte) error {
	var vv []string
	if err := json.Unmarshal(data, &vv); err != nil {
		return err
	}

	*s = *Use(vv...)
	return nil
}

func index(vs []string, v string, offset int) (int, bool) {
	pos := sort.SearchStrings(vs[offset:], v) + offset
	return pos, pos < len(vs) && vs[pos] == v
}
