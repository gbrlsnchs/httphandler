# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [0.3.1] - 2017-11-11
### Added
- Benchmark flag.

### Fixed
- goimports installation missing.

## [0.3.0] - 2017-11-11
### Added
- A writing method to the Handler.

### Changed
- How Content-Type header is set.
- Example.
- "Content" method from Responder now is called "Body".
- Enhanced the writing function method performance.

### Removed
- The stand-alone "write" function.

## [0.2.0] - 2017-10-31
### Added
- MsgPack support.
- Responder interface.

### Changed
- The writing function visibility to private. 
- Update this file to use "changelog" in lieu of "change log".

### Fixed
- Response header being set twice.
- Response when handler is nil.

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

[0.3.1]: https://github.com/gbrlsnchs/httphandler/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/gbrlsnchs/httphandler/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/gbrlsnchs/httphandler/compare/v0.1.0...v0.2.0
