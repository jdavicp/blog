package author

import (
	a "github.com/jdavicp/blog/internals/ports/left/author"
	"github.com/jdavicp/blog/internals/ports/right/database"
)

type adapter struct {
	repo database.AuthorRepositoryPort
}

func New(repo database.AuthorRepositoryPort) *adapter {
	return &adapter{
		repo,
	}
}

func (ad *adapter) InsertAuthor(author a.AuthorPort) {
	ad.repo.InsertAuthor(author)
}

func (ad *adapter) DeleteAuthor(username string) {
	ad.repo.DeleteAuthorByUsername(username)
}
