# What is?

The name is self explanatory, but this program reads all the files inside a folder and then merges all the content into a final file.

## Why?

In my day-to-day life, i create a lot of _markdown_ files with notes that i've written during some meeting, or things i need to remember

## What did i learn?

I tried to apply goroutines but the initial version doesn't need them, the code is only using them to send every file name read to a channel, but it can also work without using channels.

- How to copy files in chunks without loading into memory
- Understand `io.Writer` and `io.Reader`
- Simple usage of `Cobra Cli`

**Time spent**: +/- 1 hour
