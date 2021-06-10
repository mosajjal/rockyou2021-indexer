# rockyou2021-indexer
Simple Index and search tool for rockyou2021.txt (or any huge file)

Disclaimer: I'm not a very efficient coder. This code is the best I could come up with to solve this problem for myself.

The purpose of these two scripts is to use the filesystem as an indexing tool for a huge file like rockyou2021 (8.5 billion records) so we can make fast lookups.
The way I approached this problem was to create directories based on the first/second/third character ASCII codes. The code creates folder `indexed` in the directory it's being run in.
After that, it starts to read any huge file from stdin (`rockyou2021` for example) and puts it in different `list` files inside those subdirectories.

```
├── indexed
│   ├── 100
│   │   ├── 100
│   │   │   ├── 100
│   │   │   │   └── list
│   │   │   ├── 101
│   │   │   │   └── list
│   │   │   ├── 102
│   │   │   │   └── list
```

It does a very basic split of the file into separate `list`s and breaks our lookup time into a fraction of the original `grep` time by thousands. Keep in mind, the `index` process will take
hours to complete in large files.
## HOWTO

### build the binaries

This should be fairly straightforward since I'm not using any 3rd party libraries:

```go
go build -o indexer.bin indexer.go
go build -o lookup.bin lookup.go
```

### Download the Rockyou2021 text file

Download the big file using the following link and unzip:

```
magnet:?xt=urn:btih:JEQMEEFTBXT35RJ3GUTGXU7HP3HBU5P6&dn=rockyou2021.txt%20dictionary%20from%20kys234%20on%20RaidForums&tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A6969%2Fannounce
```

in your folder, you'll have something like this:

```
.
├── indexer.bin
├── indexer.go
├── LICENSE
├── lookup.bin
├── lookup.go
├── README.md
└── rockyou2021.txt
```

### Start the indexing process

First, make sure you have enough space for the massive number of files and folders being created. It shouldn't take more space that `rockyou2021.txt`. So you should have enough space for another copy of `rockyou2021.txt` at the minimum.

WARNING: this will take literally hours to complete. for me it took around 8 hours.

```bash
cat rockyou2021.txt | ./indexer.bin
```

if you'd like to see the progress as it's happening, use the great command line tool, `pv`:

```bash
cat rockyou2021.txt | pv | ./indexer.bin
```

after the process is finished, you'll have the folder `indexed` created in the same folder and you can look up any password you'd like very quickly.

### Lookup tool

The lookup tool has the `indexed` folder in the current directory hardcoded, so if you have moved the `indexed` folder, you need to change the variable in `lookup.go` and re-compile.

To lookup an entry or a list of entries, you can use the following:

```bash
$ ./lookup password123
password123: true
$ echo "password123" | ./lookup
password123: true
```

you can also `cat` your input list and get a list in `stdout`.