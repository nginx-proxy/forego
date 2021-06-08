package main

import "testing"

func TestMultipleEnvironmentFiles(t *testing.T) {
	envs := []string{"fixtures/envs/.env1", "fixtures/envs/.env2"}
	env, err := loadEnvs(envs)

	if err != nil {
		t.Fatalf("Could not read environments: %s", err)
	}

	if env.Get("env1") == "" {
		t.Fatal("$env1 should be present and is not")
	}

	if env.Get("env2") == "" {
		t.Fatal("$env2 should be present and is not")
	}
}
