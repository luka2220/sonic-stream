# Sonic Stream

SonicStream is an ultra-fast, lightweight multimedia conversion tool that makes it effortless to transform images and other media between formats. Written in Go, it provides a simple HTTP API for uploading files, converting them, and retrieving the results.

Features:
- Lightning-Fast Conversions:
Quickly convert between popular image formats: PNG, JPEG, GIF, BMP, WebP, etc.
- Lightweight Backend: Powered by Go – minimal resource usage and easy deployment.
- Upload your file using multipart/form-data, specify the target format, and receive a converted file or a download link.
  - Ephemeral Storage
- Optional auto-cleanup after a set time (e.g., 24 hours) to keep storage usage low.
  - Extensible
- Add new file types, transformations, and advanced features (e.g., resizing, compression settings) as your project grows.
