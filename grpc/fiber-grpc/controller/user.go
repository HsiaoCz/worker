package controller

import (
	"context"
	"net/http"
	"strings"
	"time"

	v1 "github.com/HsiaoCz/worker/grpc/fiber-grpc/internal/v1"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserSignup struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
	Email      string `json:"email"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var (
	userServiceAddr = "127.0.0.1:9001"
)

func HandleUserSignup(c *fiber.Ctx) error {
	var usersp UserSignup
	if err := c.BodyParser(usersp); err != nil {
		return err
	}
	if len(usersp.Username) == 0 || len(usersp.Password) == 0 || len(usersp.RePassword) == 0 || len(usersp.Email) == 0 {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": "请输入用户信息",
		})
	}
	if len(usersp.Username) > 20 {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": "用户名过长,请输入不超过20个字符",
		})
	}
	if len(usersp.Password) < 8 || len(usersp.Password) > 20 {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": "密码长度需要在8-20个字符字符之间",
		})
	}
	if usersp.Password != usersp.RePassword {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": "请确认密码",
		})
	}
	if strings.Contains(usersp.Email, "@") {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": "请输入正确的邮件",
		})
	}
	conn, err := grpc.Dial(userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx, canecl := context.WithTimeout(context.Background(), time.Millisecond*1500)
	defer canecl()

	client := v1.NewFiberServiceClient(conn)
	signupResponse, err := client.Signup(ctx, &v1.SignupRequest{Username: usersp.Username, Password: usersp.Password, Email: usersp.Email})
	if err != nil {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"StatusCode": signupResponse.GetStatusCode(),
			"Message":    signupResponse.GetMsg(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"StatusCode": signupResponse.GetStatusCode(),
		"Message":    signupResponse.GetMsg(),
	})
}

func HandleUserLogin(c *fiber.Ctx) error {
	return nil
}
