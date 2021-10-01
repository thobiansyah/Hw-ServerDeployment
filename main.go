package main

import (
	"belajar-golang-httprouter/router"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// data := service.Getdata()
	// user := map[string]model.User{}
	// json.Unmarshal(data, &user)

	// fmt.Println(string(data))

	app := fiber.New()

	app.Get("/followers/:username", router.GetFollowerByUsername)
	app.Get("/:userid/detail", router.GetUserDetail)

	app.Listen(":8080")
}
