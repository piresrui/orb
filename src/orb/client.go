package orb

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/piresrui/orb/config"
	"net/http"
)

type PeriodicFunc func() error

type Client interface {
	// Signup pings API to create new signup key
	Signup() error
	// Status pings API with hardware status
	Status() error
}

type VirtualOrb struct {
	config *config.EnvConfig
	orb    *orb
}

func ProvideVirtualOrb() *VirtualOrb {
	conf, err := config.ProvideConfig()
	if err != nil {
		return nil
	}
	return &VirtualOrb{
		orb:    provideOrb(),
		config: conf,
	}
}

func (v *VirtualOrb) Signup() error {
	signup, err := v.orb.Hash(v.config.AssetDir + "/some.png")
	if err != nil {
		return err
	}

	err = v.makeSignupRequest(signup)
	return err
}

func (v *VirtualOrb) Status() error {
	report := v.orb.Report()

	err := v.reportStatus(&report)
	return err
}

func (v *VirtualOrb) makeSignupRequest(signup *Signup) error {
	r, err := json.Marshal(signup)
	if err != nil {
		return err
	}
	body := bytes.NewReader(r)
	resp, err := http.Post(v.config.Hostname+v.config.SignupPath, "application/json", body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to signup")
	}

	return nil
}

func (v *VirtualOrb) reportStatus(report *Report) error {
	r, err := json.Marshal(report)
	if err != nil {
		return err
	}
	body := bytes.NewReader(r)
	resp, err := http.Post(v.config.Hostname+v.config.ReportPath, "application/json", body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to report status")
	}

	return nil
}
