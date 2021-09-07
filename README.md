# algorand-vanity

Inspired by https://algovanity.com, this is a local/offline only option written in Go, and a lot faster. On my 2019 MacBook Pro I get about 50k addresses per second.

## Usage

```bash
git clone git@github.com:smonn/algorand-vanity.git
go get
go build
./algorand-vanity '^SOMEREGEX'
```

The first param should be a regular expression. Remember that an Algorand address is a base32 string matching `^[A-Z2-7]{58}$`. So any pattern that's a subset of that should work.

**Pro-tip: run multiple instances of this to speed things up further!**

> I'm a beginner with Go so I'm sure this could be improved a lot further. Happy hunting!
