all: install fetch gen-json gen-image

install:
	cd misc/deps; make install

fetch:
	depviz fetch -github-token=${GITHUB_TOKEN} gnolang/roadmap

gen-json:
	mkdir -p output
	depviz gen json gnolang/roadmap > output/roadmap.json

gen-image:
	go run ./gen-graph > output/roadmap.dot
	dot -Tpng output/roadmap.dot > output/roadmap.png
	dot -Tsvg output/roadmap.dot > output/roadmap.svg
