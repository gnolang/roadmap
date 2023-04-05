# roadmap

### Usage:

```bash
docker build . --build-arg github_token=$GITHUB_TOKEN -t depviz-roadmap
```
    
```bash
make fetch # GITHUB_TOKEN="${GITHUB_TOKEN}" if not exported
make generate # this generate a roadmap.json inside the output directory
```
