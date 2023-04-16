all: fetch gen-image

fetch: output/indexes.bolt
gen-image: output/roadmap.svg output/roadmap.png output/gno.svg output/gno.png output/ecosystem.svg output/ecosystem.png

clean:
	rm -f output/*.json

fclean:
	rm -rf output/

## Advancedrules

gen-json: output/roadmap.json output/gno.json output/ecosystem.json
gen-dots: output/roadmap.dot output/gno.dot output/ecosystem.dot

DEPVIZ=go run moul.io/depviz/v3/cmd/depviz --store-path=output/

# depviz database, shared between repos
output/indexes.bolt:
	$(DEPVIZ) fetch -github-token=${GITHUB_TOKEN} gnolang/roadmap gnolang/gno gnolang/awesome-gno

# per-flavor rules
output/roadmap.json: output/indexes.bolt
	$(DEPVIZ) gen -hide-prs -hide-external-deps json gnolang/roadmap > $@.tmp
	@mv $@.tmp $@
output/gno.json: output/indexes.bolt
	$(DEPVIZ) gen -show-closed json gnolang/gno > $@.tmp
	@mv $@.tmp $@
output/ecosystem.json: output/indexes.bolt
	$(DEPVIZ) gen -hide-prs -hide-isolated json gnolang/gno gnolang/roadmap gnolang/awesome-gno > $@.tmp
	@mv $@.tmp $@

# generic conversions
%.dot: %.json
	go run ./gen-graph -i $< -o $@

%.svg: %.dot
	dot -Tsvg $< > $@

%.png: %.dot
	dot -Tpng $< > $@
