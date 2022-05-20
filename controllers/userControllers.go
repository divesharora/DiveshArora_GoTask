package controllers


	import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"fmt"
	"DiveshArora_GoTask/db"
	"DiveshArora_GoTask/entity"
	"github.com/gorilla/mux"
)


// reads from json file and adds to db
func AddUsers(w http.ResponseWriter, r *http.Request) {
	database.Migrate(&entity.User{})
	byteValues, err := ioutil.ReadFile("users.json")
	if err != nil {
		fmt.Println("ioutil.ReadFile ERROR:", err)
	} else {
		var users []entity.User
		json.Unmarshal(byteValues, &users)
		for _, user := range users {
			database.Connector.Create(user)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Users created successfully")
}

//searches for person name get searchkey from url query
func GetPersonByName(w http.ResponseWriter, r *http.Request) {
	database.Migrate(&entity.User{})
	vars := mux.Vars(r)
	key := r.URL.Query().Get("key")
	fmt.Println( vars["name"])

	var user []entity.User
	database.Connector.Where("name LIKE ?", key).Find(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

//to import likes from json file
func AddLikes(w http.ResponseWriter, r *http.Request) {
	database.MigrateLikes(&entity.Likes{})
	byteValues, err := ioutil.ReadFile("likes.json")
	if err != nil {
		fmt.Println("ioutil.ReadFile ERROR:", err)
	} else {
		var likes []entity.Likes
		json.Unmarshal(byteValues, &likes)
		for _, like := range likes {
			database.Connector.Create(like)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Likes created successfully")
}


//to search for users within a particular distance from the main user
func GetUsersWithinDistance(w http.ResponseWriter, r *http.Request) {
	database.Migrate(&entity.User{})
	userid := r.URL.Query().Get("uid")
	distance , err := strconv.ParseFloat(r.URL.Query().Get("distance"),64)
	if err != nil {
		fmt.Println("Error:", err)
		distance = 0
	}
	var og entity.User
	database.Connector.Where("id=?", userid).Find(&og)
	fmt.Println(og)
	var matches []entity.User
	database.Connector.Where("location BETWEEN ? and ?",og.Location+distance,og.Location-distance).Find(&matches)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(matches)
	json.NewEncoder(w).Encode(matches)
}

// to find all matches for a user
func GetMatches(w http.ResponseWriter, r *http.Request) {
	database.MigrateLikes(&entity.Likes{})
	userid := r.URL.Query().Get("uid")
	var likes []entity.Likes
	database.Connector.Where("who_likes=?", userid).Find(&likes)
	var isLiked []entity.Likes
	database.Connector.Where("who_is_liked=?", userid).Find(&isLiked)
	var matches []entity.Likes
	for _, like := range likes {
		for _, isLike := range isLiked {
			if like.Who_is_liked == isLike.Who_likes {
				matches = append(matches,like)
			}
		}
	}
	database.Migrate(&entity.User{})
	var users []entity.User
	for _, match := range matches {
		var u entity.User
		database.Connector.Where("id=?", match.Who_is_liked).First(&u)
		users = append(users,u)
	}

	
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(matches)
	json.NewEncoder(w).Encode(users)
}
