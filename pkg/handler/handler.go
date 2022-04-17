package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/suvam720/api/pkg/utils"
)

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Status string `json:"status"`
}

type response struct {
	data User
}

func Handler(ctx context.Context, UserId string) error {
	url := "https://gorest.co.in/public/v2/users/" + UserId
	var data User
	if byte, err := utils.Get(ctx, url); err == nil {
		json.Unmarshal(byte, &data)
		if err != nil {
			return err
		}

		if err := writeToFile(response{
			data: data,
		}); err != nil {
			return err
		}

	} else {
		return err
	}

	return nil
}

//function to write output in a text file
func writeToFile(response response) error {
	filename := response.data.Name
	if filename == "" {
		return errors.New("invalid userid")
	}
	f, err := os.Create(filename + ".txt")
	if err != nil {
		return err
	}
	defer f.Close()

	s := fmt.Sprintf("DETAILS:: Id:%d,\nUserName:%s,\nEmail:%s, \nGender:%s,\nStatus:%s",
		response.data.Id, response.data.Name, response.data.Email,
		response.data.Gender, response.data.Status)

	f.WriteString(s)
	fmt.Println("stored in " + filename + ".txt")
	return nil
}
