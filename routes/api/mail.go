package api

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

type APIHandler struct {
}

func RegisterRoutes(router *gin.Engine) {
	h := &APIHandler{}

	scoreRoutes := router.Group("/api")
	scoreRoutes.POST("/send", h.SendEmail)
}

func (h APIHandler) SendEmail(ctx *gin.Context) {
	from := "tmm6907@gmail.com"
	password := os.Getenv("EMAIL_PASSWORD")

	// Recipient email address
	to := "tmm6907@gmail.com"

	// SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	answer := ctx.PostForm("answer")
	fmt.Println("answer", answer)

	// Message content
	message := []byte(fmt.Sprintf("Subject: Cara's Answer!\n\n She said %s!", answer))

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Connect to the SMTP server
	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpHost, smtpPort), auth, from, []string{to}, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}
	if answer == "Yes" {
		ctx.HTML(http.StatusAccepted, "qrcode", nil)
		return
	} else {
		panic(fmt.Sprintf("She rejected us! %s!!!", answer))
	}
}
