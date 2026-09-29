# Changelog

## Documentation - 2026-09-29

- Apply the approved nine-section README format, including the three shared
  philosophy principles, a runnable supplied-HTML example and actual options.
- Specify the full structure and required content in AGENTS.md, including
  named credits for Arc90 Inc, Mozilla, Radhi Fadlillah, Felipe Martin,
  the Readeck contributors and Markus Mobius. Remove unrelated package policies.
- Align the README's six-engine comparison with the September 29 shared
  benchmark, including common units, measured versions and timing boundaries.
- Add AGENTS.md with instructions for README, UPSTREAM, CHANGELOG and
  coordinated release documentation.
- Keep `go-readabilityV2` 0.6.0, runtime source, dependencies and all measured data
  unchanged. No new module version or benchmark run.

## Documentation - 2026-09-23

- Refresh README quality and six-engine speed comparisons from the published
  [benchmark JSON](https://github.com/markusmobius/content-extractor-benchmark/blob/d433ab637f0a56c0926aa3698f470794a553472f/go_rust_shared_performance_2026_09_23.json),
  with separate metadata scores and exact provenance in [UPSTREAM.md](UPSTREAM.md).
- Readability text F1 is 87.82711% / 95.20557% / 78.47603% on LegoNews /
  ScrapingHub / WCXB. Selected Go/Rust extraction is 2.669 / 2.441 ms/page
  (1.09x); all-four means are 2.705 / 2.451 ms/page. Shared parsing is separate.
- Go-ReadabilityV2 remains 0.6.0. No source, module, dependency, fixture or tag
  changes are part of this documentation update; historical results are retained.