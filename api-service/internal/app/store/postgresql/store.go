package postgresql

import (
	"github.com/Crabocod/gpt_network/api-service/internal/app/store"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db                *sqlx.DB
	userRepository    *UserRepository
	tokenRepository   *TokenRepository
	postRepository    *PostRepository
	commentRepository *CommentRepository
}

func New(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{store: s}
	}
	return s.userRepository
}

func (s *Store) Token() store.TokenRepository {
	if s.tokenRepository == nil {
		s.tokenRepository = &TokenRepository{store: s}
	}
	return s.tokenRepository
}

func (s *Store) Post() store.PostRepository {
	if s.postRepository == nil {
		s.postRepository = &PostRepository{store: s}
	}
	return s.postRepository
}

func (s *Store) Comment() store.CommentRepository {
	if s.commentRepository == nil {
		s.commentRepository = &CommentRepository{store: s}
	}
	return s.commentRepository
}
