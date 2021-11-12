package models

type Team struct {
	TeamID        int     `json:"team_id"`
	Rating        float64 `json:"rating"`
	Wins          int     `json:"wins"`
	Losses        int     `json:"losses"`
	LastMatchTime int     `json:"last_match_time"`
	Name          string  `json:"name"`
	Tag           string  `json:"tag"`
}
