# TO-DO

- **First**, re-write the settings & configuration system from scratch! It's
  broken and messy and not worth trying to fix like this. Instead, actually
  architect how it should work, _then_ implement it.
- Refactor the directory structure processing
  - Implement zola-style structure instead
    - `zona init` command to populate the required files, _with_ defaults
      (unlike zola)
      - Interactive for setting values, also an option to create `.gitignore`
        with `public` in it.
    - `zona.yml` is **required** and should mark the root:
      - `templates`, `content`, `static`, `zona.yml`
      - multiple `zona.yml` files should be an error
      - if the folder containing `zona.yml` doesn't contain _exactly_ the
        expected directories and files, it's an error
    - Paths in page metadata should start at these folders
      - i.e. `(template|footer|header): name.html` → `root/templates/name.html`
      - `(style|icon): name.ext` → `root/static/name.ext`
    - Traverse `content` and `static` separately, applying different rules
      - everything in `static/**` should be directly copied
      - `content/**` should be processed
        - `*.md` converted, everything else copied directly
        - `./name.md` → ./name/index.html
        - Either `./name.md` or `./name/index.md` are valid, _together_ they
          cause an error!
  - What about markdown links to internal pages?
    - Relative links should be supported to play nice with LSP
      - in case of relative link, zona should attempt to resolve it, figuring
        out which file it's pointing to, and convert it to a `/` prefixed link
        pointing to appropriate place
      - so `../blog/a-post.md` → `/blog/a-post` where `/blog/a-post/index.html`
        exists
    - links from project root should also be supported
    - check link validity at build time and supply warning
    - _tl;dr_ all links should be resolved to the absolute path to that resource
      starting from the website root. that's the link that should actually be
      written to the HTML.
- Re-consider what `zona.yml` should have in it.
  - Set syntax highlighting theme here
    - a string that's not a file path: name of any built-in theme in
      [chroma](https://github.com/alecthomas/chroma)
    - path to `xml` _or_ `yml` file: custom theme for passing to chroma
      - if `xml`, pass directly
      - if `yml`, parse and convert into expected `xml` format before passing
  - Set website root URL here
  - toggle option for zona's custom image label expansion, image container div,
    etc, basically all the custom rendering stuff
- Syntax highlighting for code blocks
- Add `zona serve` command with local dev server to preview the site
- Both `zona build` and `zona serve` should output warning and error
- Write actual unit tests!
