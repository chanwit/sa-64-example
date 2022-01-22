package entity

import (
	"testing"
	"fmt"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

// ตรวจสอบค่าว่างของชื่อแล้วต้องเจอ Error
func TestUserNameNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	user := User{
		Name: "", // ผิด
		Email: "chanwit@gmail.com",
		StudentID: "B6000000",
		Password: "111",
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(user)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Name cannot be blank"))
}

func TestEmailMustBeValid(t *testing.T) {
	g := NewGomegaWithT(t)

	user := User{
		Name: "Abc",
		Email: "qwe#123", // ผิด
		StudentID: "B6000000",		
		Password: "111",
	}

	ok, err := govalidator.ValidateStruct(user)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Email: qwe#123 does not validate as email"))
}

func TestStudentIDMustBeInValidPattern(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []string{
		"X6000000",
		"BA000000",  // B ตามด้วย A และ \d 6 ตัว
		"B000000",   // B ตามด้วย \d 6 ตัว
		"B00000000", // B ตามด้วย \d 8 ตัว

		"MA000000",  // M ตามด้วย A และ \d 6 ตัว
		"M000000",   // M ตามด้วย \d 6 ตัว
		"M00000000", // M ตามด้วย \d 8 ตัว

		"DA000000",  // D ตามด้วย A และ \d 6 ตัว
		"D000000",   // D ตามด้วย \d 6 ตัว
		"D00000000", // D ตามด้วย \d 8 ตัว
	}

	for _, fixture := range fixtures {
		user := User{
			Name: "Abc",
			Email: "me@example.com",
			StudentID: fixture, // ผิด	
			Password: "111",
		}

		ok, err := govalidator.ValidateStruct(user)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal(fmt.Sprintf(`StudentID: %s does not validate as matches(^[BMD]\d{7}$)`, fixture)))
	}
}
























