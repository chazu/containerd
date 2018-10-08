/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package encryption

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"golang.org/x/crypto/openpgp"

	"github.com/pkg/errors"
)

// Uint64ToStringArray converts an array of uint64's to an array of strings
// by applying a format string to each uint64
func Uint64ToStringArray(format string, in []uint64) []string {
	var ret []string

	for _, v := range in {
		ret = append(ret, fmt.Sprintf(format, v))
	}
	return ret
}

// parsePrivateKey tries to parse a private key in DER format first and
// PEM format after, returning an error if the parsing failed
func parsePrivateKey(privKey []byte, prefix string) (interface{}, error) {
	key, err := x509.ParsePKCS8PrivateKey(privKey)
	if err != nil {
		key, err = x509.ParsePKCS1PrivateKey(privKey)
	}
	if err != nil {
		block, _ := pem.Decode(privKey)
		if block == nil {
			return nil, fmt.Errorf("%s: Could not PEM decode private key", prefix)
		}
		key, err = x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
			if err != nil {
				return nil, errors.Wrapf(err, "%s: Could not parse private key", prefix)
			}
		}
	}
	return key, err
}

// IsPrivateKey returns true in case the given byte array represents a private key
func IsPrivateKey(data []byte) bool {
	_, err := parsePrivateKey(data, "")
	return err == nil
}

// parsePublicKey tries to parse a public key in DER format first and
// PEM format after, returning an error if the parsing failed
func parsePublicKey(pubKey []byte, prefix string) (interface{}, error) {
	key, err := x509.ParsePKIXPublicKey(pubKey)
	if err != nil {
		block, _ := pem.Decode(pubKey)
		if block == nil {
			return nil, fmt.Errorf("%s: Could not PEM decode public key", prefix)
		}
		key, err = x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, errors.Wrapf(err, "%s: Could not parse public key", prefix)
		}
	}
	return key, err
}

// IsPublicKey returns true in case the given byte array represents a public key
func IsPublicKey(data []byte) bool {
	_, err := parsePublicKey(data, "")
	return err == nil
}

// ParseCertificate tries to parse a public key in DER format first and
// PEM format after, returning an error if the parsing failed
func parseCertificate(certBytes []byte, prefix string) (*x509.Certificate, error) {
	x509Cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		block, _ := pem.Decode(certBytes)
		if block == nil {
			return nil, fmt.Errorf("%s: Could not PEM decode x509 certificate", prefix)
		}
		x509Cert, err = x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, errors.Wrapf(err, "%s: Could not parse x509 certificate", prefix)
		}
	}
	return x509Cert, err
}

// IsCertificate returns true in case the given byte array represents an x.509 certificate
func IsCertificate(data []byte) bool {
	_, err := parseCertificate(data, "")
	return err == nil
}

/// IsGPGPrivateKeyRing returns true in case the given byte array represents a GPG private key ring file
func IsGPGPrivateKeyRing(data []byte) bool {
	r := bytes.NewBuffer(data)
	_, err := openpgp.ReadKeyRing(r)
	return err == nil
}
