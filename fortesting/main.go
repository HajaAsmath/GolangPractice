package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"golang.org/x/oauth2"
)

func main() {
	config := oauth1.NewConfig("FlrIzwB1IKzufwVuemhe5owrO", "LrkWJqZ8f7LWC0kkdlRsJ9uOmeEydoXBMVdWYKzGLTzR2Ho5Ld")
	token := oauth1.NewToken("2679269617-YGN7FxfABgUp3v62hSf4ruGeqBY5Dhki1GUlcge", "z1leik9r9r4GcxduKFBmZS543nBfPggIMRKxQC2yKelST")
	httpClient := config.Client(oauth2.NoContext, token)

	client := twitter.NewClient(httpClient)

	search := client.Search
	params := &twitter.SearchTweetParams{
		Query: "SuperStar",
	}
	s, _, err := search.Tweets(params)
	if err != nil {
		fmt.Println("Error :", err)
	}
	tweet := s.Statuses
	metadata := s.Metadata
	fmt.Println(metadata.Query)
	fmt.Println(tweet[0].CreatedAtTime())
	fmt.Println(tweet[0].User.ScreenName)
	fmt.Println(tweet[0].User.Name)
	fmt.Println(tweet[0].RetweetCount)
	fmt.Println(tweet[0].Text)

	tweets, _, err := client.Statuses.Update("Hello from Golang", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tweets.User.Email)
}
