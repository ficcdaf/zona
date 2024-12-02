# Zona

Zona is a tool for building a static website, optimized for lightweight blogs following minimalist design principles.

**Warning:** Zona has not yet reached **v1**. The `dev-*` branches of this repository contain the code -- however, there is no assurance of stability or functionality until the first release. Configuration and usage documentation will also be provided at this time.

## Table of Contents

- [Design Goals](#design-goals)
- [v1 Features](#v1-features)
- [Roadmap](#roadmap)
- [Contribution](#contribution)
- [Inspirations](#inspirations)

## Design Goals

Zona is intended to be easy-to-use. A user should be able to build a reasonably complex website or blog with only a directory of Markdown content and a single command, without needing to write any HTML or configuration. However, users should optionally have access to sensible and flexible configuration options, including writing HTML. The output of Zona should also be lightweight, retaining the smallest file sizes possible. These characteristics make Zona well-suited for both beginners and power users that wish to host a website on a service like Neocities or GitHub Pages.

## v1 Features

- Write pages purely in Markdown.
- Single-command build process.
- Lightweight output.
- Sensible default template and stylesheet.
- Configuration entirely optional, but very powerful.
- Site header and footer defined in Markdown.
- Declarative metadata per Markdown file.
- Automatically generated `Archive`, `Recent Posts`, and `Image Gallery` elements.
- Support for custom stylesheets, favicons, and page templates.

## Roadmap

- [ ] RSS/Atom feed generation.
- [ ] Image optimization & dithering.
- [ ] Windows, Mac, Linux releases.
- [ ] AUR package.
- [ ] Custom Markdown tags that expand to user-defined templates.
- [ ] Live preview local server.

## Contribution

Zona is a small project maintained by a very busy graduate student. If you want to contribute, you are more than welcome to submit issues and pull requests.

## Inspirations

- [Zoner](https://git.sr.ht/~ryantrawick/zoner)
- [Zonelets](https://zonelets.net/)

> Note: I am aware of `Zola`, and the similar name is entirely a coincidence. I have never used it, nor read its documentation, thus it is not listed as an inspiration.
