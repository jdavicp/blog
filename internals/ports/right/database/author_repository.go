package database

import (
	a "github.com/jdavicp/blog/internals/ports/left/author"
)

type AuthorRepositoryPort interface {
	InsertAuthor(author a.AuthorPort)
	DeleteAuthorByUsername(username string)
}
