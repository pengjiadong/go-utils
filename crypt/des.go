package crypt

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
)

type DesCBC struct {
	IV  []byte
	Key []byte
}

func (c *DesCBC) Encrypt(origData string) (string, error) {
	crypted, err := c.cbcEncrypt([]byte(origData))
	if err != nil {
		return "", err
	}
	base64Str := base64.StdEncoding.EncodeToString(crypted)
	return base64Str, nil
}

func (c *DesCBC) Decrypt(base64Data string) (string, error) {
	crypted, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", err
	}
	origData, err := c.cbcDecrypt(crypted)
	if err != nil {
		return "", err
	}
	return string(origData), nil
}

func (c *DesCBC) cbcEncrypt(origData []byte) ([]byte, error) {
	block, err := des.NewCipher(c.Key)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	origData = pkcs5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, c.IV)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func (c *DesCBC) cbcDecrypt(crypted []byte) ([]byte, error) {
	block, err := des.NewCipher(c.Key)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, c.IV)
	origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = pkcs5UnPadding(origData)
	return origData, nil
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
