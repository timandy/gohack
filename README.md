# gohack

[![Build Status](https://github.com/timandy/gohack/actions/workflows/build.yml/badge.svg)](https://github.com/timandy/gohack/actions)
[![Codecov](https://codecov.io/gh/timandy/gohack/branch/main/graph/badge.svg)](https://app.codecov.io/gh/timandy/gohack)
[![Go Report Card](https://goreportcard.com/badge/github.com/timandy/gohack)](https://goreportcard.com/report/github.com/timandy/gohack)
[![License](https://img.shields.io/github/license/timandy/gohack.svg)](https://github.com/timandy/gohack/blob/main/LICENSE)

> [中文版](README_zh.md)

`gohack` provides a design and source code with which you can `read` and `modify` arbitrary fields of the coroutine struct `runtime.g`, even the contents of other structs not exported by the runtime.

# Introduce

`gohack` obtains the variable value in the `runtime` package through assembly, through which the field offset of the structure can be obtained by reflection.

When the fields of other variables of this type are to be read, only the address is obtained through another simplified assembly, which has low overhead, and offsets the address and then reads and writes the content.

This library demonstrates how to get the value of `runtime.g0` by assembly, and get the offset to the key field by reflection.

Then read the address of the current coroutine structure `runtime.g`, offset the pointer address, and then read and write the fields.

It should be noted that this method has minimal overhead, is compatible with future `go` versions, and supports cross-platform (`386`, `amd64`, `armv6`, `armv7`, `arm64`, `loong64`, `mips`, `mipsle`, `mips64`, `mips64le`, `ppc64`, `ppc64le`, `riscv64`, `s390x`, `wasm`).

## Usage

`gohack` does not provide any exposed interface.

The goal of `gohack` is to provide an idea and source code implementation.

If you feel this repository is helpful to you, please [Fork](https://github.com/timandy/gohack/fork) and [Star](https://github.com/timandy/gohack).

[routine](https://github.com/timandy/routine) is a `tls` library, powered by `gohack`.

# Support Grid

|                | **`darwin`** | **`linux`** | **`windows`** | **`freebsd`** | **`js`** |                |
|---------------:|:------------:|:-----------:|:-------------:|:-------------:|:--------:|:---------------|
|      **`386`** |              |      ✅      |       ✅       |       ✅       |          | **`386`**      |
|    **`amd64`** |      ✅       |      ✅      |       ✅       |       ✅       |          | **`amd64`**    |
|    **`armv6`** |              |      ✅      |               |               |          | **`armv6`**    |
|    **`armv7`** |              |      ✅      |               |               |          | **`armv7`**    |
|    **`arm64`** |      ✅       |      ✅      |               |               |          | **`arm64`**    |
|  **`loong64`** |              |      ✅      |               |               |          | **`loong64`**  |
|     **`mips`** |              |      ✅      |               |               |          | **`mips`**     |
|   **`mipsle`** |              |      ✅      |               |               |          | **`mipsle`**   |
|   **`mips64`** |              |      ✅      |               |               |          | **`mips64`**   |
| **`mips64le`** |              |      ✅      |               |               |          | **`mips64le`** |
|    **`ppc64`** |              |      ✅      |               |               |          | **`ppc64`**    |
|  **`ppc64le`** |              |      ✅      |               |               |          | **`ppc64le`**  |
|  **`riscv64`** |              |      ✅      |               |               |          | **`riscv64`**  |
|    **`s390x`** |              |      ✅      |               |               |          | **`s390x`**    |
|     **`wasm`** |              |             |               |               |    ✅     | **`wasm`**     |
|                | **`darwin`** | **`linux`** | **`windows`** | **`freebsd`** | **`js`** |                |

✅: Supported

# *License*

`gohack` is released under the [Apache License 2.0](LICENSE).

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
