package store

type Store interface {
	User() UserRepository
	Token() TokenRepository
	Post() PostRepository
	Comment() CommentRepository
}
