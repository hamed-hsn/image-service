package protocol

type LinkValidator interface {
	ValidateLink(string) error
}
