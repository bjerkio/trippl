package tripletex

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	apiclient "github.com/bjerkio/tripletex-go/client"

	"github.com/bjerkio/trippl/internal/pkg/config"
	"github.com/bjerkio/trippl/internal/pkg/db"
)

// GetClient returns a Tripletex Client with Auth
func GetClient(config config.TripplConfig, db db.KeyValueStore) (*apiclient.Tripletex, runtime.ClientAuthInfoWriter, error) {
	r, err := GetTransport(config, db)
	if err != nil {
		return nil, nil, err
	}
	return apiclient.New(r, strfmt.Default), r.DefaultAuthentication, nil
}
