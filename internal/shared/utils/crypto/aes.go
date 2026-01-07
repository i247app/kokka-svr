package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"errors"
)

func evpBytesToKey(password, salt []byte, keyLen, ivLen int) (key, iv []byte) {
	var d, result []byte
	for len(result) < keyLen+ivLen {
		h := md5.New()
		h.Write(d)
		h.Write(password)
		h.Write(salt)
		d = h.Sum(nil)
		result = append(result, d...)
	}
	key = result[:keyLen]
	iv = result[keyLen : keyLen+ivLen]
	return
}

func pkcs7Unpad(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("invalid padding size")
	}
	pad := int(data[len(data)-1])
	if pad == 0 || pad > len(data) {
		return nil, errors.New("invalid padding")
	}
	return data[:len(data)-pad], nil
}

func DecryptCrypto(encryptedBase64 string, passphrase string) (string, error) {
	raw, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return "", err
	}

	if string(raw[:8]) != "Salted__" {
		return "", errors.New("missing Salted__ header")
	}

	salt := raw[8:16]
	ciphertext := raw[16:]

	key, iv := evpBytesToKey([]byte(passphrase), salt, 32, 16)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	plaintext, err := pkcs7Unpad(ciphertext)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
