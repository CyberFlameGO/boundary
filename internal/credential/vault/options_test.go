package vault

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetOpts(t *testing.T) {
	t.Parallel()
	t.Run("WithName", func(t *testing.T) {
		opts := getOpts(WithName("test"))
		testOpts := getDefaultOptions()
		testOpts.withName = "test"
		assert.Equal(t, opts, testOpts)
	})
	t.Run("WithDescription", func(t *testing.T) {
		opts := getOpts(WithDescription("test desc"))
		testOpts := getDefaultOptions()
		testOpts.withDescription = "test desc"
		assert.Equal(t, opts, testOpts)
	})
	t.Run("WithLimit", func(t *testing.T) {
		opts := getOpts(WithLimit(5))
		testOpts := getDefaultOptions()
		testOpts.withLimit = 5
		assert.Equal(t, opts, testOpts)
	})
	t.Run("WithCACert", func(t *testing.T) {
		opts := getOpts(WithCACert([]byte("test cert")))
		testOpts := getDefaultOptions()
		testOpts.withCACert = []byte("test cert")
		assert.Equal(t, opts, testOpts)
	})
	t.Run("WithNamespace", func(t *testing.T) {
		opts := getOpts(WithNamespace("namespace"))
		testOpts := getDefaultOptions()
		testOpts.withNamespace = "namespace"
		assert.Equal(t, opts, testOpts)
	})
	t.Run("WithTlsServerName", func(t *testing.T) {
		opts := getOpts(WithTlsServerName("server"))
		testOpts := getDefaultOptions()
		testOpts.withTlsServerName = "server"
		assert.Equal(t, opts, testOpts)
	})
	t.Run("WithTlsSkipVerify", func(t *testing.T) {
		opts := getOpts(WithTlsSkipVerify(true))
		testOpts := getDefaultOptions()
		testOpts.withTlsSkipVerify = true
		assert.Equal(t, opts, testOpts)
	})
	t.Run("WithClientCert", func(t *testing.T) {
		testOpts := getDefaultOptions()
		assert.Nil(t, testOpts.withClientCert)
		inCert := testClientCert(t, testCaCert(t))
		cert := inCert.Cert.Cert
		key := inCert.Cert.Key
		clientCert, err := NewClientCertificate(cert, key)
		assert.NoError(t, err)
		assert.NotNil(t, clientCert)
		opts := getOpts(WithClientCert(clientCert))
		require.NotNil(t, opts.withClientCert)
		assert.Equal(t, cert, opts.withClientCert.Certificate)
		assert.Equal(t, key, opts.withClientCert.CertificateKey)
	})
}
