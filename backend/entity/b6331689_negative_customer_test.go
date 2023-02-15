package entity

import(
	"testing"
	. "github.com/onsi/gomega"
	. "github.com/asaskevich/govalidator"
	. "github.com/joon0000/sut-final-lab"
)

func TestCustomerValidatePass (*t testing) {
	g := Newgomegawith(t)
	t.Run("check customerID is valid",func TestCustomerValidatePass (*t testing) {
		c := Customer{
			Name: "kawin",
			Email: "kawin@gmail.com",
			CustomerID: "54321",
		}

		ok,err := govalidator.ValidateStruct(c)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error).Equal("customerID is valid")
	})
}