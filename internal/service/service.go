//Package service represents balance service logic
package service

import (
	"context"
	"fmt"
	"github.com/EgorBessonov/balance-service/internal/model"
	"github.com/EgorBessonov/balance-service/internal/repository"
)

//Service struct
type Service struct {
	Repository *repository.PostgresRepository
}

//NewService returns new service instance
func NewService(rps *repository.PostgresRepository) *Service {
	return &Service{
		Repository: rps,
	}
}

//Get method returns user balance
func (service *Service) Get(ctx context.Context, userID string) (float32, error) {
	user, err := service.Repository.Get(ctx, userID)
	if err != nil {
		return 0.0, fmt.Errorf("service: can't get balance - %e", err)
	}
	return user.Balance, nil
}

//Check method checks if user has necessary balance
func (service *Service) Check(ctx context.Context, userID string, requiredBalance float32) (bool, error) {
	userBalance, err := service.Get(ctx, userID)
	if err != nil {
		return false, fmt.Errorf("service: can't check balnce - %e", err)
	}
	return userBalance >= requiredBalance, nil
}

//TopUp method increase user balance
func (service *Service) TopUp(ctx context.Context, userID string, shift float32) error {
	userBalance, err := service.Get(ctx, userID)
	if err != nil {
		return fmt.Errorf("service: can't top up balance - %e", err)
	}
	user := &model.User{
		ID:      userID,
		Balance: userBalance + shift,
	}
	err = service.Repository.Update(ctx, user)
	if err != nil {
		return fmt.Errorf("service: can't top up balance - %e", err)
	}
	return nil
}

//Withdraw method decrease user balance
func (service *Service) Withdraw(ctx context.Context, userID string, shift float32) error {
	userBalance, err := service.Get(ctx, userID)
	if err != nil {
		return fmt.Errorf("service: can't withdraw balance - %e", err)
	}
	user := &model.User{
		ID:      userID,
		Balance: userBalance - shift,
	}
	err = service.Repository.Update(ctx, user)
	if err != nil {
		return fmt.Errorf("service: can't withdraw balance - %e", err)
	}
	return nil
}
