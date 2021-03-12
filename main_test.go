package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("successfully injects env vars", func(t *testing.T) {
		os.Setenv("REACT_APP_SERVER_URL", "https://example.com")

		placeholder := "__ENV__"
		target := tempFile(t, "index.html", placeholder)

		os.Args = []string{"./shipator", target}
		out = bytes.NewBuffer(nil)
		main()

		output := out.(*bytes.Buffer).String()

		assertEqual(t, output, fmt.Sprintf("[info] [shipator] %s - __ENV__ replaced by env vars with prefix REACT_APP\n", target))

		data, _ := ioutil.ReadFile(target)
		assertEqual(t, string(data), "{\"REACT_APP_SERVER_URL\":\"https://example.com\"}")
	})

	t.Run("prefix flag", func(t *testing.T) {
		os.Setenv("VUE_APP_SERVER_URL", "https://example.com")

		placeholder := "__ENV__"
		target := tempFile(t, "index.html", placeholder)

		os.Args = []string{"./shipator", "-prefix", "VUE_APP", target}
		out = bytes.NewBuffer(nil)
		main()

		output := out.(*bytes.Buffer).String()

		assertEqual(t, output, fmt.Sprintf("[info] [shipator] %s - __ENV__ replaced by env vars with prefix VUE_APP\n", target))

		data, _ := ioutil.ReadFile(target)

		assertEqual(t, string(data), "{\"VUE_APP_SERVER_URL\":\"https://example.com\"}")
	})

	t.Run("placeholder flag", func(t *testing.T) {
		os.Setenv("REACT_APP_SERVER_URL", "https://example.com")

		placeholder := "__ANOTHER__"
		target := tempFile(t, "index.html", placeholder)

		os.Args = []string{"./shipator", "-placeholder", placeholder, target}
		out = bytes.NewBuffer(nil)
		main()

		output := out.(*bytes.Buffer).String()

		assertEqual(t, output, fmt.Sprintf("[info] [shipator] %s - __ANOTHER__ replaced by env vars with prefix REACT_APP\n", target))

		data, _ := ioutil.ReadFile(target)
		assertEqual(t, string(data), "{\"REACT_APP_SERVER_URL\":\"https://example.com\"}")
	})
}

func TestInjectEnvVars(t *testing.T) {
	t.Run("successfully injects env vars", func(t *testing.T) {
		os.Setenv("REACT_APP_SERVER_URL", "https://example.com")

		placeholder := "__ENV__"
		target := tempFile(t, "index.html", placeholder)

		injectEnvVars(target, placeholder, "REACT_APP")

		data, _ := ioutil.ReadFile(target)

		assertEqual(t, string(data), "{\"REACT_APP_SERVER_URL\":\"https://example.com\"}")
	})

	t.Run("target does not exists", func(t *testing.T) {
		err := injectEnvVars("not-exists.html", "__ENV__", "REACT_APP")

		assertEqual(t, err.Error(), "open not-exists.html: no such file or directory")
	})

	t.Run("placeholder not found", func(t *testing.T) {
		target := tempFile(t, "index.html", "__ENV__")
		err := injectEnvVars(target, "__ANOTHER__", "REACT_APP")

		assertEqual(t, err.Error(), fmt.Sprintf("placeholder __ANOTHER__ not found in %s", target))
	})
}

func tempFile(t *testing.T, name string, content string) string {
	tempDir := t.TempDir()
	target := path.Join(tempDir, "index.html")

	f, _ := os.Create(target)
	defer f.Close()
	f.WriteString(content)
	f.Sync()

	return target
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}
