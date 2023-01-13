package main

import (
	"fmt"
	"sync"
)

var VideoList = []Video{{Id: 1, Title: "go basics 1", RequiredSubType: "basic", Duration: 10},
	{Id: 1, Title: "go basics 2", RequiredSubType: "standard", Duration: 10},
	{Id: 1, Title: "go basics3", RequiredSubType: "premium", Duration: 10},
	{Id: 1, Title: "go basics 4", RequiredSubType: "premium", Duration: 10},
	{Id: 1, Title: "go basics5", RequiredSubType: "basic", Duration: 10},
	{Id: 1, Title: "go basics 6", RequiredSubType: "standard", Duration: 10}}

var Wg sync.WaitGroup

type Video struct {
	Id              int
	Title           string
	RequiredSubType string
	Duration        int
}

func GetVideos(RequiredSubscriptionType string) {
	defer Wg.Done()
	var TmpVideoList []Video
	for _, video := range VideoList {
		if video.RequiredSubType == RequiredSubscriptionType {
			TmpVideoList = append(TmpVideoList, video)
		}
	}
	fmt.Printf("\n list of videos for %v are %+v \n", RequiredSubscriptionType, TmpVideoList)
}

func main() {
	// increase the counter by 1 (any value we can pass)
	Wg.Add(1)
	go GetVideos("basic")
	Wg.Add(1)
	go GetVideos("premium")
	Wg.Add(1)

	go GetVideos("standard")

	Wg.Wait()

}
