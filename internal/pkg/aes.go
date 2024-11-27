package pkg

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func PKCS7Padding(ciphertext []byte) []byte {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(plantText []byte) ([]byte, error) {
	length := len(plantText)
	if length == 0 {
		return nil, errors.New("invalid padding size")
	}

	unpadding := int(plantText[length-1])
	if unpadding > length || unpadding == 0 {
		return nil, errors.New("invalid padding")
	}

	pad := plantText[length-unpadding:]
	if !bytes.Equal(pad, bytes.Repeat([]byte{byte(unpadding)}, unpadding)) {
		return nil, errors.New("invalid padding")
	}

	return plantText[:(length - unpadding)], nil
}

// AesEncrypt 加密
func AesEncrypt(plaintext []byte, key []byte, iv []byte) ([]byte, error) {
	// 检查密钥和 IV 的长度
	if len(key) != 32 {
		return nil, errors.New("key length must be 32 bytes (256 bits)")
	}
	if len(iv) != aes.BlockSize {
		return nil, errors.New("IV length must be 16 bytes")
	}

	// PKCS7 填充
	plaintext = PKCS7Padding(plaintext)
	ciphertext := make([]byte, len(plaintext))

	// 创建 AES 加密器
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("failed to create cipher")
	}

	// CBC 模式加密
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	return ciphertext, nil
}

// AesDecrypt 解密
func AesDecrypt(ciphertext []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("failed to create cipher: " + err.Error())
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	plaintext, err := PKCS7UnPadding(ciphertext)
	if err != nil {
		return nil, errors.New("failed to unpad: " + err.Error())
	}

	return plaintext, nil
}
