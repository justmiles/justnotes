package backend

var defaultContents = []byte(`# justNotes
A simple markdown editor powered by [Wails](https://wails.app/) and [EasyMDE](https://easymde.tk/).

## Getting started
- Click the eye icon at the top to preview your markdown
- Select the folder icon to open a markdown document

## Shortcuts

| Shortcut (Windows / Linux) | Shortcut (macOS) | Action                 |
| -------------------------- | ---------------- | ---------------------- |
| *Ctrl-'*                   | *Cmd-'*          | "toggleBlockquote"     |
| *Ctrl-B*                   | *Cmd-B*          | "toggleBold"           |
| *Ctrl-E*                   | *Cmd-E*          | "cleanBlock"           |
| *Ctrl-H*                   | *Cmd-H*          | "toggleHeadingSmaller" |
| *Ctrl-I*                   | *Cmd-I*          | "toggleItalic"         |
| *Ctrl-K*                   | *Cmd-K*          | "drawLink"             |
| *Ctrl-L*                   | *Cmd-L*          | "toggleUnorderedList"  |
| *Ctrl-P*                   | *Cmd-P*          | "togglePreview"        |
| *Ctrl-Alt-C*               | *Cmd-Alt-C*      | "toggleCodeBlock"      |
| *Ctrl-Alt-I*               | *Cmd-Alt-I*      | "drawImage"            |
| *Ctrl-Alt-L*               | *Cmd-Alt-L*      | "toggleOrderedList"    |
| *Shift-Ctrl-H*             | *Shift-Cmd-H*    | "toggleHeadingBigger"  |
| *F9*                       | *F9*             | "toggleSideBySide"     |

![](https://thumbs.gfycat.com/IncredibleRapidBluetonguelizard-size_restricted.gif)
`)
