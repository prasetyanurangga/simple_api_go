package model

import(
	"time"
)

type (
	//User Is
	User struct{
		ID int `json:"id"`
		Nama string `json:"nama"`
		Gender string `json:"gender"`
		CreateAt time.Time `json:"create_at"`
	}
)