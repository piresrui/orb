package orb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/piresrui/orb/config"
	"github.com/piresrui/orb/periodic"
	"log"
	"net/http"
	"time"
)

type Client interface {
	// Signup pings API to create new signup key
	Signup() error
	// Status pings API with hardware status
	Status() error
}

type VirtualOrb struct {
	config *config.EnvConfig
	orb    Orb
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

func (v *VirtualOrb) Run() {
	log.Println("Starting periodic calls...")
	go periodic.Periodic(v.Status, time.Duration(v.config.ReportPeriod))
	go periodic.Periodic(v.Signup, time.Duration(v.config.SignupPeriod))
}

func (v *VirtualOrb) Signup() error {
	log.Println("Making a signup...")
	signup, err := v.orb.Hash(v.config.AssetDir + "/some.png")
	if err != nil {
		return err
	}

	err = v.makeSignupRequest(signup)
	return err
}

func (v *VirtualOrb) Status() error {
	log.Println("Reporting status...")
	report := v.orb.Report()

	err := v.reportStatus(&report)
	return err
}

func (v *VirtualOrb) makeSignupRequest(signup *Signup) error {
	r, err := json.Marshal(signup)
	if err != nil {
		return err
	}
	return post(v.config.Hostname+v.config.SignupPath, r, http.StatusOK)
}

func (v *VirtualOrb) reportStatus(report *Report) error {
	r, err := json.Marshal(report)
	if err != nil {
		return err
	}
	return post(v.config.Hostname+v.config.ReportPath, r, http.StatusOK)
}

func post(api string, body []byte, statusCheck int) error {
	resp, err := http.Post(api, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("request failed to: %s", api))
	}
	return nil
}
