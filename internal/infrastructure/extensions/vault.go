package extensions

import (
	"context"
	"github.com/hashicorp/vault/api"
	"log"
)

type VaultClient struct {
	client *api.Client
}

func NewVaultClient(address, token string) *VaultClient {
	config := api.DefaultConfig()
	config.Address = address

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Error creando cliente Vault: %v", err)
	}

	client.SetToken(token)

	return &VaultClient{client: client}
}

func (v *VaultClient) GetSecret(path string) (map[string]interface{}, error) {
	secret, err := v.client.KVv2("secret").Get(context.Background(), path)
	if err != nil {
		return nil, err
	}
	return secret.Data, nil
}
