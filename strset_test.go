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
		Expect(subject.Exists("a")).To(BeFalse())
		Expect(subject.Exists("b")).To(BeTrue())
		Expect(subject.Exists("c")).To(BeFalse())
		Expect(subject.Exists("d")).To(BeTrue())
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

	It("should marshal/unmarshal JSON", func() {
		bin, err := json.Marshal(subject)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(bin)).To(Equal(`["b","d","f"]`))

		var set *Set
		err = json.Unmarshal([]byte(`["b","c","a"]`), &set)
		Expect(err).NotTo(HaveOccurred())
		Expect(set.Slice()).To(Equal([]string{"a", "b", "c"}))
	})

})

// --------------------------------------------------------------------

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "intset")
}
