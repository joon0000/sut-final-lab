package entity

import(
	"testing"
	. "github.com/onsi/gomega"
	. "github.com/asaskevich/govalidator"
	. "github.com/joon0000/sut-final-lab"
)

func TestCustomerValidatePass (*t testing) {
	g := Newgomegawith(t)
	t.Run("pass case",func TestCustomerValidatePass (*t testing) {
		c := Customer{
			Name: "kawin",
			Email: "kawin@gmail.com",
			CustomerID: "L7654321",
		}

		ok,err := govalidator.ValidateStruct(c)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error).Error.Equal("pass case")
	})
}