# Build your own wc in go

![ccwc.png](./docs/ccwc.png)

## Intro

I began learning Go and wanted to work on a project to enhance my skills. I discovered this [build your own wc tool](https://codingchallenges.fyi/challenges/challenge-wc/) challenge and decided to take it on.

## How to run

Make sure you have Go [installed](https://go.dev/doc/install).

Then clone the repo and run `make build` in the root directory.  

Then you can write this command:

```bash 

# Without any flags will show the help message

> ./ccwc 
Usage of ./ccwc:
  -c    print the byte/s count
  -l    print the line count
  -m    print the character count
  -w    print the word count

# Show line count

> ./ccwc -l <file>
791 ./data/test.txt

# Show word count

> ./ccwc -w <file>
10859 ./data/test.txt

# Show character count

> ./ccwc -m <file>
61031 ./data/test.txt

# Show byte count

> ./ccwc -c <file>
61031 ./data/test.txt

# without any flags will show all the counts
# similar to `wc` command
> ./ccwc ./data/test.txt
791 10859 61031 ./data/test.txt

```

## Why Dockerfile

I was really to see how does a Go project Dockerfile looks like, turn out it's really simple.

Just build the image and run it:

```bash
docker build -t ccwc .
docker run -it ccwc # or with any flag like above
```

## Credits

I studied other codes who have completed this challenge in Go and other languages. Here are the list:

- [André Brandão](https://github.com/andrenbrandao/wc-tool)
- [PraveshGoyal](https://github.com/praveshdev3/wc-go)
- [Hruthik Reddy Yarala](https://github.com/yaralahruthik/cc-wc)

Also I used project like [lazygit](https://github.com/jesseduffield/lazygit) and [lazydocker](https://github.com/jesseduffield/lazydocker) to help me understand how to use Go in a large scale project.

