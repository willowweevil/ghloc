#!/bin/bash

echo $GITHUB_TOKEN > .githubtoken

gh auth login --with-token < .githubtoken

rm .githubtoken

GITHUB_ACCOUNT=$(gh auth status | grep -Eo "(github.com account).*\(" | awk '{print $3}')

gh repo list | while read -r repo _; do gh repo clone "${repo}" "${repo}"; done

cloc . --include-lang=GDScript,Go,Lua,Python,"Jupyter Notebook","Bourne Shell","DOS Batch" \
	--by-percent cmb \
	--hide-rate | ./converter

rm -rf $GITHUB_ACCOUNT/
