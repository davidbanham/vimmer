# vimmer

A quick and dirty Go utility to open a series of files in vim.

In response to [this tweet](https://twitter.com/jonrowe/status/401147335222108160)

Currently uses mvim since I couldn't immediately figure out how to run terminal vim through golang's exec.

Also doesn't wait for the last process to exit before starting the next, since mvim seems to return an exit code immediately.

This is the product of about 10 minutes work, and only the second thing I've written in Go. Be gentle.

### use

    ls | xargs bin/mvimmer
