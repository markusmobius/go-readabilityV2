# Changelog

## Documentation - 2026-09-23

- Refresh README quality and six-engine speed comparisons from the published
  [benchmark JSON](https://github.com/markusmobius/content-extractor-benchmark/blob/d433ab637f0a56c0926aa3698f470794a553472f/go_rust_shared_performance_2026_09_23.json),
  with separate metadata scores and exact provenance in [UPSTREAM.md](UPSTREAM.md).
- Readability text F1 is 87.82711% / 95.20557% / 78.47603% on LegoNews /
  ScrapingHub / WCXB. Selected Go/Rust extraction is 2.669 / 2.441 ms/page
  (1.09x); all-four means are 2.705 / 2.451 ms/page. Shared parsing is separate.
- Go-ReadabilityV2 remains 0.6.0. No source, module, dependency, fixture or tag
  changes are part of this documentation update; historical results are retained.