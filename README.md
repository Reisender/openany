# OpenAny

This is a little utility to mimic os.Open("/some/file") but have it be able
to open files for more sources.

- local filesystem
- aws s3
- http/https

The goal is to make it easy to implement your own backend and have the
url scheme determine which backend to use.

### usage

Here are some examples of how you could use this in your golang code.

Open a local file.
```golang
readcloser, err := openany.Open("/some/local/file")
```

Open a URL.
```golang
readcloser, err := openany.Open("https://some/url")
```

Open from AWS S3.
```golang
readcloser, err := openany.Open("s3://bucket/some/object/key")
```
