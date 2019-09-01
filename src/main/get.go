// This consumes https://raider.io/api#!/character/get_api_v1_characters_profile api

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var character = "deeps"
var fields = "gear"
var region = "us"
var realm = "barthilas"
var url = "https://raider.io/api/v1/characters/profile?"

type raiderioGear struct {
	Name              string `json:"name"`
	Race              string `json:"race"`
	Class             string `json:"class"`
	ActiveSpecName    string `json:"active_spec_name"`
	ActiveSpecRole    string `json:"active_spec_role"`
	Gender            string `json:"gender"`
	Faction           string `json:"faction"`
	AchievementPoints int    `json:"achievement_points"`
	HonorableKills    int    `json:"honorable_kills"`
	ThumbnailURL      string `json:"thumbnail_url"`
	Region            string `json:"region"`
	Realm             string `json:"realm"`
	ProfileURL        string `json:"profile_url"`
	ProfileBanner     string `json:"profile_banner"`
	Gear              struct {
		ItemLevelEquipped int `json:"item_level_equipped"`
		ItemLevelTotal    int `json:"item_level_total"`
		ArtifactTraits    int `json:"artifact_traits"`
	} `json:"gear"`
}

func main() {
	url := url + "region=" + region + "&realm=" + realm + "&name=" + character + "&fields=" + fields
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var gear raiderioGear

	json.NewDecoder(resp.Body).Decode(&gear)

	fmt.Println(gear.Name, "iLVL: ", gear.Gear.ItemLevelEquipped)
}
