# gohack

[![Build Status](https://github.com/timandy/gohack/actions/workflows/build.yml/badge.svg)](https://github.com/timandy/gohack/actions)
[![Codecov](https://codecov.io/gh/timandy/gohack/branch/main/graph/badge.svg)](https://app.codecov.io/gh/timandy/gohack)
[![License](https://img.shields.io/github/license/timandy/gohack.svg)](https://github.com/timandy/gohack/blob/main/LICENSE)

> [English Version](README.md)

`gohack`提供了一种设计和源码，借助此你可以`读取`和`修改`协程结构体`runtime.g`的任意字段，甚至其他运行时未导出的结构的内容。

# 介绍

`gohack`通过汇编获取`runtime`包中的变量值，通过此变量可以进行反射获取到结构体的字段偏移量。

当要读取该类型的其他变量的字段时，通过另一个简化的汇编只获取地址，该方式开销小，并将地址偏移然后读写内容。

该库演示了如何通过汇编获取`runtime.g0`的值，并反射获取到关键字段的偏移量。

然后读取当前协程结构体`runtime.g`的地址，偏移指针地址，然后读写字段。

需要注意的是该方式开销极小，并且兼容未来的`go`版本，并且支持跨平台（`386`、`amd64`、`armv6`、`armv7`、`arm64`）。

## 使用

`gohack`没有提供任何公开的接口。

`gohack`的目标是提供一种思路和源码实现。

如果你觉得该仓库对你有所帮助，请 [Fork](https://github.com/timandy/gohack/fork) 和 [Star](https://github.com/timandy/gohack) 。

[routine](https://github.com/timandy/routine) 是一个 `tls` 库，由 `gohack` 提供技术支持。

# 支持网格

|             | **`darwin`** | **`linux`** | **`windows`** | **`freebsd`** |             |
|------------:|:------------:|:-----------:|:-------------:|:-------------:|:------------|
|   **`386`** |              |      ✅      |       ✅       |       ✅       | **`386`**   |
| **`amd64`** |      ✅       |      ✅      |       ✅       |       ✅       | **`amd64`** |
| **`armv6`** |              |      ✅      |               |               | **`armv6`** |
| **`armv7`** |              |      ✅      |               |               | **`armv7`** |
| **`arm64`** |      ✅       |      ✅      |               |               | **`arm64`** |
| **`ppc64`** |              |      ✅      |               |               | **`ppc64`** |
| **`s390x`** |              |      ✅      |               |               | **`s390x`** |
|             | **`darwin`** | **`linux`** | **`windows`** | **`freebsd`** |             |

✅：支持

# *许可证*

`gohack`是在 [Apache License 2.0](LICENSE) 下发布的。

```
Copyright 2021-2022 TimAndy

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
