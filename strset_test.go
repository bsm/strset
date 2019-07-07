package strset

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set", func() {
	var subject *Set

	BeforeEach(func() {
		subject = New(5)
		Expect(subject.Add("b")).To(BeTrue())
		Expect(subject.Add("d")).To(BeTrue())
		Expect(subject.Add("f")).To(BeTrue())
	})

	It("should have len", func() {
		Expect(subject.Len()).To(Equal(3))
	})

	It("should clear", func() {
		subject.Clear()
		Expect(subject.Len()).To(Equal(0))
	})

	It("should copy", func() {
		dup := new(Set)
		dup.Copy(subject)
		Expect(dup.Slice()).To(Equal([]string{"b", "d", "f"}))
		Expect(subject.Remove("d")).To(BeTrue())
		Expect(dup.Slice()).To(Equal([]string{"b", "d", "f"}))
	})

	It("should add data", func() {
		Expect(subject.Add("c")).To(BeTrue())
		Expect(subject.Add("a")).To(BeTrue())
		Expect(subject.Len()).To(Equal(5))

		Expect(subject.Add("b")).To(BeFalse())
		Expect(subject.Add("c")).To(BeFalse())
		Expect(subject.Add("d")).To(BeFalse())
		Expect(subject.Len()).To(Equal(5))
	})

	It("should remove data", func() {
		Expect(subject.Remove("c")).To(BeFalse())
		Expect(subject.Len()).To(Equal(3))
		Expect(subject.Remove("b")).To(BeTrue())
		Expect(subject.Len()).To(Equal(2))
		Expect(subject.Remove("b")).To(BeFalse())
		Expect(subject.Len()).To(Equal(2))
	})

	It("should check if exists", func() {
		Expect(subject.Has("a")).To(BeFalse())
		Expect(subject.Has("b")).To(BeTrue())
		Expect(subject.Has("c")).To(BeFalse())
		Expect(subject.Has("d")).To(BeTrue())
	})

	It("should check for intersections", func() {
		oth := New(3)
		Expect(subject.Intersects(oth)).To(BeFalse())

		oth.Add("c")
		oth.Add("e")
		Expect(subject.Intersects(oth)).To(BeFalse())

		oth.Add("g")
		oth.Add("d")
		Expect(subject.Intersects(oth)).To(BeTrue())
	})

	It("should intersect", func() {
		oth := Use("b", "c", "d", "x")
		res := new(Set)
		res.Intersection(subject, oth)
		Expect(oth.Slice()).To(Equal([]string{"b", "c", "d", "x"}))
		Expect(subject.Slice()).To(Equal([]string{"b", "d", "f"}))
		Expect(res.Slice()).To(Equal([]string{"b", "d"}))
	})

	It("should intersect with", func() {
		oth := Use("b", "c", "d", "x")
		subject.IntersectionWith(oth)
		Expect(oth.Slice()).To(Equal([]string{"b", "c", "d", "x"}))
		Expect(subject.Slice()).To(Equal([]string{"b", "d"}))
	})

	It("should union", func() {
		oth := Use("b", "c", "d", "x")
		res := new(Set)
		res.Union(subject, oth)
		Expect(oth.Slice()).To(Equal([]string{"b", "c", "d", "x"}))
		Expect(subject.Slice()).To(Equal([]string{"b", "d", "f"}))
		Expect(res.Slice()).To(Equal([]string{"b", "c", "d", "f", "x"}))
	})

	It("should union with", func() {
		oth := Use("b", "c", "d", "x")
		Expect(subject.UnionWith(oth)).To(BeTrue())
		Expect(oth.Slice()).To(Equal([]string{"b", "c", "d", "x"}))
		Expect(subject.Slice()).To(Equal([]string{"b", "c", "d", "f", "x"}))
		Expect(subject.UnionWith(oth)).To(BeFalse())
	})

	It("should marshal/unmarshal JSON", func() {
		bin, err := json.Marshal(subject)
		Expect(err).NotTo(HaveOccurred())
		Expect(bin).To(MatchJSON(`["b","d","f"]`))

		var set *Set
		err = json.Unmarshal([]byte(`["b","c","a"]`), &set)
		Expect(err).NotTo(HaveOccurred())
		Expect(set.Slice()).To(Equal([]string{"a", "b", "c"}))
	})
})

// --------------------------------------------------------------------

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "strset")
}
