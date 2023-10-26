package parser

import (
	"os"
	"testing"
)

func TestParseFromMarkdown(t *testing.T) {
	doc, _ := os.ReadFile("./test.md")
	env := Parse(string(doc))

	if _, ok := env["SVC_A_URL"]; !ok {
		t.Fatal("env SVC_A_URL not found")
	}

	if _, ok := env["SVC_B_URL"]; !ok {
		t.Fatal("env SVC_B_URL not found")
	}

	if _, ok := env["SVC_C_URL"]; !ok {
		t.Fatal("env SVC_C_URL not found")
	}

	if _, ok := env["SVC_D_URL"]; ok {
		t.Fatal("env SVC_D_URL found")
	}

	// exportEnv(env)
}
