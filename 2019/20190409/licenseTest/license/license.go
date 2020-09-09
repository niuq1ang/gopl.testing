package license

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	// "github.com/bangwork/bang-api/app/utils/config"
	"github.com/bangwork/bang-api/app/utils/log"
)

var (
	licenseDatas *LicenseDatas

	companyNameKey = "1.1.1.1.1.1.1.1"

	rootPEM = `-----BEGIN CERTIFICATE-----
MIIDCjCCAfICCQCVRTEOkvVeRzANBgkqhkiG9w0BAQsFADBHMQswCQYDVQQGEwJD
TjESMBAGA1UECAwJR3VhbmdEb25nMQ0wCwYDVQQKDARvbmVzMRUwEwYDVQQDDAxv
bmVzLmFpIFJvb3QwHhcNMTkwNDEwMDk1NDM5WhcNMjkwNDA3MDk1NDM5WjBHMQsw
CQYDVQQGEwJDTjESMBAGA1UECAwJR3VhbmdEb25nMQ0wCwYDVQQKDARvbmVzMRUw
EwYDVQQDDAxvbmVzLmFpIFJvb3QwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEK
AoIBAQDTQ+6gASvqof7SIpExtxsAZ4XZAdp5jfH8Vx1/V8T+U2aEOKeOGTqgAnke
pK+jxtrvrxRdP5123wiXG8E+ifF8QKy8qer+SPmwmsmn9i8I3ikqJ47M32c4mBj9
kfXlONZeZi1Hmh1XfhvsotELpZ4kPsAFzxG+nSWa1c1bP3I7VZqvSTIgAJ1coi0G
uidDtvhc1EyjqBH14hEOw4qDTTioSZnTEJEKJIpyI1CQP/atzrRGuK9Shgx0hlWf
qnLdW7UqYnbAJE4CHkzV56SJ99kcqwY5myqUg2BTHJDpFFRBtAz5aylT6q5HsI09
nWgSxPcaOHkZRAaQEk4tiitnzIDnAgMBAAEwDQYJKoZIhvcNAQELBQADggEBAGhO
q6oyIKlfsCfWShn9f6BQ5sXAcX5cG3YgYjh6sIZrZ5qjW5wO8rVvgM9xSaUBWW8p
q6Q5GWVWwhea3MCRzgb+tFOm/8HlZMW/EfSdIsZmsmCNfVpNC1Tj/aUhiikrKd29
2ZxkQ/mqR3pF9P/L2S0tmS/VjZl+nahoGREGrUpwqr2dkvC0YBxMqN1jZGcChf0p
KwmSecBDnVJbs+crDWD/QlqMLaaVR6VDjyGU2oqbfNneuZnJeJjHNxx5N9wJCAj7
6FxxPom/A2by6l4jReuGE9GRY2GEo1fx6YtTQyJC6CcyI+2/Y+rld9SL9C+9zSZm
GywY36FeKmS8WJ2mwo0=
-----END CERTIFICATE-----`
)

type LicenseDatas struct {
	IsValid     bool
	ErrDesc     string
	CompanyName string
	ExpiredTime int64
	CertPEM     string
}

func InitLicense() error {
	licenseDatas = new(LicenseDatas)
	certPEM, err := loadCertificate()
	if err != nil {
		return err
	}
	cert, err := parseCertificate(certPEM)
	if err != nil {
		setLiceseInvalid(err)
		return err
	}
	err = setLicenseDataByCertificate(cert)
	if err != nil {
		setLiceseInvalid(err)
		return err
	}
	licenseDatas.CertPEM = string(certPEM)
	licenseDatas.IsValid = true

	licenseDatas := GetLicenseDatas()
	if !licenseDatas.IsValid {
		log.Error("invalid license")
		os.Exit(1)
	}
	fmt.Println("done done done")
	// CheckTimeLimit()
	return nil
}

func GetLicenseDatas() (datas LicenseDatas) {
	datas = *licenseDatas
	return
}

func setLiceseInvalid(err error) {
	licenseDatas.ErrDesc = fmt.Sprintf("license data invalid:%v", err)
	licenseDatas.IsValid = false
}

func loadCertificate() ([]byte, error) {
	// licensePath := config.String("license_path", "")
	certPEM, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		return nil, fmt.Errorf("read license failed:%v", err)
	}
	return certPEM, nil
}

func parseCertificate(certPEM []byte) (*x509.Certificate, error) {
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		return nil, fmt.Errorf("failed to parse root certificate")
	}

	block, _ := pem.Decode(certPEM)
	if block == nil {
		return nil, fmt.Errorf("failed to decode api certificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil || cert == nil {
		return nil, fmt.Errorf("failed to parse api certificate")
	}

	opts := x509.VerifyOptions{
		DNSName: "ones.ai",
		Roots:   roots,
	}

	for _, e := range cert.Extensions {
		//去掉前两位Unicode码
		e.Value = e.Value[2:]
		fmt.Println(e.Id.String(), string(e.Value))
	}
	if _, err = cert.Verify(opts); err != nil {
		return nil, fmt.Errorf("failed to verify certificate:%v", err)
	}

	return cert, nil
}

func setLicenseDataByCertificate(cert *x509.Certificate) error {
	mapCertificateDatas := loadCertificateDatas(cert)

	licenseDatas = new(LicenseDatas)

	//company_name
	licenseDatas.CompanyName = mapCertificateDatas[companyNameKey]

	//expired_time
	licenseDatas.ExpiredTime = cert.NotAfter.Unix()

	return nil
}

func loadCertificateDatas(cert *x509.Certificate) map[string]string {
	datas := make(map[string]string)
	for _, e := range cert.Extensions {
		//去掉前两位Unicode码
		e.Value = e.Value[2:]
		datas[e.Id.String()] = strings.TrimSpace(string(e.Value))
		fmt.Println(e.Id.String(), string(e.Value))
	}
	return datas
}
