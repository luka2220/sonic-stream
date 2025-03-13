# Sonic Stream

## TODOs

- [ ] Design api routes
- [ ] Build functionality for converting medias file types
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
