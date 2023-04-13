all: fetch gen-json gen-image

fetch:
	go run moul.io/depviz/v3/cmd/depviz fetch -github-token=${GITHUB_TOKEN} gnolang/roadmap

gen-json:
	mkdir -p output
	go run moul.io/depviz/v3/cmd/depviz gen json gnolang/roadmap > output/roadmap.json

gen-image:
	go run ./gen-graph > output/roadmap.dot
	dot -Tpng output/roadmap.dot > output/roadmap.png
	dot -Tsvg output/roadmap.dot > output/roadmap.svg
