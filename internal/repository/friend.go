package repository

import (
	"context"
	"time"

	"github.com/MeiramSh/dostap/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FriendRepository struct {
	db *pgxpool.Pool
}

func NewFriendRepository(db *pgxpool.Pool) models.FriendRepository {
	return &FriendRepository{db: db}
}

func (fr *FriendRepository) CreateFriendRequest(c context.Context, req *models.FriendRequest) (int, error) {
	var requestId int
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	q := `INSERT INTO friend_requests (requestor_id,recipient_id,status,created_at) VALUES ($1,$2,$3,$4) RETURNING id`
	err := fr.db.QueryRow(c, q, req.RequestorId, req.RecipientId, "sended", currentTime).Scan(&requestId)
	if err != nil {
		return 0, err
	}
	return requestId, nil
}
