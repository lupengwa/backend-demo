package demo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Service interface {
	SearchMovie(name string) (MovieResp, error)
}

type ServiceImpl struct {
}

func NewService() *ServiceImpl {
	return &ServiceImpl{}
}

func (service *ServiceImpl) SearchMovie(name string) (MovieResp, error) {
	var result MovieResp
	req, err := http.NewRequest("GET", "https://api.themoviedb.org/3/search/movie", nil)
	if err != nil {
		log.Print(err)
		return result, err
	}

	q := req.URL.Query()
	q.Add("api_key", "a624463b70205f0feb942d449e888d4b")
	q.Add("query", name)
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return result, err
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)
	fmt.Println(req.URL.String())
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return result, err
	}
	// Unmarshal the response body into the SearchResult struct
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		fmt.Println("Error:", err)
		return result, err
	}
	fmt.Println(result)

	result.Results = result.Results[0:10]

	return result, nil
}

//
//func (service *ServiceImpl) AddNewUser(user UserDto) (UserDto, error) {
//	userSavedEntity, err := service.demoRepo.SaveUser(user)
//	var userSavedDto UserDto
//	if err == nil {
//		userSavedDto = UserDto{Email: userSavedEntity.Email, Id: userSavedEntity.Id}
//		return userSavedDto, nil
//	}
//	return user, err
//}
