package helper

import (
	"Coattails"
)

type (
	Api struct {
		P *Coattails.Prophet
	}
	summonerNameReq struct {
		SummonerName string `json:"summonerName"`
	}

	CurrSummoner struct {
		AccountId                   int64  `json:"accountId"`
		DisplayName                 string `json:"displayName"`
		InternalName                string `json:"internalName"`
		NameChangeFlag              bool   `json:"nameChangeFlag"`
		PercentCompleteForNextLevel int    `json:"percentCompleteForNextLevel"`
		ProfileIconId               int    `json:"profileIconId"`
		Puuid                       string `json:"puuid"`
		RerollPoints                struct {
			CurrentPoints    int `json:"currentPoints"`
			MaxRolls         int `json:"maxRolls"`
			NumberOfRolls    int `json:"numberOfRolls"`
			PointsCostToRoll int `json:"pointsCostToRoll"`
			PointsToReroll   int `json:"pointsToReroll"`
		} `json:"rerollPoints"`
		SummonerId       int64 `json:"summonerId"`
		SummonerLevel    int   `json:"summonerLevel"`
		Unnamed          bool  `json:"unnamed"`
		XpSinceLastLevel int   `json:"xpSinceLastLevel"`
		XpUntilNextLevel int   `json:"xpUntilNextLevel"`
	}
)
