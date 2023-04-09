all: install fetch gen-json gen-image

install:
	cd misc/deps; make install

fetch:
	depviz fetch -github-token=${GITHUB_TOKEN} gnolang/roadmap

gen-json:
	depviz gen json gnolang/roadmap > roadmap.json

gen-image:
	@echo "TODO"
	@exit 1
