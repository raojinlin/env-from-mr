package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/raojinlin/env-from-mr/parser"
)

var (
	baseUrl        string
	token          string
	projectID      string
	mergeRequestID int
)

func exportEnv(env map[string]string) {
	for k, v := range env {
		fmt.Printf("export %s='%v'\n", k, v)
	}
}

func main() {
	flag.StringVar(&baseUrl, "url", "", "GitLab base url")
	flag.StringVar(&token, "token", "", "GitLab api access token")
	flag.StringVar(&projectID, "pid", "", "GitLab project id")
	flag.IntVar(&mergeRequestID, "mr-iid", 0, "GitLab Merge Request internal ID")

	flag.Parse()

	// 没有知道ID，取环境变量CI_MERGE_REQUEST_IID
	if mergeRequestID == 0 {
		id := os.Getenv("CI_MERGE_REQUEST_IID")
		if id == "" {
			fmt.Println("No merge request id specify.")
			os.Exit(1)
		}

		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println("Parse merge request id from env CI_MERGE_REQUEST_IID failed: ", err.Error())
			os.Exit(2)
		}

		mergeRequestID = int(i)
	}

	mr, err := getMergeRequest(baseUrl, token, projectID, mergeRequestID)
	if err != nil {
		fmt.Println("Get merge request failed: ", err.Error())
		os.Exit(1)
	}

	exportEnv(parser.Parse(mr.Description))
}
