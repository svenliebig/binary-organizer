package service

import (
	"github.com/svenliebig/binary-organizer/internal/binaries"
	"github.com/svenliebig/binary-organizer/internal/config"
)

type service struct {
	binary binaries.Binary
	config config.Config
}

func New(binary binaries.Binary) (*service, error) {
	config, err := config.Load()
	return &service{
		binary: binary,
		config: config,
	}, err
}

