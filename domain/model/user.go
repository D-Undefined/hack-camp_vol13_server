package model

type User struct {
	Id       string `json:"uid" gorm:"primaryKey;not null"`
	UserName string `json:"user_name"`
	ImageUrl string `json:"image_url"`
	Comment  string `json:"comment"`
	Location string `json:"location"`
	Twitter  string `json:"twitter"`
	Github   string `json:"github"`
	Url      string `json:"url"`
	Follow   int    `json:"follow"`
	Follower int    `json:"follower"`
	Belong   string `json:"belong"`
	Threads  []*Thread
}
