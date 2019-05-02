# Simple Site Builder

This is the root repository for the simple site builder.

# Testing

Currently the package only has unit tests. To run them do
```bash
make test
```

# Building

To build the package run
```bash
make build
```

stage 1: context resolver
    have:
        url
    get:
        site id
        page id
        context (query, fragment?, etc)

stage 2: page loader
    have:
        site id
        page id
        context
    get:
        page fragments & content

stage 3: html generator
    have:
        page fragments & content
    get:
        html
