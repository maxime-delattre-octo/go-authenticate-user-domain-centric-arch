package ports

type Random32CharsGenerator interface {
	Execute() string
}
