package bdcrypto

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestR(t *testing.T) {
	b, e := RSAEncryptOfWapBaidu(DefaultRSAPublicKeyModulus, []byte("123"))
	fmt.Println(string(b))
	fmt.Printf("%x, %s\n", b, e)
	fmt.Println("-------------------")
	// b, e = hex.DecodeString(("150fa14de7918fe412060dd68d8cac9b99d580afcde4cf1e8f9e6ca683a6b060d57b4f19d8b479b92887577d7b3ce3f242093f9876c8ced908106e0b3c6c920d44b23d20b81a857a18ae59203fd4c80d5037d557bdb87f6405c07e8f623e78e7d446bc232f64f3be304accc44181fee4b46e1cae8f86645078b496a0d6f50ff9"))
	fmt.Println(string(b))
	b, e = RSADecryptNoPadding(DefaultRSAPrivateKey, b)
	fmt.Println(b, e)
	fmt.Println("-------------------")
	b, e = hex.DecodeString("21d499d07b240334030aa6a9be16bdb21b7026ec182fae2fb827ed118ab9638e02a063de30238fcf10f3eee6dfe7aa98e5781ab5c8254a631e9788a374dc86a49dfb0ec00bc85b1bed29ede375caccfaed7bf5d2ebc65c186c5673396824cf963a4cb06f8d93e90a28a0b0cf47fdb113fadfd74cdf3e270cc7e23cdabf4c71b4")
	fmt.Println(string(b))
	fmt.Println("-------------------")
	b, e = hex.DecodeString("5a3a6f3574d001005de8a968ef0fe85a2f242a10cd82c0e9da71d393974276f44700a5598b2b9e551ca710838bc77e3dfa8cea2feaabd736ea14480f4e4f566c120731f48c7b6de53405c18120c9bbe63d1e1e4a618b21aea140b496968c8e2efd19a3e5c95e2568d90462d5bd03799cda86e91e0dec0e0a652582f844fe962f")
	fmt.Println(string(b))
	fmt.Println("-------------------")
}

func TestReverse(t *testing.T) {
	var b = []byte("1234567890")
	fmt.Println(string(BytesReverse(b)))
}

func TestBase64(t *testing.T) {
	var b = []byte("12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
	en := Base64Encode(b)
	fmt.Println(string(en))
	de := Base64Decode(en)
	fmt.Println(string(de))
}

func BenchmarkReverse(b *testing.B) {
	var by = []byte("Pythonphp123sdif8e83")
	for i := 0; i < b.N; i++ {
		BytesReverse(by)
	}
}

func BenchmarkRSAEncrypt(b *testing.B) {
	var by = []byte("Pythonphp123sdif8e83")
	for i := 0; i < b.N; i++ {
		RSAEncryptOfWapBaidu(DefaultRSAPublicKeyModulus, by)
	}
}
