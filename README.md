TelSky
---

[![asciicast](https://asciinema.org/a/uUYrFZWRKqioSWj9jI3bdWKNT.svg)](https://asciinema.org/a/uUYrFZWRKqioSWj9jI3bdWKNT)

## Operating Proof of Concept

* Write credentials to `XDG_CONFIG` `gosky account create-session "your-handle.dev" "your-app-password" > ~/.config/telsky/client.auth`
* go install && telsky
* Cross your fingers

## Dev Notes

Let's use [charm](https://github.com/charmbracelet/bubbles?tab=readme-ov-file#list) to build out a list view of the main timeline.

This [HackerNews Example](https://github.com/bensadeh/circumflex/tree/main) would be a good model to copy. We need a basic list view
with a few actions.

We can use a [modal and other styles](https://github.com/charmbracelet/lipgloss) to add basic text previews with something like this: https://github.com/bensadeh/circumflex/blob/main/reader/reader.go#L15

For images, we can use sixels: https://github.com/mattn/go-sixel or a fallback with something like this: https://github.com/qeesung/image2ascii

Maybe we can use nerdfonts for basic icons?

All the [charm](https://charm.sh/) is here

