package repository

var TokenRepository = newTokenRepository()

func newTokenRepository() *tokenRepository {
	return &tokenRepository{}
}

type tokenRepository struct {
	BaseRepository
}