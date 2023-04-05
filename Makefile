## MAKEFILE ##
build:
	docker build -t depviz .
fetch:
	docker run -it -v ${PWD}/output:/output gno-roadmap -store-path=/output/.db fetch -github-token=${GITHUB_TOKEN} gnolang/roadmap

generate:
	docker run -it -v ${PWD}/output:/output gno-roadmap -store-path=/output/.db gen json gnolang/roadmap > ./output/roadmap.json
