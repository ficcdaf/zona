# Zona

[Zona](https://sr.ht/~ficd/zona/) is a tool for building a static website,
optimized for lightweight blogs following minimalist design principles. The
project is hosted on [sourcehut](https://sr.ht/~ficd/zona/) and mirrored on
[GitHub](https://github.com/ficcdaf/zona). You are welcome to open an issue on
GitHub or send a message to the
[mailing list](https://lists.sr.ht/~ficd/zona-devel).

<!-- prettier-ignore-start -->

> [!NOTE]
> Zona is currently in development. The `main` branch of this repository does
> not yet contain the software. The `dev-stable` branch contains the code used
> to generate [ficd.ca](https://ficd.ca) -- although the program is undocumented
> and missing features, so please proceed at your own risk. The `dev` branch
> contains the latest development updates and is not guaranteed to be functional
> (or even compile) at any given commit. Kindly note that the commit history
> will be cleaned up before the program is merged into the `main` branch.

<!-- prettier-ignore-end -->

## Design Goals

Zona is intended to be easy-to-use. A user should be able to build a reasonably
complex website or blog with only a directory of Markdown content and a single
command, without needing to write any HTML or configuration. However, users
should optionally have access to sensible and flexible configuration options,
including writing HTML. The output of Zona should also be lightweight, retaining
the smallest file sizes possible. These characteristics make Zona well-suited
for both beginners and power users that wish to host a website on a service like
Neocities or GitHub Pages.

## Features Implemented

- Write pages purely in Markdown.
- Single-command build process.
- Lightweight output.
- Sensible default template and stylesheet.
- Configuration file.
- Internal links preserved.
- Custom image element parsing & formatting.
- Site header and footer defined in Markdown.
- YAML frontmatter support.

## Planned Features

- Automatically treat contents of `posts/` directory as blog posts.
- Automatically generated `Archive`, `Recent Posts`, and `Image Gallery`
  elements.
- Support for custom stylesheets, favicons, and page templates.
- Image optimization and dithering.
- Custom markdown tags that expand to user-defined templates.
- Live preview server.
- Robust tests.

## Inspirations

- [Zoner](https://git.sr.ht/~ryantrawick/zoner)
- [Zonelets](https://zonelets.net/)
