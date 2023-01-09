package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Follower struct {
	RelationshipsFollowers []struct {
		StringListData []struct {
			Href      string `json:"href"`
			Value     string `json:"value"`
			Timestamp int    `json:"timestamp"`
		} `json:"string_list_data"`
	} `json:"relationships_followers"`
}

type Following struct {
	RelationshipsFollowers []struct {
		StringListData []struct {
			Href      string `json:"href"`
			Value     string `json:"value"`
			Timestamp int    `json:"timestamp"`
		} `json:"string_list_data"`
	} `json:"relationships_following"`
}

func main() {
	set := make(map[string]bool)

	//followers
	followers, err := os.Open("followers.json")

	if err != nil {
		fmt.Println(err)
	}
	defer followers.Close()

	byteValue, err := ioutil.ReadAll(followers)
	if err != nil {
		fmt.Print(err)
	}

	var data Follower
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Print(err)
	}

	for _, follower := range data.RelationshipsFollowers {
		for _, stringData := range follower.StringListData {
			set[stringData.Value] = true
		}
	}

	//following
	following, err := os.Open("following.json")

	if err != nil {
		fmt.Println(err)
	}
	defer following.Close()

	byteValue, err = ioutil.ReadAll(following)
	if err != nil {
		fmt.Print(err)
	}

	var followingData Following
	err = json.Unmarshal(byteValue, &followingData)
	if err != nil {
		fmt.Print(err)
	}

	for _, follower := range followingData.RelationshipsFollowers {
		for _, stringData := range follower.StringListData {
			_, ok := set[stringData.Value]

			if !ok {
				fmt.Println(stringData.Value)
			}
		}
	}

}
