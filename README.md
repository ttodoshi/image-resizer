# Image Resizer

## Console

### Usage

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

## Rest API

```
GET /
Params:
  - height           uint
        height in pixels
  - width            uint
        width in pixels
  - save-proportions bool
        save proportions 

Multipart form data
- "file"
```
