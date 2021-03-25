// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import "strconv"

type Cipher int32

const (
	CipherAES128 Cipher = 0
	CipherAES192 Cipher = 1
	CipherAES256 Cipher = 2
)

var EnumNamesCipher = map[Cipher]string{
	CipherAES128: "AES128",
	CipherAES192: "AES192",
	CipherAES256: "AES256",
}

var EnumValuesCipher = map[string]Cipher{
	"AES128": CipherAES128,
	"AES192": CipherAES192,
	"AES256": CipherAES256,
}

func (v Cipher) String() string {
	if s, ok := EnumNamesCipher[v]; ok {
		return s
	}
	return "Cipher(" + strconv.FormatInt(int64(v), 10) + ")"
}