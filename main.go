package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	target, ok := os.LookupEnv("SHIPATOR_TARGET")

	if !ok {
		panic("Env var SHIPATOR_TARGET not defined")
	}

	placeholder, ok := os.LookupEnv("SHIPATOR_PLACEHOLDER")

	if !ok {
		panic("Env var SHIPATOR_PLACEHOLDER not defined")
	}

	fmt.Println("[shipator] 1-4: Reading env vars")

	env := map[string]string{}

	for _, kv := range os.Environ() {
		key := strings.Split(kv, "=")[0]

		if strings.HasPrefix(key, "REACT_APP") || key == "NODE_ENV" {
			env[key] = os.Getenv(key)
		}
	}

	encodedEnv, err := json.Marshal(env)
	if err != nil {
		panic("Failed to encoded JSON")
	}

	fmt.Println("[shipator] 2-4: Reading template")

	data, err := ioutil.ReadFile(target)

	fmt.Println("[shipator] 3-4: Injecting env vars")
	updatedContent := strings.Replace(string(data), placeholder, string(encodedEnv), -1)

	f, err := os.Create(target)
	defer f.Close()

	fmt.Println("[shipator] 4-4: Writing template")
	f.WriteString(updatedContent)
	f.Sync()
}
