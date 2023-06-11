package models

type Site struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Created         int64  `json:"created"`
	CreatedByUserId string `json:"created_by_user_id"`
	OrganizationId  string `json:"organization"`
	Label           string `json:"label"`
	Frozen          bool   `json:"frozen"`
	Locked          bool   `json:"locked"`
}
