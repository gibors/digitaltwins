package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecryptString(t *testing.T) {
	encryptedString := "gVjhAY6mHuiDKF0dczUGIjWTd6Uf5Hw6rbipQSZTxtvTOrI6o/SJq951hOM5PsZDJ1IBhi2juVzBrl1mcnN/KniEnHJDW/dUB7vQocfzMLWxUW7DGdcoMxgbLjhAl/bI2dqB5Uwz3dFKkmhWP7Cw/RJ7h5c88N30tbB8LMr8PUOfDNTBZv9brwLHn8c+AFw0"

	value := DecryptString(encryptedString)

	assert.NotEqual(t, "", value)

}
