fetch:
	depviz fetch -github-token=${GITHUB_TOKEN} gnolang/roadmap

generate:
	depviz gen json gnolang/roadmap > roadmap.json

install:
	cd misc/deps; make install
