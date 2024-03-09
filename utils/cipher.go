package utils

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

// ECB PKCS5Padding
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// ECB PKCS5Unpadding
func PKCS5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// Des encryption
func encrypt(origData, key []byte) ([]byte, error) {
	if len(origData) < 1 || len(key) < 1 {
		return nil, errors.New("wrong data or key")
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	if len(origData)%bs != 0 {
		return nil, errors.New("wrong padding")
	}
	out := make([]byte, len(origData))
	dst := out
	for len(origData) > 0 {
		block.Encrypt(dst, origData[:bs])
		origData = origData[bs:]
		dst = dst[bs:]
	}
	return out, nil
}

// Des Decrypt
func decrypt(crypted, key []byte) ([]byte, error) {
	if len(crypted) < 1 || len(key) < 1 {
		return nil, errors.New("wrong data or key")
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	out := make([]byte, len(crypted))
	dst := out
	bs := block.BlockSize()
	if len(crypted)%bs != 0 {
		return nil, errors.New("wrong crypted size")
	}

	for len(crypted) > 0 {
		block.Decrypt(dst, crypted[:bs])
		crypted = crypted[bs:]
		dst = dst[bs:]
	}

	return out, nil
}

// [golang ECB 3DES Encrypt]
func TripleEcbDesEncrypt(origData, key []byte) ([]byte, error) {
	tkey := make([]byte, 24, 24)
	copy(tkey, key)
	k1 := tkey[:8]
	k2 := tkey[8:16]
	k3 := tkey[16:]

	block, err := des.NewCipher(k1)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	origData = PKCS5Padding(origData, bs)

	buf1, err := encrypt(origData, k1)
	if err != nil {
		return nil, err
	}
	buf2, err := decrypt(buf1, k2)
	if err != nil {
		return nil, err
	}
	out, err := encrypt(buf2, k3)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// [golang ECB 3DES Decrypt]
func TripleEcbDesDecrypt(crypted, key []byte) ([]byte, error) {
	tkey := make([]byte, 24, 24)
	copy(tkey, key)
	k1 := tkey[:8]
	k2 := tkey[8:16]
	k3 := tkey[16:]
	buf1, err := decrypt(crypted, k3)
	if err != nil {
		return nil, err
	}
	buf2, err := encrypt(buf1, k2)
	if err != nil {
		return nil, err
	}
	out, err := decrypt(buf2, k1)
	if err != nil {
		return nil, err
	}
	out = PKCS5Unpadding(out)
	return out, nil
}

func AesCbcPkcs7Encrypt(text string, secretKeyBytes, ivBytes []byte) (string, error) {
	// Create a new cipher block using AES in CBC mode with PKCS#7 padding
	block, err := aes.NewCipher(secretKeyBytes)
	if err != nil {
		return "", err
	}

	// Create a cipher.BlockMode for encryption
	cbc := cipher.NewCBCEncrypter(block, ivBytes)

	// Pad the plaintext to a multiple of the block size
	paddedText := pkcs7Padding([]byte(text), aes.BlockSize)

	// Encrypt the padded plaintext
	ciphertext := make([]byte, len(paddedText))
	cbc.CryptBlocks(ciphertext, paddedText)

	// Base64 encode the ciphertext
	encodedText := base64.StdEncoding.EncodeToString(ciphertext)

	return encodedText, nil
}

func pkcs7Padding(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, padtext...)
}

type XMLRsaKey struct {
	Modulus  string
	Exponent string
	P        string
	Q        string
	DP       string
	DQ       string
	InverseQ string
	D        string
}

func b64bigint(str string) *big.Int {
	bInt := &big.Int{}
	decoded, _ := base64.StdEncoding.DecodeString(str)
	bInt.SetBytes(decoded)
	return bInt
}

func SignPKCS1v15FromXml(doc, prKey string) (string, error) {
	var err error
	xrk := XMLRsaKey{}
	if err := xml.Unmarshal([]byte(prKey), &xrk); err != nil {
		return "", err
	}

	key := &rsa.PrivateKey{
		PublicKey: rsa.PublicKey{
			N: b64bigint(xrk.Modulus),
			E: int(b64bigint(xrk.Exponent).Int64()),
		},
		D:      b64bigint(xrk.D),
		Primes: []*big.Int{b64bigint(xrk.P), b64bigint(xrk.Q)},
		Precomputed: rsa.PrecomputedValues{
			Dp:        b64bigint(xrk.DP),
			Dq:        b64bigint(xrk.DQ),
			Qinv:      b64bigint(xrk.InverseQ),
			CRTValues: ([]rsa.CRTValue)(nil),
		},
	}
	msg := []byte(doc)
	msgHash := sha256.New()
	_, err = msgHash.Write(msg)
	if err != nil {
		return "", err
	}
	msgHashSum := msgHash.Sum(nil)
	signature, err := rsa.SignPKCS1v15(crand.Reader, key, crypto.SHA256, msgHashSum)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

func GenerateNewCode7Dig() (randomNumber string) {
	min := 1000000
	max := 9999999

	rand.NewSource(time.Now().UnixNano())
	randomNumber = strconv.Itoa(rand.Intn(max-min+1) + min)
	return
}
