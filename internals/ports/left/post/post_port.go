package post

import a "github.com/jdavicp/blog/internals/ports/left/author"

type PostPort interface {
	Author() a.AuthorPort
	Title() string
	Content() string
}
