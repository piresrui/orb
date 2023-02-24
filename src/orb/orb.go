package orb

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"math/rand"
	"os"
)

type Orb interface {
	// Report returns orb hardware status
	Report() Report
	// Hash hashes a retina image
	Hash(string) (*Signup, error)
}

type orb struct{}

func provideOrb() Orb {
	return &orb{}
}

type Report struct {
	CPU     string `json:"cpu"`
	Disk    string `json:"disk"`
	Temp    string `json:"temp"`
	Battery string `json:"battery"`
}

type Signup struct {
	Image string `json:"image"`
	Key   string `json:"key"`
}

// Report
// This simulates actual hardware information
// On an actual orb it would fetch data from system
func (o *orb) Report() Report {
	cpu := fmt.Sprintf("%d", rand.Intn(100))
	disk := fmt.Sprintf("%d", rand.Intn(1000+100)+100)
	temp := fmt.Sprintf("%d", rand.Intn(100+30)+30)
	battery := fmt.Sprintf("%d", rand.Intn(100))

	return Report{
		CPU:     cpu,
		Disk:    disk,
		Temp:    temp,
		Battery: battery,
	}
}

// Hash
// This simulates hashing an image
func (o *orb) Hash(path string) (*Signup, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	imgBase64Str := base64.StdEncoding.EncodeToString(data)

	return &Signup{
		Key:   uuid.NewString(),
		Image: imgBase64Str,
	}, nil
}
