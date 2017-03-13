package contract

// Plugin represet a valid plugin.
type Plugin interface {
	IsAcceptable() bool
	Version() string
}
