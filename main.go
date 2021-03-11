package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var placeholder string
	var prefix string

	flag.StringVar(&placeholder, "placeholder", "__ENV__", "Placeholder in the target")
	flag.StringVar(&prefix, "prefix", "REACT_APP", "Prefix of the env vars to inject")
	flag.Usage = func() {
		fmt.Printf("Usage\n")
		fmt.Printf("  $ shipator [options] target\n")
		fmt.Printf("\n")
		fmt.Printf("Options\n")
		flag.PrintDefaults()
		fmt.Printf("\n")
		fmt.Printf("Examples\n")
		fmt.Printf("  $ shipator build/index.html\n")
		fmt.Printf("  $ shipator -prefix REACT_APP -placeholder __ENV__ build/index.html\n")
		fmt.Printf("  $ shipator -placeholder __VARS__ build/index.html\n")
		fmt.Printf("  $ shipator -prefix VUE_APP build/index.html\n")
	}

	flag.Parse()

	target := flag.Arg(0)
	if target == "" {
		fmt.Printf("Argument Error: target is not defined\n\n")

		flag.Usage()
		os.Exit(1)
	}

	env := map[string]string{}

	for _, kv := range os.Environ() {
		key := strings.Split(kv, "=")[0]

		if strings.HasPrefix(key, prefix) || key == "NODE_ENV" {
			env[key] = os.Getenv(key)
		}
	}

	encodedEnv, err := json.Marshal(env)
	if err != nil {
		panic("Failed to encoded JSON")
	}

	data, err := ioutil.ReadFile(target)

	updatedContent := strings.Replace(string(data), placeholder, string(encodedEnv), -1)

	f, err := os.Create(target)
	defer f.Close()

	f.WriteString(updatedContent)
	f.Sync()

	fmt.Printf("[OK] %s - placeholder %s was replaced by env vars prefixed by %s\n", target, placeholder, prefix)
}
