package redis

import "context"

func (s *Store) CreateUser(phone, password string, ctx context.Context) error {
	err := s.redis.Set(ctx, phone, password, 0).Err()
	if err != nil {
		return err
	}

	return err
}
