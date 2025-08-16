# PingKratos

简单快速的 Kratos 框架 Ping 服务包，同时支持 gRPC 和 HTTP 协议类型。

## 英文文档

[ENGLISH README](README.md)

## 功能特性

- ✅ **简洁设计** - 清晰直观的服务结构
- ✅ **双协议支持** - 同时支持 gRPC 和 HTTP 协议
- ✅ **原生集成** - 与 Kratos 框架无缝集成
- ✅ **内置测试** - 包含完整的测试覆盖
- ✅ **现代架构** - 使用 Protocol Buffers 定义服务
- ✅ **无需配置** - 开箱即用的 Ping 服务实现

## 快速开始

### 安装

```bash
go get github.com/orzkratos/pingkratos
```

### 使用方法

参见 [测试文件](serverpingkratos/ping_test.go) 作为示例。

## 依赖项

### 核心依赖
- `github.com/go-kratos/kratos/v2` - Kratos 框架
- `google.golang.org/grpc` - gRPC 支持
- `google.golang.org/protobuf` - Protocol Buffers

### 测试依赖
- `github.com/stretchr/testify` - 测试框架
- `github.com/orzkratos/zapkratos` - Zap 日志集成
- `github.com/yyle88/*` - 实用工具包

---

## 许可证类型

项目采用 MIT 许可证，详情请参阅 [LICENSE](LICENSE)。

---

## 贡献新代码

非常欢迎贡献代码！贡献流程：

1. 在 GitHub 上 Fork 仓库 （通过网页界面操作）。
2. 克隆Forked项目 (`git clone https://github.com/yourname/repo-name.git`)。
3. 在克隆的项目里 (`cd repo-name`)
4. 创建功能分支（`git checkout -b feature/xxx`）。
5. 添加代码 (`git add .`)。
6. 提交更改（`git commit -m "添加功能 xxx"`）。
7. 推送分支（`git push origin feature/xxx`）。
8. 发起 Pull Request （通过网页界面操作）。

请确保测试通过并更新相关文档。

---

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉

Give me stars. Thank you!!!
