# shipator

[![GitHub release](https://img.shields.io/github/v/release/brainnco/shipator)](https://github.com/brainnco/shipator/releases/latest)
[![CI](https://github.com/brainnco/shipator/workflows/CI/badge.svg?branch=main)](https://github.com/brainnco/shipator/actions/workflows/CI.yml?query=branch%3Amain)
[![License](https://img.shields.io/github/license/brainnco/shipator)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/brainnco/shipator)](https://goreportcard.com/report/github.com/brainnco/shipator)

Inject environment variables at runtime into your SPA bundle.

TODO:

- [ ] Create a logo
- [ ] Write readme. delete old readme.
- [ ] Release version.

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
  -version
        Prints current version

Examples
  $ shipator build/index.html
  $ shipator -prefix REACT_APP -placeholder __ENV__ build/index.html
  $ shipator -placeholder __VARS__ build/index.html
  $ shipator -prefix VUE_APP build/index.html
```

## Example

```docker
FROM node:12.16.1-alpine as builder
WORKDIR /app
COPY yarn.lock /app/yarn.lock
COPY package.json /app/package.json
RUN yarn install --frozen-lockfile
COPY . /app
RUN yarn build

FROM brainnco/shipator:0.1.0-rc5
COPY --from=builder /app/build /app/shipator/html
```

TODO: example deploying a react application with ngnix + k8s.

## Changelog

See the [changelog](CHANGELOG.md).

## Contributing

See the [contributing file](CONTRIBUTING.md).

## License

Copyright 2020 brainn.co.

Shipator source code is released under Apache 2 License.

Check [LICENSE](LICENSE) file for more information.
