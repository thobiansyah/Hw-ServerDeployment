package router

import (
	"belajar-golang-httprouter/model"
	"belajar-golang-httprouter/service"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetFollowerByUsername(ctx *fiber.Ctx) error {

	data, err := service.Getdata()
	if err != nil {
		return ctx.JSON(err)
	}

	account := map[string]model.User{}
	json.Unmarshal(data, &account)
	var result string
	findUsername := ctx.Params("username")

	for _, value := range account {
		if value.Username == findUsername {
			result = "followers: " + strconv.Itoa(value.Follower)
			break
		}
	}

	return ctx.SendString(result)
}

func GetUserDetail(ctx *fiber.Ctx) error {
	data, err := service.Getdata()
	if err != nil {
		return ctx.JSON(err)
	}
	account := map[string]model.User{}
	json.Unmarshal(data, &account)
	userId := ctx.Params("userid")
	result := account[userId]

	return ctx.JSON(result)
}
