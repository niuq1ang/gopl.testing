package license

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/bangwork/bang-api/app/utils/config"
	"github.com/bangwork/bang-api/app/utils/log"
)

var (
	licenseDatas *LicenseDatas

	companyNameKey     = "1.1.1.1.1.1.1.1"
	projectExpireTime  = "1.1.1.1.1.1.1.2"
	wikiExpireTime     = "1.1.1.1.1.1.1.3"
	testcaseExpireTime = "1.1.1.1.1.1.1.4"
	pipelineExpireTime = "1.1.1.1.1.1.1.5"
	maxTeamUserCount   = "1.1.1.1.1.1.1.6"
	orgIdentity        = "1.1.1.1.1.1.1.7"
)

type LicenseDatas struct {
	IsValid            bool
	ErrDesc            string
	CompanyName        string
	ProjectExpireTime  int64
	WikiExpireTime     int64
	TestcaseExpireTime int64
	PipelineExpireTime int64
	MaxMemberCount     int64
	OrgIdentity        string
	ExpiredTime        int64
	CertPEM            string
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
	CheckTimeLimit()
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
	licensePath := config.String("license_path", "")
	certPEM, err := ioutil.ReadFile(licensePath)
	if err != nil {
		return nil, fmt.Errorf("read license failed:%v", err)
	}
	return certPEM, nil
}

func parseCertificate(certPEM []byte) (*x509.Certificate, error) {
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(RootPEM))
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

	var err error
	licenseDatas.MaxMemberCount, err = strconv.ParseInt(mapCertificateDatas[maxTeamUserCount], 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse maxUserCount:%v", err)
	}

	licenseDatas.ProjectExpireTime, err = strconv.ParseInt(mapCertificateDatas[projectExpireTime], 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse projectExpireTime:%v", err)
	}

	licenseDatas.WikiExpireTime, err = strconv.ParseInt(mapCertificateDatas[wikiExpireTime], 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse wikiExpireTime:%v", err)
	}

	licenseDatas.TestcaseExpireTime, err = strconv.ParseInt(mapCertificateDatas[testcaseExpireTime], 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse testcaseExpireTime:%v", err)
	}

	licenseDatas.PipelineExpireTime, err = strconv.ParseInt(mapCertificateDatas[pipelineExpireTime], 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse pipelineExpireTime:%v", err)
	}

	licenseDatas.OrgIdentity = mapCertificateDatas[orgIdentity]

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
	}
	return datas
}
