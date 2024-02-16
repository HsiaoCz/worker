package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/HsiaoCz/worker/moo/endpoints"
	"github.com/HsiaoCz/worker/moo/internal/pb"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Code":    http.StatusBadRequest,
			"Message": err.Error(),
		})
	}
	conn, err := grpc.Dial(endpoints.MooAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewMooClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1500)
	defer cancel()
	resp, err := client.Login(ctx, &pb.LoginRequest{Username: user.Username, Password: user.Password})
	if err != nil {
		return err
	}
	return c.Status(int(resp.GetCode())).JSON(fiber.Map{
		"Code":    resp.GetCode(),
		"Message": resp.GetMsg(),
	})
}
