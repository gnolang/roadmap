build:
	docker build -t depviz .

fetch:
	docker run -it -v "$(PWD):$(PWD)" -w "$(PWD)/output" depviz -store-path=db fetch -github-token=${GITHUB_TOKEN} gnolang/roadmap

generate:
	docker run -it -v "$(PWD):$(PWD)" -w "$(PWD)/output" depviz -store-path=db gen json gnolang/roadmap > roadmap.json
