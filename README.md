# WriterSeeker [![CircleCI](https://circleci.com/gh/orcaman/writerseeker.svg?style=svg)](https://circleci.com/gh/orcaman/writerseeker) [![GoDoc](https://godoc.org/github.com/orcaman/writerseeker?status.svg)](https://godoc.org/github.com/orcaman/writerseeker)

WriterSeeker is the in-memory io.WriteSeeker implementation missing in the standard lib :-)


## Use-case
In serverless / PaaS environemnts there is usually no file system access - you cannot read or write files to the container you are running in. This means that if you are using a function or library in Go that expects a `File` type (which implements the `io.WriteSeeker` interface), you are pretty much screwed. WriterSeeker solves this by letting you write and seek inside an in-memory buffer.

## Usage Example

Let's say that you are using a library to generate PDF files. The library usually expects a `File` type to perform the writing to. You would create a `File` by using `os.Open` and then feed this to the `Write` function like so:

```go
fWrite, err := os.Create(outputPath)
if err != nil {
    return err
}

defer fWrite.Close()

err = pdfWriter.Write(fWrite)
```

With `WriterSeeker`, you do not need the file, just work in-memory:

```go
writerSeeker := &writerseeker.WriterSeeker{}
err = pdfWriter.Write(writerSeeker)
```

Now you can get a an `io.Reader` from the `writerSeeker` instance and boogie! for example, copy it's buffer to an `io.Writer`.

```
r := writerSeeker.Reader()
w := getWriter()
if _, err := io.Copy(w, r); err != nil {
 ...
 ...
 ...
}
```