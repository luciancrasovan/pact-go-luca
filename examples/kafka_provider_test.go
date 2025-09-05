package main

import (
	"github.com/pact-foundation/pact-go/v2/message"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/pact-foundation/pact-go/v2/provider"
	"os"
	"testing"
)

func TestKafkaProvider(t *testing.T) {
	verifier := provider.NewVerifier()
	var genre *Genre

	// Map test descriptions to message producer (handlers)
	functionMappings := message.Handlers{
		"a gebre event": func([]models.ProviderState) (message.Body, message.Metadata, error) {
			if genre != nil {
				return genre, message.Metadata{
					"Content-Type": "application/json",
				}, nil
			} else {
				return models.ProviderStateResponse{
					"message": "not found",
				}, nil, nil
			}
		},
	}

	stateMappings := models.StateHandlers{
		"Genre with id 1000 exists": func(setup bool, s models.ProviderState) (models.ProviderStateResponse, error) {
			if setup {
				genre = &Genre{
					ID:   100,
					Name: "SciFi",
				}
			}

			return models.ProviderStateResponse{"id": genre.ID}, nil
		},
	}

	verifyRequest := provider.VerifyRequest{
		Provider:                   "GenreEventProducer",
		BrokerURL:                  "https://webflowluca.pactflow.io",
		BrokerToken:                os.Getenv("PACTFLOW_TOKEN"),
		PublishVerificationResults: true,
		ProviderVersion:            "1.0.0",
		Transports: []provider.Transport{
			{
				Protocol: "message",
			},
		},
		MessageHandlers: functionMappings,
		StateHandlers:   stateMappings,
	}

	if err := verifier.VerifyProvider(t, verifyRequest); err != nil {
		t.Fatalf("Error verificando el provider: %v", err)
	}
}

type Genre struct {
	ID   int    `json:"id" pact:"example=1000"`
	Name string `json:"name" pact:"example=SciFi"`
}
