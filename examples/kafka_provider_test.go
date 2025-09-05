package main

import (
	"context"
	"fmt"
	"github.com/pact-foundation/pact-go/v2/provider"
	"log"
	"os"
)

func main() {
	verifier := provider.NewVerifier()

	// Configuración del verificador
	verifyRequest := provider.VerifyRequest{
		Provider:                   "GenreEventProducer",
		BrokerURL:                  "https://webflowluca.pactflow.io",
		BrokerToken:                os.Getenv("PACTFLOW_TOKEN"), // Usa variable de entorno para seguridad
		PublishVerificationResults: true,
		ProviderVersion:            "1.0.0",
		Transports: []provider.Transport{
			{
				Protocol: "message",
			},
		},
		MessageHandlers: map[string]func() (interface{}, error){
			"genre message": func() (interface{}, error) {
				// Simula el mensaje que se publicaría en Kafka
				message := map[string]interface{}{
					"id":    123,
					"nobre": "Rock",
				}
				return message, nil
			},
		},
	}

	// Ejecuta la verificación
	if err := verifier.VerifyProvider(context.Background(), verifyRequest); err != nil {
		log.Fatalf("Error verificando el provider: %v", err)
	}

	fmt.Println("✅ Verificación completada correctamente")
}
