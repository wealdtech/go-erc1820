# go-erc1820

[![Tag](https://img.shields.io/github/tag/wealdtech/go-erc1820.svg)](https://github.com/wealdtech/go-erc1820/releases/)
[![License](https://img.shields.io/github/license/wealdtech/go-erc1820.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/wealdtech/go-erc1820?status.svg)](https://godoc.org/github.com/wealdtech/go-erc1820)
[![Travis CI](https://img.shields.io/travis/wealdtech/go-erc1820.svg)](https://travis-ci.org/wealdtech/go-erc1820)
[![codecov.io](https://img.shields.io/codecov/c/github/wealdtech/go-erc1820.svg)](https://codecov.io/github/wealdtech/go-erc1820)

Go module to simplify talking to the [ERC-1820](https://eips.ethereum.org/EIPS/eip-1820) registry contract.


## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contribute](#contribute)
- [License](#license)

## Install

`go-erc1820` is a standard Go module which can be installed with:

```sh
go get github.com/wealdtech/go-erc1820
```

## Usage

`go-erc1820` provides simple access to the [ERC-1820](https://eips.ethereum.org/EIPS/eip-1820) registry contract.

### Example

```go
package main

import (
	erc1820 "github.com/wealdtech/go-erc1820"
)

func main() {
    client, err := ethclient.Dial("https://infura.io/v3/SECRET")
    if err != nil {
        panic(err)
    }

    registry, err := erc1820.New(client)
    if err != nil {
        panic(err)
    }

    // Fetch the implementer for an interface
    implementer, err := registry.InterfaceImplementer("ERC777TokensRecipient", )
    if err != nil {
        panic(err)
    }
}
```

## Maintainers

Jim McDonald: [@mcdee](https://github.com/mcdee).

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/wealdtech/go-erc1820/issues).

## License

[Apache-2.0](LICENSE) © 2019 Weald Technology Trading Ltd
