package quandoscript

type QuandoScript interface {
	Generate() string // returns the generated template QuandoScript for a widget in a block
}
