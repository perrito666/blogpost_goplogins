package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"plugin"
	"sort"
	"strings"

	"github.com/perrito666/blogpost_goplogins/contract"
)

const (
	pluginSuffix = ".so"
	pluginFolder = "./plugins"
)

func findAvailableVersions(folder string) ([]string, error) {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}
	plugins := []string{}
	for _, file := range files {
		fname := file.Name()
		if strings.HasSuffix(fname, pluginSuffix) {
			plugins = append(plugins, fname)
		}
	}
	sort.Strings(plugins)
	for i := len(plugins)/2 - 1; i >= 0; i-- {
		opp := len(plugins) - 1 - i
		plugins[i], plugins[opp] = plugins[opp], plugins[i]
	}
	return plugins, nil
}

func main() {

	pluginVersions, err := findAvailableVersions(pluginFolder)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to list plugins: %v", err))
	}

	var acceptedPlugin contract.Plugin
	var ok bool
	for _, version := range pluginVersions {
		fmt.Println(version)
		p, err := plugin.Open(path.Join(pluginFolder, version))
		if err != nil {
			fmt.Println(fmt.Errorf("showcase plugin is not available: %v", err))
			continue
		}

		e, err := p.Lookup("Showcase")
		if err != nil {
			fmt.Println(fmt.Errorf("showcase element is not present: %v", err))
			continue
		}

		var pluginAccessor func() contract.Plugin
		pluginAccessor, ok = e.(func() contract.Plugin)
		if ok {
			acceptedPlugin = pluginAccessor()
			if acceptedPlugin.IsAcceptable() {
				break
			}
		}
	}
	if !ok {
		fmt.Println(fmt.Errorf("no suitable plugin version found"))
		return
	}
	fmt.Println(fmt.Sprintf("Found newest valid plugin version: %q", acceptedPlugin.Version()))
}
