package comment

import (
	"github.com/jdavicp/blog/internals/ports/left/author"
	"github.com/jdavicp/blog/internals/ports/left/post"
)

type CommentPort interface {
	Author() author.AuthorPort
	Content()
	Post() post.PostPort
}
