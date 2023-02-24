package orb_test

import (
	"encoding/base64"
	orb2 "github.com/piresrui/orb/orb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestOrbReport(t *testing.T) {
	orb := orb2.ProvideOrb()

	r := orb.Report()
	require.NotEmpty(t, r.CPU)
	require.NotEmpty(t, r.Disk)
	require.NotEmpty(t, r.Temp)
	require.NotEmpty(t, r.Battery)
}

func TestOrbHash(t *testing.T) {
	orb := orb2.ProvideOrb()

	tmpfile, err := os.CreateTemp("", "test")
	require.NoError(t, err)
	defer os.Remove(tmpfile.Name())

	data := []byte("hello, world")
	expectedBase64 := base64.StdEncoding.EncodeToString(data)
	_, err = tmpfile.Write(data)
	assert.NoError(t, err)

	err = tmpfile.Close()
	require.NoError(t, err)

	h, err := orb.Hash(tmpfile.Name())
	require.NoError(t, err)

	require.Equal(t, h.Image, expectedBase64)
	require.NotEmpty(t, h.Key)
}
