//go:build provider
// +build provider

// Package main contains a runnable Provider Pact test example.
package main

import (
	"os"
	"testing"

	"github.com/pact-foundation/pact-go/v2/log"
	"github.com/pact-foundation/pact-go/v2/message"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/pact-foundation/pact-go/v2/provider"
)

func TestV3MessageProvider(t *testing.T) {
	log.SetLogLevel("TRACE")
	var user *User

	verifier := provider.NewVerifier()

	// Map test descriptions to message producer (handlers)
	functionMappings := message.Handlers{
		"a user event": func([]models.ProviderState) (message.Body, message.Metadata, error) {
			if user != nil {
				return user, message.Metadata{
					"Content-Type": "application/json",
				}, nil
			} else {
				return models.ProviderStateResponse{
					"message": "not found",
				}, nil, nil
			}
		},
	}

	// Setup any required states for the handlers
	stateMappings := models.StateHandlers{
		"User with id 127 exists": func(setup bool, s models.ProviderState) (models.ProviderStateResponse, error) {
			if setup {
				user = &User{
					ID:       127,
					Name:     "Billy",
					Date:     "2020-01-01",
					LastName: "Sampson",
				}
			}

			return models.ProviderStateResponse{"id": user.ID}, nil
		},
	}

	// Verify the Provider with pact contract from PactFile broker
	verifier.VerifyProvider(t, provider.VerifyRequest{
		StateHandlers:   stateMappings,
		Provider:        "V3MessageProvider",
		ProviderVersion: "1.0.0",
		BrokerURL:       "https://webflowluca.pactflow.io",
		BrokerToken:     os.Getenv("PACTFLOW_TOKEN"),
		MessageHandlers: functionMappings,
	})

}

type User struct {
	ID       int    `json:"id" pact:"example=27"`
	Name     string `json:"name" pact:"example=billy"`
	LastName string `json:"lastName" pact:"example=Sampson"`
	Date     string `json:"datetime" pact:"example=2020-01-01'T'08:00:45,format=yyyy-MM-dd'T'HH:mm:ss,generator=datetime"`
}
