package testing_test

import (
	test "ginkoGomega/testing"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testing Suite")
}

var _ = Describe("Testing Person", func() {

	var person *test.Person

	BeforeEach(func() {

		person = &test.Person{Age: 12}

	})

	Context("When the person is child", func() {

		It("Returns True", func() {

			response := person.IsChild()

			Expect(response).To(BeTrue())

		})
	})

	DescribeTable("IsChild Table Test",
		func(age int, expectedResponse bool) {

			p := test.Person{Age: age}
			response := p.IsChild()

			Expect(response).To(Equal(expectedResponse))
		},
		Entry("When Child", 10, true),
		Entry("Not Child", 19, false))
})
