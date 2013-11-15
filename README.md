# vimmer

A quick and dirty Go utility to open a series of files in vim.

In response to [this tweet](https://twitter.com/jonrowe/status/401147335222108160)

### use

    ls | xargs bin/vimmer

This opens a sequence of files in vim, opening the next one after you exit the editor.

It can also do mvim

    ls | xargs bin/mvimmer

But this one doesn't wait for the last process to exit before starting the next, since mvim seems to return an exit code immediately.

This is the product of about 10 minutes work, and only the second thing I've written in Go. Be gentle.

