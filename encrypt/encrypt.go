package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

const (
	ivDefValue = "0102030405060708"
)

//PKCS5Padding 不知道是干啥的
func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

//PKCS5UnPadding 不知道是干啥的
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		fmt.Println("unpadding>length")
		return nil
	}

	return src[:(length - unpadding)]
}

//GetKey 获取加密所用的Key
func GetKey(openid string) (int, []byte) {
	s := openid
	h := sha256.New()
	h.Write([]byte(s))

	signature := hex.EncodeToString(h.Sum(nil))
	content := signature[0:32]
	fmt.Println(content)
	return 200, []byte(content)
}

//ASEEncrypt 数据加密
func ASEEncrypt(plaintext []byte, openid string) ([]byte, int, error) {

	s := openid
	h := sha256.New()
	h.Write([]byte(s))

	signature := hex.EncodeToString(h.Sum(nil))
	content := signature[0:32]
	key := []byte(content)

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return nil, 300, errors.New("invalid decrypt key")
	}

	blockSize := block.BlockSize()
	plaintext = PKCS5Padding(plaintext, blockSize)
	iv := []byte(ivDefValue)
	blockMode := cipher.NewCBCEncrypter(block, iv)

	ciphertext := make([]byte, len(plaintext))
	blockMode.CryptBlocks(ciphertext, plaintext)

	return ciphertext, 200, nil
}

//ASEDecrypt 数据解密
func ASEDecrypt(ciphertext []byte, openid string) (string, int, error) {

	//查询数据库，key的值
	s := openid
	h := sha256.New()
	h.Write([]byte(s))

	signature := hex.EncodeToString(h.Sum(nil))
	content := signature[0:32]
	key := []byte(content)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", 300, errors.New("invalid decrypt key")
	}

	blockSize := block.BlockSize()

	if len(ciphertext) < blockSize {
		return "", 300, errors.New("ciphertext too short")
	}

	iv := []byte(ivDefValue)
	if len(ciphertext)%blockSize != 0 {
		return "", 300, errors.New("ciphertext is not a multiple of the block size")
	}

	blockModel := cipher.NewCBCDecrypter(block, iv)

	plaintext := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plaintext, ciphertext)
	plaintext = PKCS5UnPadding(plaintext)
	if plaintext == nil {
		return "", 400, nil
	}
	pass64 := base64.StdEncoding.EncodeToString(plaintext)
	return pass64, 200, nil
}
