package cipher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const key = "AAAABBBBCCCC"
const iv = "AAAABBBBCCCCDDDD"
const crypted = `Zo+mD1I8fzF1FSnVDWacu8xDbYennAzEcHea9vDJ9fivIiGl8CXe1NOAc9tzB0Eq`
const decrypted = `{"tenant_id":"1","salt":"secret"}`

func TestEncrypt(t *testing.T) {
	result := Encrypt([]byte(key), []byte(iv), decrypted)
	assert.Equal(t, crypted, result)

	result = Decrypt([]byte(key), []byte(iv), result)
	assert.Equal(t, decrypted, result)
}

func TestDecrypt(t *testing.T) {
	result := Decrypt([]byte(key), []byte(iv), crypted)
	assert.Equal(t, decrypted, result)
}