# Sonic Stream

## TODOs

- [ ] Set up CI/CD
- [ ] Design api routes
- [ ] Build functionality for converting medias file types
- [ ] Build the image API endpoint
- [x] Current file limit (image, video, audio) <= 100mb
- [x] Think about how files are going to be sent over the wire (whole, chunks, stream) (whole file for now...)

### API Routes

- Bundled static files from vite will be sent through /
- Base route for API is going to be /api

#### API Routes - /api

- POST => /api/image/ (image file to be sent through form data)
  - Form data should include: file, image format type to convert
- POST => /api/video/ (video file to be sent through form data)
  - Form data should include: video file, video format type to convert
- POST => /api/audio/ (audio file to be sent through form data)
  - Form data should include: audio file, audio format type to convert

#### Immage Route - POST -> /api/image

TODOS:

- [ ] Throw/handle error for invalid file type
- [ ] Abstract sending json response
- [ ] Unit tests

The max image file size is 150kb
Note: The go standard library supports image file encoding and decoding for multiple formats

- Structure of form-data key-value (http client request body)

  - file: (uploaded file from client) -> file
  - generate: (image file type to generate) -> string

- Response to endpoint:
  - For a successful response the client will need to make a second get request to the "downloaded_url" to download the stored file
  - To acomplish this the converted filw will be stored in the db with a short uuid
  - The POST reqest from /api/image responsed with the dowload url with the file uuid
  - The client will then use that route i.e /api/image/download/file_id to get the downloaded converted file

```json
{
  "download_url": "https://server.com/files/convereted/img.png"
}
```

- What Image file types take accept, each one will need to be parsed differently base on byte structure?

  - .png
  - .jpeg
  - .gif
  - .bmp
  - .webp

- Library image format conversion:
  1. PNG (import \_ "image/png")
  2. JPEG (import \_ "image/jpeg")
  3. GIF (import \_ "image/gif")
     • BMP: import _ "golang.org/x/image/bmp"
     • Provides a decoder for BMP. Encoding support is more limited.
     • WebP: import _ "golang.org/x/image/webp"
     • Provides a decoder for WebP. A built-in encoder is not included, so may need third-party or C bindings for WebP encoding.

### File Storage

My current idea of storage, since this is just a side project, is to store in an sqlite db with a timestamp, and have some cron job clean up and remove files that are a day or more old. I will have the cron job run once every day. If this project needs scaling and needs a more concurrent approach I will forward requests to a message queue, which then get fed out to read and write files to an aws s3 bucket
