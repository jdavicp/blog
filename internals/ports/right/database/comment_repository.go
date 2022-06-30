package database

import a "github.com/jdavicp/blog/internals/ports/left/comment"

type CommentRepositoryPort interface {
	InsertComment(comment a.CommentPort)
}
