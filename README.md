# go-readabilityV2

`go-readabilityV2` extracts article HTML, plain text and metadata from supplied
HTML. It is a core-only fork of
[readeck/go-readability](https://codeberg.org/readeck/go-readability), descended
from `go-shiori/go-readability` and `mozilla/readability`.

## Philosophy

Our extractor packages share three principles:

1. **Bring your own HTML.** Keep page acquisition separate from extraction.
	The primary workflow uses HTML supplied by the caller, who controls fetching,
	caching, rendering, retries and scheduling.
2. **Stay close to upstream.** Preserve the algorithms and behavior of each
	package's declared upstream reference as closely as possible. Document
	deliberate differences and compatibility limits in [UPSTREAM.md](UPSTREAM.md)
	rather than claiming exact equivalence on every page.
3. **Provide very fast Go and Rust packages.** Run extraction natively, without
	a Python or Java runtime. Improve throughput and allocation efficiency while
	preserving intended behavior, and substantiate performance with reproducible
	benchmarks that report quality alongside speed.

## Overview

The current `go-readabilityV2` release is **v0.6.0**. It accepts an `io.Reader`
or an existing HTML tree and returns an article with HTML/text renderers and
metadata getters for title, byline, excerpt, site, image, language and dates.
It has no page fetcher, CLI or HTTP server; the original URL is context only.

Its reference is the `readeck/go-readability` v2.1.2 fork, based on
`mozilla/readability` 0.6.0 with the Go forks' improvements. The corresponding
Rust package, [rust-readability-v2](https://github.com/markusmobius/rust-readability),
uses `go-readabilityV2` as its behavioral reference. Callers own concurrency;
each extraction runs on the calling goroutine without an internal worker pool.

## Installation

```sh
go get github.com/markusmobius/go-readabilityV2@v0.6.0
```

Use Go 1.26.0 or newer. [go.mod](go.mod) selects Go 1.27.1 for development.
The module has no `/v2` suffix. See [CHANGELOG.md](CHANGELOG.md) for releases.

## Usage

Extract text from HTML already held in memory:

```go
package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	readability "github.com/markusmobius/go-readabilityV2"
)

func main() {
	pageURL, err := url.Parse("https://example.org/research")
	if err != nil {
		panic(err)
	}

	source := `<html><head><title>Research results</title></head><body><article>
<h1>Research results</h1>
<p>The research team compared several methods for extracting articles from saved
web pages. Every method received the same original HTML, and the evaluation
kept the reference text separate from the input supplied to each extractor.</p>
<p>The report records the complete experiment, including errors and repeated
measurements. Its results describe this collection of pages and do not promise
the same quality or execution time for every website.</p>
<p>The archived pages and analysis make it possible to repeat the comparison
and inspect the evidence behind each result.</p>
</article></body></html>`
	article, err := readability.FromReader(strings.NewReader(source), pageURL)
	if err != nil {
		panic(err)
	}

	fmt.Println(article.Title())
	if err := article.RenderText(os.Stdout); err != nil {
		panic(err)
	}
}
```

| Entry Point | Input |
| --- | --- |
| `FromReader` | HTML from an `io.Reader`, with decoding and parsing |
| `FromDocument` | An existing `*html.Node` and optional original URL |
| `CheckDocument` | A parsed tree for a quick article-likelihood check |
| `NewParser` | A parser whose extraction controls can be customized |

Use `Article.RenderHTML` or `Article.RenderText` for output. `Article.Node`
exposes the selected tree; getters expose metadata. See the
[go-readabilityV2 API reference](https://pkg.go.dev/github.com/markusmobius/go-readabilityV2)
for complete signatures and date getters.

## Options

Start with `NewParser()` and change only the controls you need:

| Option | Default | Effect |
| --- | --- | --- |
| `MaxElemsToParse` | `0` | No element-count limit; a positive value bounds accepted input trees. |
| `NTopCandidates` | `5` | Number of top candidates considered during article selection. |
| `CharThresholds` | `500` | Article-length threshold used by extraction retries. |
| `KeepClasses` | `false` | Remove classes except those explicitly preserved. |
| `ClassesToPreserve` | `page` | Classes retained when general class preservation is off. |
| `DisableJSONLD` | `false` | Use JSON-LD metadata unless disabled. |
| `AllowedVideoRegex` | `nil` | Use the default video-host matcher unless a custom expression is supplied. |

## Current Quality and Speed

The [2026-09-29 shared benchmark](https://github.com/markusmobius/content-extractor-benchmark/blob/ec719092d12f4d2a438dd29d9f4405aab6e0a321/README.md#results-2026-09-29)
compares the six packages below on **2,659 saved pages**: 983 LegoNews,
181 ScrapingHub and 1,495 WCXB.

### Extraction Speed

| Go Package (Measured Version) | Rust Package (Measured Version) | Go ms/page | Rust ms/page | Go/Rust |
| --- | --- | ---: | ---: | ---: |
| `go-readabilityV2` 0.6.0 | `rust-readability-v2` 0.6.5 | 4.755 | 3.945 | 1.21x |
| `go-domdistiller` 1.0.0 | `rust-domdistiller` 1.0.1 | 6.159 | 3.400 | 1.81x |
| `go-trafilatura` 2.2.6 (FAST) | `rust-trafilatura` 2.2.6 (FAST) | 11.329 | 6.570 | 1.72x |

Times are means of **all four measured passes after one warmup**. Go/Rust is
the named Go package's time divided by the named Rust package's time, not an
old/new release speedup. Later documentation-only releases do not change the
versions actually measured.

The run used Windows 11, Ryzen AI 7 PRO 350, Go 1.27.1 and Rust 1.98.1 GNU
with ThinLTO/mimalloc. Extraction includes required working copies, metadata
and text rendering. File I/O, startup, IPC, response serialization and scoring
are excluded. Comments and pagination are off; tables are on.
`go-trafilatura` and `rust-trafilatura` use FAST with external fallback disabled.
Power and sleep checks passed.

Parsing is separate: **Go 11.283 / Rust 6.386 ms/page**, charged once per
language/page for the shared suite. It includes decoding, DOM construction and
the separate `go-trafilatura` / `rust-trafilatura` noscript tree when needed.
These are extraction-stage comparisons, not complete request latencies.

### Text Quality

Each named pair has equal text scores. Errors are listed in LegoNews /
ScrapingHub / WCXB order and remain in the scoring denominators.

| Go Package | Rust Package | LegoNews F1 | ScrapingHub F1 | WCXB F1 | Errors |
| --- | --- | ---: | ---: | ---: | --- |
| `go-readabilityV2` | `rust-readability-v2` | 87.82711% | 95.20557% | 78.47603% | 7 / 0 / 28 |
| `go-domdistiller` | `rust-domdistiller` | 86.74080% | 92.74280% | 74.39696% | 0 / 0 / 0 |
| `go-trafilatura` (FAST) | `rust-trafilatura` (FAST) | 90.91534% | 96.15663% | 78.51703% | 4 / 0 / 10 |

The corpora use different scoring rules; their F1 scores must not be averaged.
Equal text scores do not imply identical metadata: `go-trafilatura` and
`rust-trafilatura` differ on one title and one author field. The
[full report](https://github.com/markusmobius/content-extractor-benchmark/blob/49c426d6135df81b7d492bea7e6aec8e6d77d80c/go_rust_shared_performance_2026_09_29.json)
contains metadata scores, differences, every pass and source/build identities.

## Compatibility and Limitations

- **Declared reference.** Follow `readeck/go-readability` v2.1.2 and its
	`mozilla/readability` 0.6.0 ancestry, with the pinned dependencies recorded in
	[UPSTREAM.md](UPSTREAM.md). Different HTML parser versions can change output.
- **No browser rendering.** The library does not execute JavaScript, compute
	layout or retrieve content absent from the supplied HTML.
- **Heuristic results.** Boilerplate can remain or content can be missed.
	Check extraction and rendering errors; a quick `CheckDocument` result is
	not a substitute for full extraction.
- **Caller-owned concurrency.** Use a separate parser and working DOM for
	concurrent extractions. Parser reuse can retain language state.
- **Not a sanitizer.** Sanitize extracted HTML before displaying untrusted input.

Historical measurements remain in
[benchmark-results-0.6.0-crate.json](benchmark-results-0.6.0-crate.json) and the
[technical reference](UPSTREAM.md); they use different timing boundaries.

## Development

From a checkout with Go installed:

```sh
go test -mod=readonly ./...
go vet -mod=readonly ./...
```

Generated parser helpers are checked in. Regenerating them with `make generate`
requires re2go and the repository's generation toolchain; ordinary use does not.
Documentation and release rules are in [AGENTS.md](AGENTS.md).

## License and Credits

The original MIT [LICENSE](LICENSE), naming Radhi Fadlillah, is retained unchanged.
Arc90 Inc created the original JavaScript algorithm credited by
[mozilla/readability](https://github.com/mozilla/readability), developed and
maintained by Mozilla and its contributors. Radhi Fadlillah
created the original [go-shiori/go-readability](https://github.com/go-shiori/go-readability)
port, subsequently maintained by Felipe Martin and its contributors. The Readeck
contributors developed the [readeck/go-readability](https://codeberg.org/readeck/go-readability)
v2 fork from which this package descends. Markus Mobius maintains `go-readabilityV2`.
These credits distinguish the original work from its later forks and maintenance.
