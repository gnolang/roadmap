all: install fetch generate

install:
	cd misc/deps; make install

fetch:
	depviz fetch -github-token=${GITHUB_TOKEN} gnolang/roadmap

generate:
	depviz gen json gnolang/roadmap > roadmap.json
