package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// Batch size for
const BATCH = 10
const WORKERS = 200

var youtubeAPIToken = os.Getenv("YOUTUBE_API_TOKEN")
var huggingFaceAPIToken = os.Getenv("HUGGING_FACE_API_TOKEN")
var Service *youtube.Service

type RobertaResponse []struct {
	Label string  `json:"label"`
	Score float64 `json:"score"`
}

type Counter struct {
	Tag     string
	Count   int
	Channel chan int
}

func init() {
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(youtubeAPIToken))
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	Service = service
}

func getVideoData(videoId string) (*youtube.VideoListResponse, error) {
	var part = []string{"snippet", "contentDetails", "statistics"}
	call := Service.Videos.List(part)
	call.Id(videoId)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}
	if len(response.Items) == 0 {
		return nil, fmt.Errorf("video not found")
	}
	return response, nil
}

func buildRequestBody(comments []*youtube.CommentThread) string {
	listFormatted := "["
	for i, comment := range comments {
		listFormatted += fmt.Sprintf("%q", comment.Snippet.TopLevelComment.Snippet.TextOriginal)
		if i < len(comments)-1 {
			listFormatted += ","
		}
	}
	listFormatted += "]"
	return fmt.Sprintf("{ \"inputs\": %s }", listFormatted)
}

func printYoutubeVideoData(videoData *youtube.VideoListResponse) {
	fmt.Printf("Video Data\n")
	fmt.Printf("Title: %s\n", videoData.Items[0].Snippet.Title)
	fmt.Printf("-------------------------------------------------\n")
	fmt.Printf("Description: %s\n", videoData.Items[0].Snippet.Description)
	fmt.Printf("-------------------------------------------------\n")
	fmt.Printf("Likes: %d\n", videoData.Items[0].Statistics.LikeCount)
	fmt.Printf("Comments Count: %d\n", videoData.Items[0].Statistics.CommentCount)
	fmt.Printf("Views Count: %d\n", videoData.Items[0].Statistics.ViewCount)
}

func printAnalysisResults(
	positive *Counter,
	neutral *Counter,
	negative *Counter,
	errors *Counter,
	numComments int,
) {
	fmt.Printf("Here are the sentiment analysis results for your video\n")
	fmt.Printf("From %d comments\n", numComments)
	fmt.Printf("Positive comments: %d\n", positive.Count)
	fmt.Printf("Neutral comments: %d\n", neutral.Count)
	fmt.Printf("Negative comments: %d\n", negative.Count)
	fmt.Printf("Errors found during the analysis: %d\n", errors.Count)
}

func main() {
	var videoId string
	var numComments int
	flag.StringVar(&videoId, "videoId", "", "The id of a youtube video e.g. yyUHQIec83I")
	flag.IntVar(&numComments, "numComments", 100, "The number of comments to analyze (default: 100)")
	flag.Parse()

	if numComments <= 0 {
		log.Fatal("numComments flag must be a positive number")
	}

	videoData, err := getVideoData(videoId)

	if err != nil {
		log.Fatal(err.Error())
	}

	printYoutubeVideoData(videoData)

	fmt.Printf("\nAnalyzing video (Sequential method)...\n")
	startTime := time.Now()
	analyzeVideoSequential(videoId, numComments)
	fmt.Printf("Analysis took %.3f seconds to execute.\n", time.Since(startTime).Seconds())

	fmt.Printf("\nAnalyzing video (Parallel method)...\n")
	startTime = time.Now()
	analyzeVideoParallel(videoId, numComments)
	fmt.Printf("Analysis took %.3f seconds to execute.\n", time.Since(startTime).Seconds())
}
