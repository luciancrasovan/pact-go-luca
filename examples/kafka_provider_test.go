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
		MessageHandlers: message.Handlers{
			"genre message": func([]models.ProviderState) {
				message := map[string]interface{}{
					"id":    123,
					"nobre": "Rock",
				}
				return message, nil
			},
		},
	}

	if err := verifier.VerifyProvider(t, verifyRequest); err != nil {
		t.Fatalf("Error verificando el provider: %v", err)
	}
}
