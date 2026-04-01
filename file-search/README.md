# FileSearcher

This FileSearcher module is a Go implementation for finding substrings within
files. It partially overlaps functionality with `grep` and `find`.  The purpose
for this module is to provide a simple way to search for substrings within
a GitHub repo.

It was also just fun to build :D.

## Usage

```bash

go run . --repo path/to/repo/folder --search "strings,to,search,for"

```

For example, a big set of strings to search for:

```bash
go run . --repo /Users/someone/repos/vertex-ai-platform \

--output /Users/someone/Desktop/output.txt \

--search "Vertex AI Conversation,Med-PaLM,Conversational AI Platform,Vector Engine,generative AI agent,Vertex AI Infra: GDC,AI Platform Pipelines,Vertex Deep Learning Containers,AI Platform Deep Learning VM Image,AI Platform Deep Learning Containers,AI Platform Notebooks,Video Search,Video Search AI,Video Search API,AI Platform Data Labeling Service,Chat Experience,Gen App Builder,Generative AI App Builder,Embeddings Gecko,Multimodal Embeddings,Code Bison,Code Generation,Code Chat,Code Completion,Image Generation,Vertex AI Matching Engine,Enterprise Search" 

```