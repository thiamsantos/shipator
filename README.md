# shipator

[![GitHub release](https://img.shields.io/github/v/release/brainnco/shipator)](https://github.com/brainnco/shipator/releases/latest)
[![CI](https://github.com/brainnco/shipator/workflows/CI/badge.svg?branch=main)](https://github.com/brainnco/shipator/actions/workflows/CI.yml?query=branch%3Amain)
[![License](https://img.shields.io/github/license/brainnco/shipator)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/brainnco/shipator)](https://goreportcard.com/report/github.com/brainnco/shipator)

Inject environment variables at runtime into your SPA bundle.

TODO:

- [ ] Create a logo
- [ ] Write readme. delete old readme.
- [ ] Release CLI, maybe using github releases. https://goreleaser.com/ ?
- [ ] Official docker image?

## Background

TODO: why?

## How it works

TODO: explaing what it does.

## Usage

TODO: how to use the cli. describe CLI options

```
Usage
  $ shipator [options] target

Options
  -placeholder string
        Placeholder in the target (default "__ENV__")
  -prefix string
        Prefix of the env vars to inject (default "REACT_APP")

Examples
  $ shipator build/index.html
  $ shipator -prefix REACT_APP -placeholder __ENV__ build/index.html
  $ shipator -placeholder __VARS__ build/index.html
  $ shipator -prefix VUE_APP build/index.htm
```

## Example

TODO: example deploying a react application with ngnix + k8s.

## Changelog

See the [changelog](CHANGELOG.md).

## Contributing

See the [contributing file](CONTRIBUTING.md).

## License

Copyright 2020 brainn.co.

Shipator source code is released under Apache 2 License.

Check [LICENSE](LICENSE) file for more information.
