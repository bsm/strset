package strset

import (
	"encoding/json"
	"sort"
)

// Set represents a set of unique strings.
type Set struct{ items []string }

// New creates a set with a given cap.
func New(size int) *Set {
	return &Set{items: make([]string, 0, size)}
}

// Use turns a slice into a set, re-using the underlying slice.
// WARNING: this function is destructive and will mutate the passed slice.
func Use(vv ...string) *Set {
	sort.Strings(vv)
	return &Set{items: vv}
}

// Copy sets s to the value of x.
func (s *Set) Copy(x *Set) {
	s.items = append(s.items[:0], x.items...)
}

// Len returns the set length.
func (s *Set) Len() int { return len(s.items) }

// Clear removes all elements from the set s.
func (s *Set) Clear() { s.items = s.items[:0] }

// Equals reports whether the sets s and t have the same elements.
func (s *Set) Equals(t *Set) bool {
	if len(s.items) != len(t.items) {
		return false
	}
	for i, v := range s.items {
		if v != t.items[i] {
			return false
		}
	}
	return true
}

// Add adds x to the set s, and reports whether the set grew.
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

// Remove removes x from the set s, and reports whether the set shrank.
func (s *Set) Remove(v string) bool {
	if pos := sort.SearchStrings(s.items, v); pos < len(s.items) && s.items[pos] == v {
		s.items = s.items[:pos+copy(s.items[pos:], s.items[pos+1:])]
		return true
	}
	return false
}

// Has reports whether x is an element of the set s.
func (s *Set) Has(v string) bool {
	pos := sort.SearchStrings(s.items, v)
	return pos < len(s.items) && s.items[pos] == v
}

// Intersection sets s to the intersection x ∩ y.
func (s *Set) Intersection(x, y *Set) {
	ix, iy := x.items, y.items
	if len(iy) < len(ix) {
		ix, iy = iy, ix
	}

	s.Clear()

	var offset int
	var ok bool
	for _, v := range ix {
		if offset, ok = index(iy, v, offset); ok {
			s.Add(v)
		}
	}
}

// IntersectionWith sets s to the intersection s ∩ x.
func (s *Set) IntersectionWith(x *Set) {
	s.Intersection(s, x)
}

// Intersects reports whether s ∩ x ≠ ∅.
func (s *Set) Intersects(x *Set) bool {
	si, xi := s.items, x.items
	sn, xn := len(si), len(xi)
	if xn < sn {
		si, xi = xi, si
		sn, xn = xn, sn
	}
	if sn == 0 || si[0] > xi[xn-1] || xi[0] > si[sn-1] {
		return false
	}

	offset := 0
	for _, v := range si {
		if pos, ok := index(xi, v, offset); ok {
			return true
		} else if pos >= xn {
			return false
		} else {
			offset = pos
		}
	}
	return false
}

// Union sets s to the union x ∪ y.
func (s *Set) Union(x, y *Set) {
	xi, yi := x.items, y.items

	s.Clear()
	for _, v := range xi {
		s.Add(v)
	}
	for _, v := range yi {
		s.Add(v)
	}
}

// UnionWith sets s to the union s ∪ x, and reports whether s grew.
func (s *Set) UnionWith(x *Set) bool {
	sz := s.Len()
	for _, v := range x.items {
		s.Add(v)
	}
	return s.Len() > sz
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
