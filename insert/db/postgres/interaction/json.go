package interaction

import "time"

type Post struct {
	Name string `json:"name"`
}

type Item struct {
	Id          int64       `json:"id"`
	CampaignId  int64       `json:"campaignId"`
	Name        string      `json:"name"`
	Description interface{} `json:"description"`
	Priority    string      `json:"priority"`
	Removed     bool        `json:"removed"`
	CreatedAt   time.Time   `json:"createdAt"`
}

type Update struct {
	Name        string      `json:"name"`
	Description interface{} `json:"description"`
}

type Delete struct {
	Id         int64 `json:"description"`
	CampaignId int64 `json:"campaignId"`
	Removed    bool  `json:"removed"`
}
