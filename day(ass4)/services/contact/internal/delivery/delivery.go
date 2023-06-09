package delivery

import (
	"Projects/day/services/contact/internal/useCase"
	"net/http"
)

type UserDelivery struct {
	userUseCase useCase.UserUseCase
}

func NewUserDelivery(userUseCase *useCase.UserUseCase) *UserDelivery {
	return &UserDelivery{}
}
func (d *UserDelivery) HandleUsersRequest(w http.ResponseWriter, r *http.Request) {

}
func (d *UserDelivery) HandleUserRequest(w http.ResponseWriter, r *http.Request) {

}
