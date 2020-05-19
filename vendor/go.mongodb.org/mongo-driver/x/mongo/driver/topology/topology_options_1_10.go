// +build go1.10

package topology

import "crypto/x509"

func x509CertSubject(cert *x509.Certificate) string {
	return cert.Subject.String()
}
