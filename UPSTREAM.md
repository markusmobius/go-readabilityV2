# Upstream Reference

## Released Suite Benchmark

The [2026-09-23 JSON](https://github.com/markusmobius/content-extractor-benchmark/blob/d433ab637f0a56c0926aa3698f470794a553472f/go_rust_shared_performance_2026_09_23.json)
is authoritative for the current [README tables](README.md#current-quality-and-speed).
Its published-file SHA-256 (LF line endings) is `7d7be9839f1652606cb91850af5134b188f2508be25623df889372dab4a06cc6`.
Read text scores at `quality[worker][engine].evaluations[corpus].overall.f1`,
selected timings at `overall`, and all-four timings at `all_passes`.

Go pins are Readability 0.6.0 (`db6ab179f951f80ad176850001aaf486e2fbd367`),
DomDistiller 1.0.0 (`25b8d046ffb4053bf68345d6fa59bc9ae1961ad8`) and Trafilatura
2.2.2 (`f4684e100869274311107325e3b72e47cc78db20`). Rust uses Readability 0.6.3,
DomDistiller 1.0.1 and Trafilatura 2.2.4; exact commits, dependency graphs and
binary hashes are in the build receipts. Rust-Trafilatura remains private;
reproducing that suite requires authorized access. No Go version or tag changes.

| Implementation | Author Sets Exact / 1,290 | Author-Unit F1 | Titles Exact / 2,364 | Dates Exact / 1,530 |
| --- | ---: | ---: | ---: | ---: |
| go-readabilityV2-0.6.0 | 640 | 56.38767% | 1,247 | 763 |
| rust-readability-0.6.3 | 640 | 56.38767% | 1,247 | 763 |
| go-domdistiller-1.0.0 | 0 | 0.00000% | 1,106 | 0 |
| rust-domdistiller-1.0.1 | 0 | 0.00000% | 1,106 | 0 |
| go-trafilatura-2.2.2 | 695 | 58.80923% | 1,228 | 1,227 |
| rust-trafilatura-2.2.4 | 696 | 58.86640% | 1,227 | 1,227 |

Metadata uses only nonempty supplied annotations; unannotated is not negative,
and missing output is not filled by another engine. All six scored-output
digests match the preceding September 22 report. Trafilatura's two Go/Rust
differences concern one title and one author, not extracted text.

The full 2,659-page development run used seed 20260922, one warmup and four
measured passes. Passes 1 and 3 were selected by combined extraction time for
every row (5,318 observations each); all-four means retain 10,636 observations.
Worker order is balanced per page; three-engine order is a partial six-pass
block. Go uses `GOMAXPROCS=1`, `GOGC=100`, without forced collection. Native
timers exclude file reads and IPC; parsing and extraction stay separate.
The 26,590-response audit passed with no recorded sleep and AC power throughout.
This is a coordinated released-suite comparison using Go-Trafilatura's graph,
not an isolated language change, standalone-reader test or unseen holdout.

## Source and Scope

The module is `github.com/markusmobius/go-readabilityV2`, version 0.6.0.
It forks Readeck's `v2` branch, v2.1.2, at
[`b18540d99ebf105cd67122585a0a41ec299b70bc`](https://codeberg.org/readeck/go-readability/commit/b18540d99ebf105cd67122585a0a41ec299b70bc).
Readeck derives from Go-Shiori Readability and Mozilla Readability.js; the
inherited algorithm tracks Mozilla 0.6.0 plus the Go forks' improvements.

This fork retains reader decoding, DOM extraction, parser options, metadata,
date getters and rendering. URL fetching, request modifiers, CLI/server,
network helper scripts and diagnostic parser logging were removed. Extraction
runs on the caller's goroutine; acquisition and concurrency belong to callers.
Dependencies remain recorded in [go.mod](go.mod) and [go.sum](go.sum).

The historical 983-page comparison and its different timing protocol remain in
[benchmark-results-0.6.0-crate.json](benchmark-results-0.6.0-crate.json) and the
[README](README.md#historical-quality-and-performance). Equal corpus quality is
not a claim of identical HTML across different parser dependency versions.
The upstream MIT [LICENSE](LICENSE) remains unchanged.