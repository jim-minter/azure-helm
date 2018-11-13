package tls

import (
	"crypto/rsa"
	"crypto/x509"
)

var (
	dummyPrivateKey  *rsa.PrivateKey
	dummyCertificate *x509.Certificate
)

func init() {
	var err error
	dummyPrivateKey, err = x509.ParsePKCS1PrivateKey([]byte{
		0x30, 0x82, 0x01, 0x39, 0x02, 0x01, 0x00, 0x02, 0x41, 0x00, 0xd2, 0x7a, 0xaa, 0x17, 0x19, 0x88, 0x0f, 0x8a,
		0x95, 0x00, 0xc3, 0x35, 0x58, 0xc0, 0xeb, 0x83, 0x13, 0x22, 0x6f, 0xb0, 0x5d, 0xa5, 0x51, 0xbf, 0x30, 0x9f,
		0xa6, 0x7a, 0x2b, 0x90, 0x5e, 0x3d, 0x44, 0x6d, 0x55, 0xfd, 0x59, 0xdf, 0x54, 0x13, 0x5e, 0x3b, 0x9d, 0x28,
		0xa2, 0xae, 0x5d, 0x72, 0x2a, 0x48, 0x37, 0x7a, 0x3c, 0x9d, 0xea, 0x5c, 0x38, 0xf0, 0x0d, 0x5e, 0x73, 0x68,
		0x54, 0x43, 0x02, 0x03, 0x01, 0x00, 0x01, 0x02, 0x40, 0x1d, 0x21, 0xeb, 0x4e, 0xfd, 0x32, 0xae, 0xec, 0x00,
		0x89, 0xb5, 0x7b, 0x39, 0xc8, 0xa2, 0x10, 0x67, 0x62, 0x99, 0xd0, 0xf0, 0x75, 0x44, 0x66, 0x51, 0x25, 0x97,
		0xd6, 0x4b, 0x17, 0x55, 0x63, 0xa7, 0x32, 0xaa, 0xb1, 0x63, 0xcb, 0x15, 0x0d, 0xee, 0x3e, 0xf5, 0x28, 0xe6,
		0x9a, 0x7d, 0x2d, 0x89, 0xfe, 0xfb, 0x3f, 0x81, 0x97, 0x33, 0x74, 0x25, 0x24, 0x9d, 0x56, 0x90, 0xaa, 0x00,
		0xf1, 0x02, 0x21, 0x00, 0xf3, 0x99, 0xf1, 0xc0, 0xfb, 0xb7, 0x9d, 0xc0, 0x93, 0x6d, 0xfd, 0x86, 0xea, 0x3e,
		0x2c, 0x64, 0x20, 0xe4, 0x56, 0x22, 0x59, 0x4d, 0xc1, 0x15, 0x8d, 0xae, 0x50, 0x3e, 0x03, 0x98, 0x3f, 0xcb,
		0x02, 0x21, 0x00, 0xdd, 0x31, 0x25, 0xc0, 0x4a, 0x7f, 0xda, 0x6b, 0x32, 0x11, 0xe8, 0xac, 0x7a, 0xca, 0x73,
		0x7c, 0xf8, 0xc0, 0xb9, 0x1d, 0x0c, 0xce, 0xb4, 0xe1, 0x2a, 0xa0, 0xea, 0xd7, 0x49, 0x29, 0x3e, 0x69, 0x02,
		0x20, 0x61, 0x22, 0xf0, 0xd0, 0xc8, 0x4f, 0x06, 0x9b, 0xa4, 0xee, 0x46, 0x1b, 0x47, 0x4a, 0xb5, 0x7e, 0xd3,
		0xd2, 0xd9, 0x39, 0xe7, 0x2c, 0x67, 0x23, 0x06, 0x15, 0x0d, 0x30, 0x3d, 0x54, 0xb7, 0x93, 0x02, 0x20, 0x30,
		0x81, 0xaa, 0xa9, 0xb4, 0xac, 0xbd, 0x15, 0x6c, 0xf1, 0x88, 0x78, 0xea, 0xa4, 0xa3, 0x16, 0xf0, 0xe6, 0x32,
		0xb2, 0x30, 0x30, 0xd4, 0x71, 0xdc, 0x8b, 0x55, 0x74, 0xc8, 0xd2, 0x86, 0xa1, 0x02, 0x20, 0x65, 0x56, 0xd4,
		0x31, 0xe3, 0x1d, 0x01, 0x65, 0x2f, 0x10, 0x89, 0xb3, 0xd1, 0xc0, 0xe0, 0x0a, 0x92, 0x8b, 0xfa, 0x3e, 0x34,
		0xd8, 0x99, 0x9b, 0x75, 0x5b, 0xaf, 0x1c, 0xce, 0xfe, 0xc0, 0x05,
	})
	if err != nil {
		panic(err)
	}
	dummyCertificate, err = x509.ParseCertificate([]byte{
		0x30, 0x82, 0x01, 0xea, 0x30, 0x82, 0x01, 0x94, 0xa0, 0x03, 0x02, 0x01, 0x02, 0x02, 0x09, 0x00, 0xf6, 0x74,
		0xff, 0x6b, 0x76, 0x8d, 0xaa, 0xf2, 0x30, 0x0d, 0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x01, 0x01,
		0x0b, 0x05, 0x00, 0x30, 0x15, 0x31, 0x13, 0x30, 0x11, 0x06, 0x03, 0x55, 0x04, 0x03, 0x13, 0x0a, 0x6b, 0x75,
		0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x30, 0x1e, 0x17, 0x0d, 0x31, 0x38, 0x31, 0x31, 0x31, 0x32,
		0x31, 0x37, 0x32, 0x32, 0x32, 0x31, 0x5a, 0x17, 0x0d, 0x32, 0x30, 0x31, 0x31, 0x31, 0x32, 0x31, 0x37, 0x32,
		0x32, 0x32, 0x31, 0x5a, 0x30, 0x15, 0x31, 0x13, 0x30, 0x11, 0x06, 0x03, 0x55, 0x04, 0x03, 0x13, 0x0a, 0x6b,
		0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x30, 0x5c, 0x30, 0x0d, 0x06, 0x09, 0x2a, 0x86, 0x48,
		0x86, 0xf7, 0x0d, 0x01, 0x01, 0x01, 0x05, 0x00, 0x03, 0x4b, 0x00, 0x30, 0x48, 0x02, 0x41, 0x00, 0xd2, 0x7a,
		0xaa, 0x17, 0x19, 0x88, 0x0f, 0x8a, 0x95, 0x00, 0xc3, 0x35, 0x58, 0xc0, 0xeb, 0x83, 0x13, 0x22, 0x6f, 0xb0,
		0x5d, 0xa5, 0x51, 0xbf, 0x30, 0x9f, 0xa6, 0x7a, 0x2b, 0x90, 0x5e, 0x3d, 0x44, 0x6d, 0x55, 0xfd, 0x59, 0xdf,
		0x54, 0x13, 0x5e, 0x3b, 0x9d, 0x28, 0xa2, 0xae, 0x5d, 0x72, 0x2a, 0x48, 0x37, 0x7a, 0x3c, 0x9d, 0xea, 0x5c,
		0x38, 0xf0, 0x0d, 0x5e, 0x73, 0x68, 0x54, 0x43, 0x02, 0x03, 0x01, 0x00, 0x01, 0xa3, 0x81, 0xc6, 0x30, 0x81,
		0xc3, 0x30, 0x0e, 0x06, 0x03, 0x55, 0x1d, 0x0f, 0x01, 0x01, 0xff, 0x04, 0x04, 0x03, 0x02, 0x05, 0xa0, 0x30,
		0x0c, 0x06, 0x03, 0x55, 0x1d, 0x13, 0x01, 0x01, 0xff, 0x04, 0x02, 0x30, 0x00, 0x30, 0x81, 0xa2, 0x06, 0x03,
		0x55, 0x1d, 0x11, 0x04, 0x81, 0x9a, 0x30, 0x81, 0x97, 0x82, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
		0x74, 0x65, 0x73, 0x82, 0x0d, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x82, 0x0d, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x82, 0x0d, 0x6d,
		0x61, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x30, 0x30, 0x30, 0x30, 0x30, 0x32, 0x82, 0x0a, 0x6b, 0x75, 0x62, 0x65,
		0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x82, 0x12, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
		0x2e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x82, 0x16, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
		0x65, 0x73, 0x2e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x73, 0x76, 0x63, 0x82, 0x24, 0x6b, 0x75,
		0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x73,
		0x76, 0x63, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x30, 0x0d,
		0x06, 0x09, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x01, 0x01, 0x0b, 0x05, 0x00, 0x03, 0x41, 0x00, 0xd0, 0x2f,
		0xe7, 0x48, 0xab, 0x78, 0x3b, 0x63, 0x44, 0x36, 0xf5, 0x56, 0xbb, 0xf6, 0x0e, 0xb7, 0x45, 0x7b, 0x19, 0x0e,
		0x85, 0xa8, 0x37, 0x23, 0x7c, 0x03, 0x52, 0x23, 0xa8, 0x22, 0x16, 0xf9, 0x9c, 0xd0, 0x8c, 0x17, 0xb8, 0x34,
		0x02, 0xe4, 0xc9, 0x35, 0x91, 0xc4, 0x42, 0xa3, 0x57, 0xdd, 0x7a, 0x17, 0x6c, 0x58, 0x85, 0x03, 0x5b, 0x8b,
		0x29, 0xd1, 0x4b, 0xb0, 0xfb, 0x8e, 0x59, 0x24,
	})
	if err != nil {
		panic(err)
	}
}

// GetDummyPrivateKey return dummy RSA private key
func GetDummyPrivateKey() *rsa.PrivateKey {
	return dummyPrivateKey
}

// GetDummyCertificate return dummy x509 certificate
func GetDummyCertificate() *x509.Certificate {
	return dummyCertificate
}
