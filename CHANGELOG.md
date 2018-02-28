# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added
- Runtime error handler.
- `internal` package with dummy data.
- Custom error handling support.

### Changed
- Refactored tests.
- `RuntimeErrorFunc` now accepts an `error` object as parameter.

### Removed
- Logger function support.
- Mock-up test files.
- `Handler.ServeHTTP` doesn't set a `Content-Type` header anymore.

### Changed
- `Error` now implements both `error` and `Responder`.

## [0.6.0] - 2018-01-29
### Changed
- `LoggerFunc` is now `ErrorLoggerFunc`.
- `DefaultLoggerFunc` is now `DefaultErrorLoggerFunc`

## [0.5.1] - 2018-01-29
### Removed
- `response.go` file, as it is now useless.

## [0.5.0] - 2018-01-29
### Added
- `Responder` interface for HTTP responses.
- `example_test.go` to README file as snippet.

### Changed
- README file tab length (through EditorConfig) for better printing Go example.

### Fixed
- Logger function not being called when an `Error` occurs.

### Removed
- `Response` struct for HTTP responses.

## [0.4.0] - 2018-01-24
### Changed
- `Handler` structure (marshaller functions are now pluggable).

### Fixed
- Panicking when response is `nil`
- Some parts of the changelog were changed to use code marking.

### Removed
- MsgPack support (its marshaller function can now be plugged in).

## [0.3.1] - 2017-11-11
### Added
- Benchmark flag.

### Fixed
- `goimports` installation missing.

## [0.3.0] - 2017-11-11
### Added
- A writing method to `Handler`.

### Changed
- How Content-Type header is set.
- Example.
- `Responder.Content` is now called `Responder.Body`.
- Enhanced the writing function method performance.

### Removed
- The stand-alone `write` function.

## [0.2.0] - 2017-10-31
### Added
- MsgPack support.
- Responder interface.

### Changed
- The writing function visibility to private. 
- Update this file to use "changelog" in lieu of "change log".

### Fixed
- Response header being set twice.
- Response when handler is `nil`.

### Removed
- Global header.
- Global error code status and message.

## 0.1.0 - 2017-09-21
### Added
- This changelog file.
- README file.
- MIT License.
- Travis CI configuration file and scripts.
- Git ignore file.
- Editorconfig file.
- This package's source code, including examples and tests.
- Go dep files.

[Unreleased]: https://github.com/gbrlsnchs/httphandler/compare/v0.6.0...HEAD
[0.6.0]: https://github.com/gbrlsnchs/httphandler/compare/v0.5.1...v0.6.0
[0.5.1]: https://github.com/gbrlsnchs/httphandler/compare/v0.5.0...v0.5.1
[0.5.0]: https://github.com/gbrlsnchs/httphandler/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/gbrlsnchs/httphandler/compare/v0.3.1...v0.4.0
[0.3.1]: https://github.com/gbrlsnchs/httphandler/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/gbrlsnchs/httphandler/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/gbrlsnchs/httphandler/compare/v0.1.0...v0.2.0
