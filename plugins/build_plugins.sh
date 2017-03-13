#!/bin/sh
if [ ! -d "../plugin_sources" ]; then
    echo "this script should be run from inside the plugins directory, sorry for the poor resiliency"
    exit 1
fi

for source_plugin in $(ls ../plugin_sources); do
    go build -buildmode=plugin ../plugin_sources/$source_plugin/...;
    if [ $? -eq 0 ]; then
        echo "added version ${source_plugin}";
    fi
done