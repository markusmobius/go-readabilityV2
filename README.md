# Go-Readability V2

This is a fork of [codeberg.org/readeck/go-readability](https://codeberg.org/readeck/go-readability),
branch `v2`, taken at commit
[`b18540d99ebf105cd67122585a0a41ec299b70bc`](https://codeberg.org/readeck/go-readability/commit/b18540d99ebf105cd67122585a0a41ec299b70bc).
Readeck's implementation is itself a fork of
[github.com/go-shiori/go-readability](https://github.com/go-shiori/go-readability),
originally written by Radhi Fadlillah and maintained by Felipe Martin and GitHub contributors.

The inherited implementation matches the
[Mozilla Readability.js 0.6.0](https://github.com/mozilla/readability/blob/main/CHANGELOG.md#060---2025-03-03)
baseline, plus the Go forks' additional fixes and performance improvements.
Going forward, we strive to mirror Mozilla's original JavaScript Readability
implementation as it evolves.

Our philosophy is **bring your own HTML**. The library extracts readable article
content and metadata from an `io.Reader` or an existing `html.Node`. Everything
outside the core extraction library has been removed: URL fetching, request
modifiers, the CLI and HTTP server, network helper scripts, and diagnostic parser
logging. The caller controls how HTML is acquired and how work is scheduled.

Reader decoding, DOM extraction, parser options, metadata and date getters, and
HTML/plain-text rendering are retained.

**The library is single-threaded.** Each extraction runs on the calling goroutine,
with no internal extraction workers or worker pool. It is suitable for servers
running many engines in parallel: give each engine its own parser and input DOM,
and let the server control concurrency.

## Usage

Version **0.6.0** uses the module path `github.com/markusmobius/go-readabilityV2`
and requires Go 1.26 or newer. Supply HTML from your own reader or DOM:

```go
package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	readability "github.com/markusmobius/go-readabilityV2"
)

func main() {
	srcFile, err := os.Open("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	baseURL, err := url.Parse("https://example.com/path/to/article")
	if err != nil {
		log.Fatal(err)
	}
	article, err := readability.FromReader(srcFile, baseURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found article with title %q\n\n", article.Title())
	// Print the parsed, cleaned-up HTML markup of the article.
	if err := article.RenderHTML(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
```

The base URL resolves relative links and related metadata; it is never fetched.
Use `FromDocument` for an existing DOM, or `NewParser` to customize extraction.
`CheckDocument` provides the fast readability check without full extraction.

## Current Quality and Speed

The [2026-09-23 benchmark JSON](https://github.com/markusmobius/content-extractor-benchmark/blob/d433ab637f0a56c0926aa3698f470794a553472f/go_rust_shared_performance_2026_09_23.json)
is the source for these tables. All six released engines use the same 2,659
development pages: 983 LegoNews, 181 ScrapingHub and 1,495 WCXB. The three F1
scores use different scoring rules and must not be averaged. Errors are shown
in that corpus order and remain in the denominators.

| Implementation | LegoNews F1 | ScrapingHub F1 | WCXB F1 | Errors |
| --- | ---: | ---: | ---: | --- |
| go-readabilityV2-0.6.0 | 87.82711% | 95.20557% | 78.47603% | 7 / 0 / 28 |
| rust-readability-0.6.3 | 87.82711% | 95.20557% | 78.47603% | 7 / 0 / 28 |
| go-domdistiller-1.0.0 | 86.74080% | 92.74280% | 74.39696% | 0 / 0 / 0 |
| rust-domdistiller-1.0.1 | 86.74080% | 92.74280% | 74.39696% | 0 / 0 / 0 |
| go-trafilatura-2.2.2 | 90.88412% | 96.15663% | 78.49352% | 3 / 0 / 10 |
| rust-trafilatura-2.2.4 | 90.88412% | 96.15663% | 78.49352% | 3 / 0 / 10 |

| Implementation | Shared Parse ms/page | Extraction ms/page | Extraction ms/page, All Four Passes |
| --- | ---: | ---: | ---: |
| go-readabilityV2-0.6.0 | 5.638 | 2.669 | 2.705 |
| rust-readability-0.6.3 | 2.525 | 2.441 | 2.451 |
| go-domdistiller-1.0.0 | 5.638 | 3.618 | 3.628 |
| rust-domdistiller-1.0.1 | 2.525 | 1.965 | 1.973 |
| go-trafilatura-2.2.2 | 5.638 | 6.815 | 6.839 |
| rust-trafilatura-2.2.4 | 2.525 | 4.000 | 4.026 |

One full warmup precedes four measured passes. The first two timing columns
use the common best two complete passes (1 and 3), an optimistic estimate;
the final column retains the all-four mean. Go/Rust extraction ratios from
unrounded means are **1.09x Readability, 1.84x DomDistiller and 1.70x Trafilatura**.
Parsing is charged once per language/page, not once per engine. All-four parse
means are Go 5.667 and Rust 2.531 ms/page.

These Windows 11 / Ryzen AI 7 PRO 350 measurements use Go 1.27.1 and Rust
1.98.1 GNU, the released Trafilatura dependency graphs, and Rust ThinLTO/mimalloc.
Parsing includes eager decoding, normalization and DOM construction after the
file read; extraction includes private working copies, native metadata and text
rendering. Rust temporary trees are destroyed inside the timer; Go uses normal
GC, which can cross stage boundaries. File I/O, startup, IPC and scoring are
excluded. Fallbacks, comments and pagination are off; tables are on.

Every repeated scored output was stable. Go/Rust Readability and DomDistiller
match all scored outputs; Trafilatura retains two metadata-only differences.
Separate metadata scores, exact source pins and protocol limits are in
[UPSTREAM.md](UPSTREAM.md#released-suite-benchmark). Go-ReadabilityV2 remains
0.6.0; [CHANGELOG.md](CHANGELOG.md) records this documentation-only refresh.
These are shared-input suite timings, not standalone end-to-end latency.
Older measurements below use different protocols.

## Historical Quality and Performance

Measured on 2026-09-16 using all **983 labeled pages** from
[content-extractor-benchmark](https://github.com/markusmobius/content-extractor-benchmark/tree/466fdbee8a504441eb78ed11d71c1da220681cab),
the same pinned corpus used for Go-DomDistiller and Rust-DomDistiller.
The Codeberg row uses the unmodified v2.1.2 source at the commit above and its
original dependencies. The Go fork and Rust rows use version 0.6.0.

### Extraction Quality

| Extractor | Precision | Recall | F1 | Accuracy |
| --- | ---: | ---: | ---: | ---: |
| Codeberg Go-Readability v2.1.2 | 0.8705 | 0.8862 | 0.8783 | 0.8774 |
| Go-ReadabilityV2 | 0.8705 | 0.8862 | 0.8783 | 0.8774 |
| Rust-Readability | 0.8705 | 0.8862 | 0.8783 | 0.8774 |

All three engines have **exactly the same quality counts**: TP 2,601, FN 334,
FP 387 and TN 2,561. The same seven empty extractions are scored as empty text,
not skipped. Scores use the benchmark's case-sensitive snippet matching and
globally aggregated counts, not token-level scoring or metadata-quality scores.

The Go fork and Rust match exactly on all 983 pages for text, HTML, title, byline,
excerpt, site name, image URL, favicon, language and errors. Codeberg's original
dependencies use x/net 0.41.0 and x/text 0.26.0; the fork uses 0.59.0 and 0.42.0.
With the original dependencies, 387 pages differ only in the HTML field; all
other compared fields and quality counts agree. A separate Codeberg control
using the fork's dependency versions matches all ten fields exactly on 983/983
pages. Thus equal quality is not a claim of byte-identical HTML across different
parser versions.

### Extraction Time

Median time per complete 983-page pass; ranges show the measured minimum and
maximum across all **102 measured passes per engine**.
Speedup is Go-ReadabilityV2's median time divided by each engine's median time.

| Extractor | Median | Range | Speedup vs Go-ReadabilityV2 |
| --- | ---: | ---: | ---: |
| Codeberg Go-Readability v2.1.2 | 2,044 ms | 1,854-2,792 ms | 0.96x |
| Go-ReadabilityV2 | 1,955 ms | 1,811-2,712 ms | 1.00x |
| Rust-Readability | 960 ms | 876-1,521 ms | 2.04x |

Rust's median speedup was **2.04x over Go-ReadabilityV2** and **2.13x over
Codeberg**.

Measured on an AMD Ryzen AI 7 PRO 350 under Linux/WSL2, pinned to one logical CPU,
with Go 1.27.1 and Rust 1.98.1 release builds. Each engine received two warmup
passes followed by 102 measured passes, cycling all six engine orders 17 times.
The dependency-aligned Codeberg control is for correctness only;
the timed Codeberg row retains its original dependency versions.

Timing includes extraction from pre-parsed DOMs, plain-text rendering and snippet
scoring. It excludes file I/O, decoding, initial HTML/URL parsing, startup, IPC
and the extra HTML/metadata collection used for exact-output checks. Timing varies
between runs: these are single-machine measurements, not a guaranteed speedup or an
end-to-end reader/network benchmark.

Raw samples, dependency graphs and source/binary fingerprints are retained in
[benchmark-results-0.6.0-crate.json](benchmark-results-0.6.0-crate.json).
The runner and reproduction instructions are maintained with
[Rust-Readability](https://github.com/markusmobius/rust-readability/blob/main/UPSTREAM.md#shared-benchmark).

## License

The upstream MIT [LICENSE](LICENSE) is retained unchanged.
