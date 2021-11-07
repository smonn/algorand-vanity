# algorand-vanity

Inspired by https://algovanity.com, this is a local/offline only option written in Go, and a lot faster. On my 2019 MacBook Pro I get about ~300k addresses per second.

## Install

### Option 1, clone and build yourself

```bash
git clone git@github.com:smonn/algorand-vanity.git
cd algorand-vanity
go get
go build
```

### Option 2, go install

```bash
go install github.com/smonn/algorand-vanity
```

### Option 3, download latest release

Visit the [releases page](https://github.com/smonn/algorand-vanity/releases) and download the latest release for your platform.

## Usage

```bash
algorand-vanity '^ABC[2-7]'
```

The first param should be a regular expression. Remember that an Algorand address is a base32 string matching `^[A-Z2-7]{58}$`. So any pattern that's a subset of that should work.

> I'm a beginner with Go so I'm sure this could be improved a lot further. Happy hunting!
