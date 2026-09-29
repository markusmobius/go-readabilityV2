# Go-ReadabilityV2 Maintenance

Read these instructions before changing README.md, UPSTREAM.md, CHANGELOG.md,
benchmark claims or releases. Inspect current files and git status; preserve
unrelated edits and follow explicit user constraints.

## Library Identity

This is the supplied-HTML, core-only Go fork of Readeck/Mozilla Readability.
Fetching, CLI/server behavior and scheduling belong to callers. Describe the
actual pinned reference and deliberate differences, not complete equivalence
inferred from a finite test corpus. Standalone Mozilla Readability is not
Trafilatura's bundled readability-lxml. Removing Trafilatura fallback candidates
does not change this library's independent extraction API.

## Document Roles

- README.md: readable purpose, scope, usage and current quality/speed. Keep
  existing useful examples; do not turn it into an implementation or audit log.
- UPSTREAM.md: ancestry, exact source/dependency pins, compatibility boundaries,
  reproductions, hashes, coverage and remaining differences.
- CHANGELOG.md: dated/versioned user-visible changes and why they were made.
  Label documentation-only changes; preserve the meaning of historical entries.
- Release notes: the same changes, benchmark definitions and limitations as the
  README/changelog, with links to detailed evidence instead of copied logs.
- AGENTS.md: durable instructions, not current results or work-in-progress.

## Six-Repository Benchmark Contract

Follow the complete
[benchmark maintenance guide](https://github.com/markusmobius/content-extractor-benchmark/blob/master/AGENTS.md).
Coordinate go-domdistiller, rust-domdistiller, go-readabilityV2, rust-readability,
go-trafilatura and rust-trafilatura, not only this library's language pair.

1. Every README has `## Current Quality and Speed` with the same six-engine
	comparison from one completed shared-suite report. Use measured version labels.
2. Columns: `Extractor`, `Go Version`, `Rust Version`, `Go ms/page`,
	`Rust ms/page`, `Go/Rust`. Rows: Readability, DomDistiller, Trafilatura FAST.
	Display milliseconds/page to three decimals and Go/Rust ratios to two decimals.
3. Read structured JSON and compute ratios from unrounded means. The current
	shared protocol uses all four measured passes after one warmup. Never mix
	dates, environments, modes, means/medians or selected/all-pass aggregates.
4. Report parsing separately, charged once per language/page. Include decoding,
	normalization, DOM construction and any separately required Trafilatura tree.
	State corpus/counts, options, hardware, toolchains and included/excluded work.
5. Keep non-FAST Trafilatura as a separate comparison with its own report. A
	paired language ratio is not an isolated version speedup or request latency.
6. Report named-corpus F1 percentages to five decimals, retaining errors in the
	denominators. Do not average different scoring definitions. Matching text
	scores do not prove byte-identical HTML or metadata; preserve known differences.
7. Link immutable source reports and commits. Keep historical results clearly
	dated and old artifacts unchanged. Older 102-pass standalone measurements
	are not directly comparable to the newer shared-input timing boundary.
8. Documentation-only patches retain actual measured versions; do not relabel
	old measurements or rerun benchmarks just to change wording. Keep unrelated
	application-worker incidents and fallback traces out of this README.

## Editing and Publication

1. Identify the requested scope and evidence. Documentation is not permission
	to change extraction, parsing, dependency versions or application behavior.
2. Update the three documents according to their roles and coordinate the common
	benchmark section across all six repositories. Use plain, library-facing prose.
3. Check counts, percentages, units, ratios, version labels, option names and
	Markdown links. Compare the common sections and run `git diff --check`.
	Use existing example/tests and CI commands for relevant verification; disclose
	any gate not run and do not weaken independent reference assertions.
4. Obtain authorization before commits, pushes, tags, new package versions or
	release publication. Go documentation edits do not automatically require a
	new module version. Never move an already published tag.
5. Finalize documentation before an authorized release. Keep runtime and
	dependency pins unchanged for documentation-only work. Verify ordinary module
	download with checksum checking enabled and verify the GitHub release page.
	A successful tag push alone does not create a release page.
6. Coordinated Rust updates must account for immutable crates.io READMEs. New
	crate documentation requires an authorized patch version, finalized docs,
	archive inspection and published-byte verification, not a release-body edit.
7. Hash committed/remote bytes rather than assumed Windows worktree bytes;
	CRLF and LF differ. Preserve historical evidence and unrelated dirty files.
	Never print credentials or infer hosted CI success from publication success.
8. Report actual validation, publication versions/URLs and remaining limitations
	concisely. Detailed source receipts belong in UPSTREAM or benchmark reports.