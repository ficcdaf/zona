# TODO

- Fix the relative URL situation in the headers
  - The link that's defined in header file should be relative to root, not the
    page being processed.
  - How to handle this?
- Syntax highlighting for code blocks
- Implement zola-style directory structure
  - `templates`, `content`, `static`?
  - Paths in page metadata should start at these folders
  - What about markdown links to internal pages?
  - Steal Zola's syntax?
    - Link starting with `@` → `content/`
- Implement zola-style sections with `_index.md`

## Thoroughly test directory processing

- Is the pagedata being constructed as expected?
- Is the processmemory struct working as expected?

NOTE: I should really write these as actual tests, not just running tests on my
testing directory myself.
