package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"math/big"
	// "net"
	"os"
	"time"
)

var rootPEM = `-----BEGIN CERTIFICAET-----
MIIDSjCCAjKgAwIBAgIQQwwzsUG5VrNdLZnlKJvBvjANBgkqhkiG9w0BAQsFADBD
MRAwDgYDVQQKEwdPTkVTLkFJMR0wGwYDVQQLExREZXZlbG9wbWVudCBHcm91cCBJ
VDEQMA4GA1UEAxMHb25lcy5haTAgFw0xOTA0MTAwODM5NDZaGA8yMTE5MDMxNzA4
Mzk0NlowQzEQMA4GA1UEChMHT05FUy5BSTEdMBsGA1UECxMURGV2ZWxvcG1lbnQg
R3JvdXAgSVQxEDAOBgNVBAMTB29uZXMuYWkwggEiMA0GCSqGSIb3DQEBAQUAA4IB
DwAwggEKAoIBAQDsDCrCneyaNAMuD/jX+ZzmDVijkXEFGGV3F94xD39mgSp/NbZ5
pXKdTiCPjQ9BKPGto2g40ayCsS+YNXpgdfDpKS7TS63Mqfa5dtTK3hP5WLFdIg4a
3D7PC2E4f/D6xqklFYVLh0JARKT/qcrMpd+nQ1XWSesiCU3HUdpUJ0H12kLtJ1KZ
86kIibT/ya6+ZGW92kF1CqcnJSMsLJl7zsHM5dM4Mbq2fRMUOhcErLzD0ZRz27q8
c6Upu29WuyNIQIIhJJAyCq0upk9avd45vXvHMGT99MtRMPo1ITBE3J1d6j4Vo0AJ
T/1DqW3Cfk7v5dD7kZmsJBvhUbk51TIW2uwRAgMBAAGjODA2MA4GA1UdDwEB/wQE
AwIFoDATBgNVHSUEDDAKBggrBgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MA0GCSqG
SIb3DQEBCwUAA4IBAQCin85m1eV7WLpIX9eQy6QRwETKf3Qt9UKhpiP7bbFr5TpI
vn6KggHRWqgz65vYaYSZbqIKW3KdHr7LmOJWsiE4t9yprpzCwUaE2+kHhQ0NuafM
aSI1tfa+LROOhjYSVL0S+fuLCDj7RtD8AizK6M27TQaFOteQDcet7fIfDbXKG4ej
VDxO+/1aPDPBhf9PJBKBhLOAGcWaCxeTJWpHzxgLW5sHMsgrbmcb6YoTmh6+grm6
xOGbFPskIyb8AjlcnPietTDiYO5edOJ44/JbUMx/Fs/hRW53Nkl09Bt3np21UDaW
PQu8UxLAcso2C8no6v2MxUF8p+7Qkwe0pXTX8KFA
-----END CERTIFICAET-----
`

var privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAw+BbLQK4lzcIVPvSVZCj3Qnrqm+IBp0LmQL8sqDJHUgnYw6W
fTYrFn19AJExCFHBb9QW4zXX1qZIS+m7sGtnZ+JftU98eEGrVwS6As3ciFFUW0+7
GyrKmgFynBvE7rQc6nLmwtTDC3LcSu7bneJyp8bRpAuG6R2VAsv6Ofr9Fdt7MmtT
aUQJ92BpgmP2ugVsYUwoqPBrgJ9jrhjkW1ZaZFCTKdw100GvlF+KCSHbwT+q3FaR
3Ff4TG5mng8OS1ZDRrl8i2hFQ8MjoAibtZ/nEtQ1Rbneq5528xLizJAoFeCRwiEN
m8XKFBeD/66bMW6R9VaxRUzOZr+dBhCYv/kaJwIDAQABAoIBAQCXAvuIIhbqDfWQ
f2zTpuWOXj2gcsguQRlV6CWmGauztZMYHYEMG5zwg+LBcy6YOV+mtZ7HoRU8WYpX
bGXVR5i7UJ11DUjq9cYBGpXF36xsNDeTBk88TKtSPFd93zLezWGVpuwdvSZ6Rx1r
b+dFO/YlP3UUwgZK5xGgzF2oIcbI98oK3IGXw5qF73oSMDAyXNE1eCpUh+T5j+/i
VYCpKjrYfZWpem7PuOWhMmcOds5yUVZF3rFSaXD9X++T8HPwRbaI+eQeDNQx31o1
j5Fq8RgjnYaYae4IR92W8NgzUiKz2dUVx0/oaIe1S7fbDHE9nCatLhBnYlHv6YOf
F6x5Q8GxAoGBANqWQGi8azia9V7SiMVSiu5MBaU7zlJFBzGaRRLnKPD+lf3EnIIV
xQYlonSvhANTJJVJDOUrHTvsj0aEQF8MIA3boB853VmINqmgp8YiIv5xW+HY7IWj
5aT9A6wXLqa1ok0z7qKJSKZ2bqedxbbfqzXxGo9RnqS2uG1VJ1CtnQclAoGBAOVn
AUvp9JT1SmG8CQY+R/WdfeIejsWIn7zXwn7DDR3UOKHS5EbyfwlN3Z7V6dyPD4oS
qOHGi9YhJCOeI6TBU4pvHrBsPUcKhUVzh0PqhWUgRVqUWaHbmgaCReXacnfzs6Ng
KK6AUo7ZLQuR3p15/uwOWCIgfyW20fpqriUtfFBbAoGANM7h5qvGl5lQOEVSx6Em
cWWqY33sXZHRqBkmIOf5yycmY3WRfhf2ToLfGaOhmBAJIYvRcrgUyvperuGaj5vl
TpPFvmWDZXL0Cc0HW2W54JDQauK2WZ78Bv0dWwEKwznUaSRnvZq/E4A/QX1rW7/z
pAeIrCFC23ZNXhSZd/NJoYUCgYBIazEUlwHYOGtH2UV8DBNJbXN8+6qUn2HWCsLi
o46zePh71YitAQarmEI7T0ppj1wXeXaI99XTPlkbSWP1Mfa+f//PMAxULfUOjL8J
6LI9z2c1hhmZqCrCnUOVoVV7mpp+G/AbznTBlbvLqDv/uLCKMqkN9PZOoJ0/EYhy
Qh65ewKBgD4/0h55943ajcEpzVoAdvdXLyP+edTEAZ6m2tMHB0B1bHziGFZt75BW
maKJbxB9gHulZRHr/zdTrbO2V499CCjFLDNNPGh8Lz6YVODqicYnZERnAUmCP89D
e5AUMwrrAiwAA6b6oUmpF/27r2tLduZlJ+FEFsTtSMPA6BMYz7Vq
-----END RSA PRIVATE KEY-----
`

func main() {
	GenerateRootCert()
}

func GenerateUserCert() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)   //把 1 左移 128 位，返回给 big.Int
	serialNumber, _ := rand.Int(rand.Reader, max) //返回在 [0, max) 区间均匀随机分布的一个随机值
	subject := pkix.Name{                         //Name代表一个X.509识别名。只包含识别名的公共属性，额外的属性被忽略。
		Organization:       []string{"ONES.AI"},
		OrganizationalUnit: []string{"Development Group IT"},
		CommonName:         "ones.ai",
	}
	extens := []pkix.Extension{
		pkix.Extension{
			Id:       asn1.ObjectIdentifier{1, 1, 1, 1, 1, 1, 1, 1},
			Critical: true,
			Value:    []byte("ASN1:UTF8String:companyNameTest"),
		},
		pkix.Extension{
			Id:       asn1.ObjectIdentifier{1, 1, 1, 1, 1, 1, 1, 2},
			Critical: true,
			Value:    []byte("ASN1:UTF8String:2019-08-10 00:00:00"),
		},
	}
	template := x509.Certificate{
		SerialNumber: serialNumber, // SerialNumber 是 CA 颁布的唯一序列号，在此使用一个大随机数来代表它
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(100 * 365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, //KeyUsage 与 ExtKeyUsage 用来表明该证书是用来做服务器认证的
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},               // 密钥扩展用途的序列
		//IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
		Extensions: extens,
	}
	pk, _ := rsa.GenerateKey(rand.Reader, 2048) //生成一对具有指定字位数的RSA密钥
	// bobPriKeyEncode, err := x509.ParsePKCS1PrivateKey(bobPriKey)

	block, _ := pem.Decode([]byte(rootPEM))
	if block == nil {
		fmt.Println("failed to decode api certificate")
	}
	rootCert, _ := x509.ParseCertificate(block.Bytes)
	if rootCert == nil {
		fmt.Println("failed to parse api certificate")
	}
	//CreateCertificate基于模板创建一个新的证书
	//第二个第三个参数相同，则证书是自签名的
	//返回的切片是DER编码的证书
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, rootCert, &pk.PublicKey, pk) //DER 格式
	if err != nil {
		fmt.Println(err)
	}

	//certOut, err := os.Create("cert.pem")
	certOut, err := os.OpenFile("root_cert.pem", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICAET", Bytes: derBytes})
	certOut.Close()
	// keyOut, err := os.Create("key.pem")
	keyOut, err := os.OpenFile("root_privata_key.pem", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)

	if err != nil {
		fmt.Println(err)
	}
	pem.Encode(keyOut, &pem.Block{Type: "CERTIFICAET", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()

	fmt.Println("done")
}

func GenerateRootCert() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)   //把 1 左移 128 位，返回给 big.Int
	serialNumber, _ := rand.Int(rand.Reader, max) //返回在 [0, max) 区间均匀随机分布的一个随机值
	subject := pkix.Name{                         //Name代表一个X.509识别名。只包含识别名的公共属性，额外的属性被忽略。
		Organization:       []string{"ONES.AI"},
		OrganizationalUnit: []string{"Development Group IT"},
		CommonName:         "ones.ai",
	}
	template := x509.Certificate{
		SerialNumber:          serialNumber, // SerialNumber 是 CA 颁布的唯一序列号，在此使用一个大随机数来代表它
		Subject:               subject,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(100 * 365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, //KeyUsage 与 ExtKeyUsage 用来表明该证书是用来做服务器认证的
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},               // 密钥扩展用途的序列
		BasicConstraintsValid: true,
		IsCA:         true,
		SubjectKeyId: []byte{1, 2, 3},
	}
	pk, _ := rsa.GenerateKey(rand.Reader, 2048) //生成一对具有指定字位数的RSA密钥
	//CreateCertificate基于模板创建一个新的证书
	//第二个第三个参数相同，则证书是自签名的
	//返回的切片是DER编码的证书
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk) //DER 格式
	if err != nil {
		fmt.Println(err)
	}

	//certOut, err := os.Create("cert.pem")
	certOut, err := os.OpenFile("ca.crt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICAET", Bytes: derBytes})
	certOut.Close()

	// keyOut, err := os.Create("key.pem")
	keyOut, err := os.OpenFile("ca.key", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
	fmt.Println("generate ca done")
}
