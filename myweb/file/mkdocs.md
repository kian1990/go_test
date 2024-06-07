---
date: 2023-12-25
authors:
  - kian
categories:
  - 文档
---

# **MkDocs**
[官方文档](https://squidfunk.github.io/mkdocs-material){ .md-button }
<!-- more -->

## 安装

### 通过 `pip` 安装

``` bash
pip install mkdocs-material
```

### 修改配置文件 `mkdocs.yml`:

``` yaml
theme:
  name: material
```

## 创建新项目
```bash
mkdocs new [dir-name]
```

## 启动动态服务器
```bash
mkdocs serve
```

## 部署静态页面
```bash
mkdocs build
```

## 查看帮助
```bash
mkdocs -h
```

## 项目描述
```bash
    mkdocs.yml    # 配置文件
    docs/
        index.md  # 主页
        ...       # 其他 markdown 页面、图片和其他文件
```

## 许可

**MIT License**

Copyright (c) 2016-2024 Martin Donath

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to
deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NON-INFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.

