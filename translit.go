package translit

type Transliterator interface {
	Convert(str string) string
}
