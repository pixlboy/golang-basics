// HTTP and JSON

package main

import ( 
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

// User is a github user info
type User struct {
	Name string `json:"name"`
	PublicRepos int `json:"public_repos"`
}

func userInfo(login string) (*User, error){
	//Make HTTP call

	url := fmt.Sprintf("https://api.github.com/users/%s", login)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	//Decode JSON

	user := &User{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

func main() {

	fmt.Printf("---------- \n\n")

	user, err := userInfo("rach8garg")
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	fmt.Printf("%+v \n", user)


	fmt.Println("\n\n ----------")

}