package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecryptString(t *testing.T) {
	encryptedString := "KSW2WyinU0AehJHZiXlBt/0iQke6D2/cUEN75Uq94YxHalJsA8Dz82eHC3spu3J5D7dDrpSPaGwD7Qn6fNGhBU4MU3TZvX+BClGnu7PLoMxYfSzamMZnDPJ/6qmdvty0eLOn8iuiRh7tum36D+eg+WbdNSeogCmN+TT4ScenPwU="

	value := DecryptString(encryptedString)

	assert.NotEqual(t, "", value)

}
