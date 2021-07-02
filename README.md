# Telegram Login Widget Verify

This is a simple module to help validate the [Telegram Login Widget](https://core.telegram.org/widgets/login)
authentication response.

## Example Usage

```go
package main

import (
	"github.com/AlexanderMatveev/telegram_login_verify"
	"log"
)

func main() {
	token := "1234567890:ThisIsTotallyFakeTokenByYouCanTry"
	loginData := &telegram_login_verify.LoginData{
		AuthDate:  1625259833,
		FirstName: "John",
		Id:        1122334455,
		LastName:  "Doe",
		PhotoUrl:  "https://t.me/i/userpic/320/qDWB1ViXSJK5nbarvLY7TjHMBjpy_ymn91EluZ-WofY.jpg",
		Username:  "opt0ut",
		Hash:      "428099efb5916398f093a83930444f0ca35b10a03e5eafa92b347cbe4e953bec",
	}
	if err := telegram_login_verify.Verify(loginData, token); err != nil {
		log.Fatalf("Telegram verify error: %v", err)
	}
	log.Print("Validation success")
}
```