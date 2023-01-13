package main

import "fmt"

var VideoList = []Video{{Id: 1, Title: "go basics 1", RequiredSubType: "basic", Duration: 10},
	{Id: 1, Title: "go basics 2", RequiredSubType: "standard", Duration: 10},
	{Id: 1, Title: "go basics3", RequiredSubType: "premium", Duration: 10},
	{Id: 1, Title: "go basics 4", RequiredSubType: "premium", Duration: 10},
	{Id: 1, Title: "go basics5", RequiredSubType: "basic", Duration: 10},
	{Id: 1, Title: "go basics 6", RequiredSubType: "standard", Duration: 10}}

type Video struct {
	Id              int
	Title           string
	RequiredSubType string
	Duration        int
}

func GetVideos(RequiredSubscriptionType string) {
	var TmpVideoList []Video
	for _, video := range VideoList {
		if video.RequiredSubType == RequiredSubscriptionType {
			TmpVideoList = append(TmpVideoList, video)
		}
	}
	fmt.Println("--------------------------------------------------------------------------------------------------------")

	fmt.Printf("list of videos for %v are %+v", RequiredSubscriptionType, TmpVideoList)
	fmt.Println("")
}

func main() {
	// it will print result of below 3 calls as we are using general control flow for execution
	GetVideos("basic")
	// GetVideos("standard")
	// GetVideos("premium")
	// Here we will miss out the result of the function call that was made using go routine
	go GetVideos("basic")
	GetVideos("standard")

}
