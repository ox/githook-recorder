#githook-recorder

Passed a non-existing folder name, `githook-recorder` runs a server that can receive GitHub
git commit hook messages, and write them to individual, ordered, json files in
that folder name.

## Installation

    $ go install github.com/ArtemTitoulenko/githook-recorder

## Usage

    $ githook-recorder [--port=9090] <session name>


