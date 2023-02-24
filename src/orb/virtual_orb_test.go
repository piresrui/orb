package orb_test

import (
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/piresrui/orb/config"
	"github.com/piresrui/orb/orb"
	"github.com/piresrui/orb/orb/mocks"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func provideTestConfig(url string) *config.EnvConfig {
	return &config.EnvConfig{
		Hostname:   url,
		SignupPath: "/signup",
		ReportPath: "/status",
	}
}

func TestReport(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOrb := mocks.NewMockOrb(ctrl)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/status" {
			require.Equal(t, r.URL.Path, "/status")
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	conf := provideTestConfig(server.URL)
	o := orb.NewVirtualOrb(mockOrb, server.Client(), conf)

	mockOrb.EXPECT().Report().Return(orb.Report{
		CPU:     "10",
		Disk:    "10",
		Temp:    "10",
		Battery: "10",
	}).Times(1)

	err := o.Status()
	require.NoError(t, err)
}

func TestReportFailedRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOrb := mocks.NewMockOrb(ctrl)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/status" {
			require.Equal(t, r.URL.Path, "/status")
		}
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()
	conf := provideTestConfig(server.URL)
	o := orb.NewVirtualOrb(mockOrb, server.Client(), conf)

	mockOrb.EXPECT().Report().Return(orb.Report{
		CPU:     "10",
		Disk:    "10",
		Temp:    "10",
		Battery: "10",
	}).Times(1)

	err := o.Status()
	require.Error(t, err)
	require.ErrorContains(t, err, "request failed to")
}

func TestSignup(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOrb := mocks.NewMockOrb(ctrl)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/signup" {
			require.Equal(t, r.URL.Path, "/signup")
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	conf := provideTestConfig(server.URL)
	o := orb.NewVirtualOrb(mockOrb, server.Client(), conf)

	hash := uuid.NewString()
	mockOrb.EXPECT().Hash(gomock.Any()).Return(&orb.Signup{Image: "base64", Key: hash}, nil).Times(1)

	err := o.Signup()
	require.NoError(t, err)
}

func TestSignupFailedRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOrb := mocks.NewMockOrb(ctrl)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/signup" {
			require.Equal(t, r.URL.Path, "/signup")
		}
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()
	conf := provideTestConfig(server.URL)
	o := orb.NewVirtualOrb(mockOrb, server.Client(), conf)

	hash := uuid.NewString()
	mockOrb.EXPECT().Hash(gomock.Any()).Return(&orb.Signup{Image: "base64", Key: hash}, nil).Times(1)

	err := o.Signup()
	require.Error(t, err)
	require.ErrorContains(t, err, "request failed to")
}
