# PingKratos

A simple and fast Ping service package for the Kratos framework, supporting both gRPC and HTTP protocols.

## CHINESE README

[中文说明](README.zh.md)

## Features

- ✅ **Simple Design** - Clean and straightforward service structure
- ✅ **Both Protocols** - Supports both gRPC and HTTP protocols
- ✅ **Native Support** - Seamless integration with Kratos framework
- ✅ **Built-in Tests** - Comprehensive test coverage included
- ✅ **Modern Schema** - Uses Protocol Buffers for service definition
- ✅ **Zero Config** - Out-of-the-box Ping service implementation

## Quick Start

### Installation

```bash
go get github.com/orzkratos/pingkratos
```

### Usage

See [TEST](serverpingkratos/ping_test.go) as demo.

## Dependencies

### Core Dependencies
- `github.com/go-kratos/kratos/v2` - Kratos framework
- `google.golang.org/grpc` - gRPC support
- `google.golang.org/protobuf` - Protocol Buffers

### Test Dependencies
- `github.com/stretchr/testify` - Testing framework
- `github.com/orzkratos/zapkratos` - Zap logging integration
- `github.com/yyle88/*` - Utilities

---

## License

MIT License. See [LICENSE](LICENSE).

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repo on GitHub (using the webpage interface).
2. Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. Navigate to the cloned project (`cd repo-name`)
4. Create a feature branch (`git checkout -b feature/xxx`).
5. Stage changes (`git add .`)
6. Commit changes (`git commit -m "Add feature xxx"`).
7. Push to the branch (`git push origin feature/xxx`).
8. Open a pull request on GitHub (on the GitHub webpage).

Please ensure tests pass and include relevant documentation updates.

---

## Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

If you find this package valuable, give me some stars on GitHub! Thank you!!!

**Thank you for your support!**

**Happy Coding with this package!** 🎉

Give me stars. Thank you!!!
