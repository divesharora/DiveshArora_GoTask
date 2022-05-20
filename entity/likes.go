package entity


type Likes struct {
Id        int `json:"id"`
Who_likes int `json:"who_likes"`
Who_is_liked int `json:"who_is_liked"`
}