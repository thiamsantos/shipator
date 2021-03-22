package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

const version = "0.1.0-rc3"

func main() {
	var placeholder string
	var prefix string
	var versionFlag bool

	cli := flag.NewFlagSet("shipator", flag.ExitOnError)

	cli.SetOutput(out)
	cli.StringVar(&placeholder, "placeholder", "__ENV__", "Placeholder in the target")
	cli.StringVar(&prefix, "prefix", "REACT_APP", "Prefix of the env vars to inject")
	cli.BoolVar(&versionFlag, "version", false, "Prints current version")
	cli.Usage = func() {
		fmt.Fprintf(out, "Usage\n")
		fmt.Fprintf(out, "  $ shipator [options] target\n")
		fmt.Fprintf(out, "\n")
		fmt.Fprintf(out, "Options\n")
		cli.PrintDefaults()
		fmt.Fprintf(out, "\n")
		fmt.Fprintf(out, "Examples\n")
		fmt.Fprintf(out, "  $ shipator build/index.html\n")
		fmt.Fprintf(out, "  $ shipator -prefix REACT_APP -placeholder __ENV__ build/index.html\n")
		fmt.Fprintf(out, "  $ shipator -placeholder __VARS__ build/index.html\n")
		fmt.Fprintf(out, "  $ shipator -prefix VUE_APP build/index.html\n")
	}

	cli.Parse(os.Args[1:])

	if versionFlag {
		fmt.Fprintf(out, "%s\n", version)
		os.Exit(0)
	}

	target := cli.Arg(0)
	if target == "" {
		fmt.Fprintf(out, "Argument Error: target is not defined\n\n")

		cli.Usage()
		os.Exit(1)
	}

	err := injectEnvVars(target, placeholder, prefix)

	if err != nil {
		fmt.Fprint(out, err)
		os.Exit(1)
	}

	fmt.Fprintf(out, "[info] [shipator] %s - %s replaced by env vars with prefix %s\n", target, placeholder, prefix)
}

func injectEnvVars(target string, placeholder string, prefix string) error {
	env := map[string]string{}

	for _, kv := range os.Environ() {
		key := strings.Split(kv, "=")[0]

		if strings.HasPrefix(key, prefix) || key == "NODE_ENV" {
			env[key] = os.Getenv(key)
		}
	}

	encodedEnv, err := json.Marshal(env)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(target)
	if err != nil {
		return err
	}

	content := string(data)

	if !strings.Contains(content, placeholder) {
		return fmt.Errorf("placeholder %s not found in %s", placeholder, target)
	}

	updatedContent := strings.Replace(content, placeholder, string(encodedEnv), -1)

	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(updatedContent)
	f.Sync()

	return nil
}
