# Console Image Resizer

## Build

```shell
make build
```

## Usage

simple usage

```shell
./bin/app <image-path1> <image-path2>
```

all pics from directory

```shell
find <dir-path> -maxdepth 1 -type f | xargs -I {} ./bin/app {}
```

### Help

```
Usage of ./bin/app:
  -h    show help
  -mh uint
        height in pixels (default 512)
  -mw uint
        width in pixels (default 512)
  -o string
        output directory (default "resized")
```