package orb

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/piresrui/orb/config"
	"github.com/piresrui/orb/periodic"
	"log"
	"net/http"
	"time"
)

type VirtualOrb interface {
	// Signup pings API to create new signup key
	Signup() error
	// Status pings API with hardware status
	Status() error
	// Run sets up the periodic calls
	Run(ctx context.Context)
}

type virtualOrb struct {
	config *config.EnvConfig
	orb    Orb
	client *http.Client
}

func NewVirtualOrb(orb Orb, client *http.Client, conf *config.EnvConfig) VirtualOrb {
	return &virtualOrb{
		orb:    orb,
		client: client,
		config: conf,
	}
}

func ProvideVirtualOrb() VirtualOrb {
	conf, err := config.ProvideConfig()
	if err != nil {
		return nil
	}
	client := http.DefaultClient
	return &virtualOrb{
		orb:    ProvideOrb(),
		config: conf,
		client: client,
	}
}

func (v *virtualOrb) Run(ctx context.Context) {
	log.Println("Starting periodic calls...")
	go periodic.Periodic(ctx, v.Status, time.Duration(v.config.ReportPeriod))
	go periodic.Periodic(ctx, v.Signup, time.Duration(v.config.SignupPeriod))
}

func (v *virtualOrb) Signup() error {
	log.Println("Making a signup...")
	signup, err := v.orb.Hash(v.config.AssetDir + "/some.png")
	if err != nil {
		return err
	}

	err = v.makeSignupRequest(signup)
	return err
}

func (v *virtualOrb) Status() error {
	log.Println("Reporting status...")
	report := v.orb.Report()

	err := v.reportStatus(&report)
	return err
}

func (v *virtualOrb) makeSignupRequest(signup *Signup) error {
	r, err := json.Marshal(signup)
	if err != nil {
		return err
	}
	return v.post(v.config.Hostname+v.config.SignupPath, r, http.StatusOK)
}

func (v *virtualOrb) reportStatus(report *Report) error {
	r, err := json.Marshal(report)
	if err != nil {
		return err
	}
	return v.post(v.config.Hostname+v.config.ReportPath, r, http.StatusOK)
}

func (v *virtualOrb) post(api string, body []byte, statusCheck int) error {
	resp, err := v.client.Post(api, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("request failed to: %s", api))
	}
	return nil
}
