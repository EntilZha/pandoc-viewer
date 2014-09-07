pandoc-viewer
=============
This project is a quick hack that helps my workflow when working with Pandoc Markdown documents.

My goal was to be able to write Pandoc style markdown specifically to be able to write LaTeX inline
or block code within a markdown document. My preferred workflow would be vim with an auto updating
pdf viewer alongside whenever file changes are written (perhaps something more fancy in the future).

The code is setup as a Go project that watches for files changes. On file changes it executes this
chain of command
```
pandoc <filename>.md -o <filename>.pdf; open -a Preview; open -a MacVim
```
This compiles the pandoc document, brings Preview into focus so it pulls changes from disk, then opens
MacVim to change focus back. For reference, I also have this in my .bashrc which makes this more seamless
since by default MacVim from Terminal and Application launch are treated differently.

```
alias mvim="open -a MacVim"
```

Contributions are welcome if you are interested in extending this to more than just my own workflow
