package main

import "github.com/perrito666/blogpost_goplogins/contract"

// ShowcaseElement is the variable used as the plugin entry
// point.
var ShowcaseElement contract.Plugin = &plugin{"0.2"}

// Showcase returns the current ShowcaseElement.
func Showcase() contract.Plugin { return ShowcaseElement }

type plugin struct {
	version string
}

func (p *plugin) Version() string {
	return p.version
}

func (p *plugin) IsAcceptable() bool {
	return false
}

func main() {}
