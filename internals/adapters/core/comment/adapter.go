package comment

import (
	a "github.com/jdavicp/blog/internals/ports/left/author"
	p "github.com/jdavicp/blog/internals/ports/left/post"
)

type adapter struct {
	author  a.AuthorPort
	content string
	post    p.PostPort
}

func New(author a.AuthorPort, content string, post p.PostPort) *adapter {
	if author.Email() == "" || post.Title() == "" {
		return &adapter{}
	}
	return &adapter{
		author,
		content,
		post,
	}
}

func (ad *adapter) Author() a.AuthorPort {
	return ad.author
}

func (ad *adapter) Content() string {
	return ad.content
}

func (ad *adapter) Post() p.PostPort {
	return ad.post
}
