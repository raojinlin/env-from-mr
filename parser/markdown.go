package parser

import (
	"regexp"
	"strings"
)

func parseEnv(codeBlock string) map[string]string {
	if codeBlock == "" {
		return nil
	}

	envs := make(map[string]string)
	for _, line := range strings.Split(codeBlock, "\n") {
		if line != "" {
			envSlice := strings.SplitN(line, "=", 2)
			if len(envSlice) > 1 {
				envs[envSlice[0]] = envSlice[1]
			}
		}
	}

	return envs
}

func Parse(doc string) map[string]string {
	codeBlockRegex := regexp.MustCompile("```env\\n([\\s\\S]*?)```")

	codeBlocks := codeBlockRegex.FindAllStringSubmatch(doc, -1)
	envs := make(map[string]string)
	for _, block := range codeBlocks {
		for k, v := range parseEnv(block[1]) {
			envs[k] = v
		}
	}

	return envs
}
