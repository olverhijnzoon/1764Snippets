package main

import (
	"fmt"
	"os"
	"strconv"

	vault "github.com/hashicorp/vault/api"
)

func main() {

	fmt.Println("# 42Snippets")
	fmt.Println("## Golang Hashicorp Vault Client")

	// Get the Vault token from log - I know
	vaultToken := "s.H4zL3fJwBe2AuebW6GkLmUUf"

	// Create a new Vault client
	config := &vault.Config{
		Address: "http://localhost:31082",
	}

	client, err := vault.NewClient(config)
	if err != nil {
		fmt.Println("Error creating Vault client: ", err)
		os.Exit(1)
	}

	// Set the Vault token
	client.SetToken(vaultToken)

	// Mount a new KV v2 secrets engine at the path "1764Secrets"
	err = client.Sys().Mount("1764Secrets", &vault.MountInput{
		Type:        "kv-v2",
		Description: "My KV v2 secrets engine",
	})
	if err != nil {
		fmt.Println("Error mounting secrets engine: ", err)
		os.Exit(1)
	}

	fmt.Println("Successfully mounted KV v2 secrets engine at path '1764Secrets'")

	// Create 42 secrets in the "1764Secrets" secrets engine
	for i := 1; i <= 42; i++ {
		_, err = client.Logical().Write("1764Secrets/data/secret"+strconv.Itoa(i), map[string]interface{}{
			"data": map[string]interface{}{
				"value": "This is secret " + strconv.Itoa(i),
			},
		})
		if err != nil {
			fmt.Println("Error writing secret: ", err)
			os.Exit(1)
		}
	}

	fmt.Println("Successfully created 42 secrets in the '1764Secrets' secrets engine")

}
