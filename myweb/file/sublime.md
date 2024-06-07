---
date: 2024-04-08
authors:
  - kian
categories:
  - 文档
---

# **Sublime**
<!-- more -->

## 安装 Markdown 插件
Preferences -> Package Control -> Package Control:Install Package

安装 MarkdownEditing 和 MarkdownPreview 插件

## 配置 Markdown 插件
Preferences -> Key Bindings

```json
{ "keys": ["f1"], "command": "markdown_preview", "args": {"target": "browser", "parser":"markdown"} }
```

Preferences -> Package Settings -> Markdown Preview -> Settings

```json
{ "enable_autoreload": true }
```