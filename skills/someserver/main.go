package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type MeetingRoom struct {
	Name      string `json:"name"`
	Capacity  int    `json:"capacity"`
	Status    int    `json:"status"`
	ID        int    `json:"id"`
	AvatarURL string `json:"avatar_url"`
	IsDelete  int    `json:"is_delete"`
}

func main() {
	r := fiber.New()
	r.Get("/api/meetingRoomList", GetMeetingRoomList)
	r.Listen(":9001")
}

func GetMeetingRoomList(c *fiber.Ctx) error {
	page := c.Params("page")
	page_size := c.Params("page_siez")
	fmt.Println(page)
	fmt.Println(page_size)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"meeting_room_list": []MeetingRoom{
			{Name: "修改后的会议室 1", Capacity: 30, Status: 1, ID: 1, AvatarURL: "", IsDelete: 0},
		},
		"meeting_room_count": 1,
	})
}
