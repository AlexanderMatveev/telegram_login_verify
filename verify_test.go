package telegram_login_verify

import (
	"testing"
)

func TestVerifyValid(t *testing.T) {
	t.Parallel()
	loginData := &LoginData{
		AuthDate:  1625259833,
		FirstName: "Alexander",
		Id:        128111077,
		LastName:  "Matveev",
		PhotoUrl:  "https://t.me/i/userpic/320/qDWB1ViXSJK5nbarvLY7TjHMBjpy_ymn91EluZ-WofY.jpg",
		Username:  "opt0ut",
		Hash:      "428099efb5916398f093a83930444f0ca35b10a03e5eafa92b347cbe4e953bec",
	}
	token := "1234567890:ThisIsTotallyFakeTokenByYouCanTry"
	if Verify(loginData, token) != nil {
		t.Errorf("Expected hash to be valid")
	}
}

func TestVerifyInValid(t *testing.T) {
	t.Parallel()
	loginData := &LoginData{
		AuthDate:  1625259833,
		FirstName: "Alexander",
		Id:        128111077,
		LastName:  "Matveev",
		PhotoUrl:  "https://t.me/i/userpic/320/qDWB1ViXSJK5nbarvLY7TjHMBjpy_ymn91EluZ-WofY.jpg",
		Username:  "opt0ut",
		Hash:      "428099efb5916398f093a83930444f0ca35b10a03e5eafa92b347cbe4e953bed",
	}
	token := "1234567890:ThisIsTotallyFakeTokenByYouCanTry"
	if Verify(loginData, token) == nil {
		t.Errorf("Expected hash to be invalid")
	}
}
