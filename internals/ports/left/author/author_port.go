package author

type AuthorPort interface {
	Name() string
	Email() string
	Username() string
}
