package smtp

import (
	"fmt"
	"github.com/spf13/viper"
	"net/smtp"
)

func SendMail(to []string, content string) error {
	from := fmt.Sprintf("%v", viper.Get("MAIL_ADDRESS"))
	password := fmt.Sprintf("%v", viper.Get("MAIL_PASS"))

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "Verification Code"

	message := []byte(subject + " " + content)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		return err
	}

	return nil
}
