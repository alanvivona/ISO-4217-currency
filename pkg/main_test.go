package currency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	c1, err1 := GetByCode("GBP")
	assert.Equal(t, UnitedKingdomPound, *c1, "GBP should be a valid currency code")
	assert.Equal(t, nil, err1, "GBP should not trigger an error")

	c2, err2 := GetByCode("123")
	assert.Nil(t, c2, "123 should not be a valid currency code")
	assert.NotEqual(t, nil, err2, "123 should trigger an error")
}
