package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := Validator(TestValidationData(t))
		assert.NoError(t, err)
	})
	t.Run("error", func(t *testing.T) {
		data := TestValidationData(t)
		data.Email = ""

		err := Validator(data)
		assert.Error(t, err)

		errValidation := err.(*ErrValidation)
		assert.NotEmpty(t, errValidation.Errors["email"])
	})
}
