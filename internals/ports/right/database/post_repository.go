package database

import (
	a "github.com/jdavicp/blog/internals/ports/left/post"
)

type PortRepositoryPort interface {
	InsertPost(post a.PostPort)
	DeletePost()
}
