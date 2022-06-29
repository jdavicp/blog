package post

import a "github.com/jdavicp/blog/internals/ports/left/author"

type adapter struct {
	title   string
	content string
	author  a.AuthorPort
}

func New(title string, content string, author a.AuthorPort) *adapter {
	if author.Username() == "" {
		return &adapter{}
	}
	return &adapter{
		title,
		content,
		author,
	}
}

func (ad *adapter) Author() a.AuthorPort {
	return ad.author
}

func (ad *adapter) Title() string {
	return ad.title
}

func (ad *adapter) Content() string {
	return ad.content
}
