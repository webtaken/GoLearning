package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"google.golang.org/api/youtube/v3"
)

func handleModelResponseSequential(
	sentimentAnalysisResponse []RobertaResponse,
	positive *Counter,
	neutral *Counter,
	negative *Counter,
) {
	for _, analysisResponse := range sentimentAnalysisResponse {
		maxScoreObj := struct {
			label string
			score float64
		}{
			label: analysisResponse[0].Label,
			score: analysisResponse[0].Score,
		}
		for _, analysis := range analysisResponse {
			if analysis.Score > maxScoreObj.score {
				maxScoreObj.label = analysis.Label
				maxScoreObj.score = analysis.Score
			}
		}

		switch maxScoreObj.label {
		case "Positive":
			positive.Count += 1
		case "Negative":
			negative.Count += 1
		case "Neutral":
			neutral.Count += 1
		}
	}
}

func batchAnalyzerSequential(
	comments []*youtube.CommentThread,
	positive *Counter,
	neutral *Counter,
	negative *Counter,
	errors *Counter,
) {
	modelEndpoint := "https://api-inference.huggingface.co/models/cardiffnlp/twitter-xlm-roberta-base-sentiment"
	var jsonData = []byte(buildRequestBody(comments))
	request, err := http.NewRequest("POST", modelEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		errors.Count += len(comments)
		return
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", huggingFaceAPIToken))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		errors.Count += len(comments)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, _ := io.ReadAll(response.Body)
		var sentimentAnalysisResponse []RobertaResponse

		err := json.Unmarshal(body, &sentimentAnalysisResponse)
		if err != nil {
			errors.Count += len(comments)
			return
		}
		handleModelResponseSequential(sentimentAnalysisResponse, positive, negative, neutral)
		return
	}
	errors.Count += len(comments)
}

func analyzeVideoSequential(videoId string, numComments int) {
	positive := Counter{
		Tag:     "Positive",
		Count:   0,
		Channel: make(chan int),
	}
	negative := Counter{
		Tag:     "Negative",
		Count:   0,
		Channel: make(chan int),
	}
	neutral := Counter{
		Tag:     "Neutral",
		Count:   0,
		Channel: make(chan int),
	}
	errors := Counter{
		Tag:     "Errors",
		Count:   0,
		Channel: make(chan int),
	}

	var part = []string{"id", "snippet"}
	nextPageToken := ""
	call := Service.CommentThreads.List(part)
	call.VideoId(videoId)
	call.MaxResults(int64(BATCH))
	commentsRetrieved := 0
	commentsRetrieveFailed := 0
	for commentsRetrieved < numComments {
		if nextPageToken != "" {
			call.PageToken(nextPageToken)
		}

		response, err := call.Do()
		if err != nil {
			commentsRetrieveFailed += BATCH
			continue
		}

		commentsToAnalyze := len(response.Items)
		if commentsRetrieved+commentsToAnalyze >= numComments {
			commentsToAnalyze = numComments - commentsRetrieved
		}
		commentsRetrieved += commentsToAnalyze

		batchAnalyzerSequential(response.Items[:commentsToAnalyze], &positive, &negative, &neutral, &errors)

		nextPageToken = response.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	// After all the process we show the results
	printAnalysisResults(&positive, &neutral, &negative, &errors, numComments)
}
