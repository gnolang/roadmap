# Contributing

Manfred's disclaimer: Sorry if the repo looks messy right now. I'm transitioning from an old tool and noticed one of my dependencies (cayley) is no longer maintained. If you run into issues, ping me on Signal. I'm working on improving things, so expect a better experience soon.

To avoid issues fetching data or sharing your GitHub token, you can grab the roadmap.json file from this URL: https://github.com/gnolang/roadmap/tree/generate/output.

## Recommended flows

If you're working on the graph rendering, I recommend focusing solely on the `gen-graph` command and skipping the networking component.

You can use the following command which offers fast automatic preview (tested on Mac):

    find . | entr -c sh -xec "make gen-image && open output/roadmap.svg".
