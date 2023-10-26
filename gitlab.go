package main

import "github.com/xanzy/go-gitlab"

func getMergeRequest(baseUrl, token, projectID string, mergeRequest int) (*gitlab.MergeRequest, error) {
	gitlabClient, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseUrl))
	if err != nil {
		return nil, err
	}

	mr, _, err := gitlabClient.MergeRequests.GetMergeRequest(projectID, mergeRequest, &gitlab.GetMergeRequestsOptions{})
	return mr, err
}
