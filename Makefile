all: fetch gen-json gen-image

fetch:
	go run moul.io/depviz/v3/cmd/depviz --store-path=output/ fetch -github-token=${GITHUB_TOKEN} gnolang/roadmap

output/roadmap.json:
	go run moul.io/depviz/v3/cmd/depviz --store-path=output/ gen json gnolang/roadmap > $@.tmp
	@mv $@.tmp $@

gen-image: output/roadmap.json
	go run ./gen-graph
	dot -Tpng output/roadmap.dot > output/roadmap.png
	dot -Tsvg output/roadmap.dot > output/roadmap.svg

clean:
	rm -f output/roadmap.json

fclean:
	rm -rf output/
