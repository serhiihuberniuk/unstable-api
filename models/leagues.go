package models

type Leagues struct {
	LeagueID int    `json:"leagueid"`
	Ticket   string `json:"ticket"`
	Banner   string `json:"banner"`
	Tier     string `json:"tier"`
	Name     string `json:"name"`
}
