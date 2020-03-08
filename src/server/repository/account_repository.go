package repository

var AccountRepository = newAccountRepository()

// new
func newAccountRepository() *accountRepository {
	AccountRepository := &accountRepository{}
	 return AccountRepository
}

type accountRepository struct {
	BaseRepository
}

