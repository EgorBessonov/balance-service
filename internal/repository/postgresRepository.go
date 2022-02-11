//Package repository applies for connection with databases
package repository

import (
	"context"
	"fmt"
	"github.com/EgorBessonov/balance-service/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

//PostgresRepository type
type PostgresRepository struct {
	DBconn *pgxpool.Pool
}

//NewPostgresRepository returns new repository instance
func NewPostgresRepository(conn *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		DBconn: conn,
	}
}

//Get method returns user balance from repository
func (rps *PostgresRepository) Get(ctx context.Context, userID string) (*model.User, error) {
	user := model.User{ID: userID}
	err := rps.DBconn.QueryRow(ctx, `select balance from users where id=$1`, userID).Scan(&user.Balance)
	if err != nil {
		return nil, fmt.Errorf("postgres repository: can't get user balance - %e", err)
	}
	return &user, nil
}

//Update method send new balance value to repository
func (rps *PostgresRepository) Update(ctx context.Context, user *model.User) error {
	_, err := rps.DBconn.Exec(ctx, `update users
	set balance=$1 where id=$2`, user.Balance, user.ID)
	if err != nil {
		return fmt.Errorf("postgres repository: can't update balance - %e", err)
	}
	return nil
}
