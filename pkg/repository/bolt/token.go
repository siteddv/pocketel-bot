package bolt

import (
	"github.com/boltdb/bolt"
	"github.com/siteddv/pocketel_bot/pkg/repository"
	"strconv"
)

type TokenRepository struct {
	db *bolt.DB
}

func NewTokenRepository(db *bolt.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

func (r *TokenRepository) Save(chatID int64, token string, bucket repository.Bucket) error {
	return r.db.Update(
		func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(bucket))
			return b.Put(intToBytes(chatID), []byte(token))
		},
	)
}

func intToBytes(v int64) []byte {
	return []byte(strconv.FormatInt(v, 10))
}
