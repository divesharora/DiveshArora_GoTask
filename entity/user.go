package entity


type User struct {
Id        int `json:"id"`
Name string `json:"name"`
Location  float64 `json:"location"`
Gender       string `json:"gender"`
Email 	string `json:"email"`
}