package redis

import (
	"context"
	"fmt"
)

func (s *Store) CheckUser(phone string, ctx context.Context) (string, error) {
	password := s.redis.Get(ctx, phone)
	if password == nil {
		return "", fmt.Errorf("user not found")
	}

	return password.String(), nil
}
