// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package net

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

// TLSClientConfig represents the standard client TLS config.
type TLSClientConfig struct {
	CaCerts            []string `json:"ca_certs" toml:"ca_certs"`
	Cert               string   `json:"cert" toml:"cert"`
	CertKey            string   `json:"cert_key" toml:"cert_key"`
	InsecureSkipVerify bool     `json:"insecure_skip_verify" toml:"insecure_skip_verify"`
	ServerName         string   `json:"server_name" toml:"server_name"`
}

// TLSConfig returns a tls.Config, may be nil without error if TLS is not configured.
func (c *TLSClientConfig) TLSConfig() (*tls.Config, error) {
	// This check returns a nil (aka, "use the default")
	// tls.Config if no field is set that would have an effect on
	// a TLS connection. That is, any of:
	//     * client certificate settings,
	//     * peer certificate authorities,
	//     * disabled security, or
	//     * an SNI server name.

	if len(c.CaCerts) == 0 &&
		c.CertKey == "" &&
		c.Cert == "" &&
		!c.InsecureSkipVerify && //nolint:gosec
		c.ServerName == "" {
		return nil, nil
	}
	tlsConfig := &tls.Config{
		InsecureSkipVerify: c.InsecureSkipVerify, //nolint:gosec
		Renegotiation:      tls.RenegotiateNever,
	}

	if len(c.CaCerts) != 0 {
		pool, err := makeCertPool(c.CaCerts)
		if err != nil {
			return nil, err
		}

		tlsConfig.RootCAs = pool
	}

	if c.Cert != "" && c.CertKey != "" {
		if err := loadCertificate(tlsConfig, c.Cert, c.CertKey); err != nil {
			return nil, err
		}
	}

	tlsConfig.ServerName = c.ServerName

	return tlsConfig, nil
}

func makeCertPool(certFiles []string) (*x509.CertPool, error) {
	pool := x509.NewCertPool()
	for _, certFile := range certFiles {
		pem, err := os.ReadFile(filepath.Clean(certFile))
		if err != nil {
			return nil, fmt.Errorf("could not read certificate %q: %w", certFile, err)
		}

		if ok := pool.AppendCertsFromPEM(pem); !ok {
			return nil, fmt.Errorf("could not parse any PEM certificates %q: %w", certFile, err)
		}
	}

	return pool, nil
}

func loadCertificate(config *tls.Config, certFile, keyFile string) error {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return fmt.Errorf("could not load keypair %s:%s: %w", certFile, keyFile, err)
	}

	config.Certificates = []tls.Certificate{cert}
	config.BuildNameToCertificate()

	return nil
}

func LoadTLSConfigByBase64(caCerts []string, cert, certKey string, insecureSkipVerify bool, serverName string) (*tls.Config, error) {
	if len(caCerts) == 0 &&
		cert == "" &&
		certKey == "" &&
		!insecureSkipVerify && //nolint:gosec
		serverName == "" {
		return nil, nil
	}

	tlsConfig := &tls.Config{
		Renegotiation:      tls.RenegotiateNever,
		InsecureSkipVerify: insecureSkipVerify, //nolint:gosec
		ServerName:         serverName,
	}

	if len(caCerts) != 0 {
		pool := x509.NewCertPool()
		for _, caCert := range caCerts {
			caCertPEMBlock, err := base64.StdEncoding.DecodeString(caCert)
			if err != nil {
				return nil, fmt.Errorf("could not read caCert data %w", err)
			}
			if ok := pool.AppendCertsFromPEM(caCertPEMBlock); !ok {
				return nil, fmt.Errorf("could not parse any PEM certificates request %w", err)
			}
		}
		tlsConfig.RootCAs = pool
	}

	if cert != "" && certKey != "" {
		certPEMBlock, err := base64.StdEncoding.DecodeString(cert)
		if err != nil {
			return nil, fmt.Errorf("could not read cert data %w", err)
		}
		keyPEMBlock, err := base64.StdEncoding.DecodeString(certKey)
		if err != nil {
			return nil, fmt.Errorf("could not read certKey data %w", err)
		}

		cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
		if err != nil {
			return nil, fmt.Errorf("could not load keypair %w", err)
		}

		tlsConfig.Certificates = []tls.Certificate{cert}
		tlsConfig.BuildNameToCertificate()
	}

	return tlsConfig, nil
}

func DefaultTLSConfigWithInsecureSkipVerify(insecureSkipVerify bool) *tls.Config {
	return &tls.Config{
		Renegotiation:      tls.RenegotiateNever,
		InsecureSkipVerify: insecureSkipVerify, //nolint:gosec
	}
}
