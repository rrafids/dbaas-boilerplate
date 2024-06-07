#!/bin/bash -e
# gofmt wrapper that will exit with non-zero code if there's any format violations


dir="${1:-.}"

# print the diff
gofmt -s -d "$dir"

# list files with diff
files="$(gofmt -s -l "$dir" | tr '\r\n' ' ')"


echo "Running gofmt in $dir"
if [[ "$files" == "" ]]; then
    echo "OK"
    exit 0
else
    echo
    echo "Files with gofmt formatting issues:"
    for f in $files; do
        echo "- $f"
    done

    echo
    echo "Run this command to fix, then commit the changed files:"
    echo "gofmt -w -s $files"
    exit 1
fi
