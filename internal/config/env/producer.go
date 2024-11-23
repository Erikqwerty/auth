package env

import (
	"errors"
	"os"
	"strings"

	"github.com/erikqwerty/auth/internal/config"
)

const (
	brockers = "Brockers"
)

type producerConfig struct {
	brockers []string
}

func NewProducerConfig() (config.KafkaProducerConfig, error) {
	bstr := os.Getenv(brockers)
	brokers := strings.Split(bstr, ",")
	if len(brokers) == 0 {
		return nil, errors.New("producer config not found")
	}
	return &producerConfig{brockers: brokers}, nil
}

func (p *producerConfig) Brockers() []string {
	return p.brockers
}
