package entites

import "github.com/pborman/uuid"

type Tweet struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

func NewTweet() *Tweet {
	tweet := Tweet{
		Id: uuid.New(),
	}

	return &tweet
}
