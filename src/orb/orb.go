package orb

import "github.com/piresrui/orb/config"

type Orb interface {
	// Signup pings API to create new signup key
	Signup(string) error
	// Status pings API with hardware status
	Status() error
}

type VirtualOrb struct {
	Config config.EnvConfig
}

func (v *VirtualOrb) Signup(path string) error {
	return nil
}

func (v *VirtualOrb) Status() error {
	return nil
}
