[简体中文](https://github.com/Xeron2000/gitTrend/blob/main/README.zh.md) | [English](https://github.com/Xeron2000/gitTrend/blob/main/README.md)

## 目录

- [介绍](#介绍)
- [安装](#安装)
- [使用](#使用)
- [贡献](#贡献)
- [许可证](#许可证)

## 介绍

![gt](https://raw.githubusercontent.com/Xeron2000/gitTrend/main/gt.png)

## 安装

### 使用构建好的二进制程序

见[Releases](https://github.com/Xeron2000/gitTrend/releases)

### 从源代码安装

- [Go](https://golang.org/dl/) 1.22.4+
- [Git](https://git-scm.com/downloads)

1. 克隆仓库

   ```
   git clone https://github.com/Xeron2000/gitTrend.git
   cd gitTrend
   ```

2. 安装依赖项

   ```
   go mod tidy
   ```

3. 构建二进制程序

   ```
   go build -o gt
   cp gt /usr/local/bin/
   ```

4. 运行项目

   ```
   gt scrape
   ```

## 使用

```
gt scrape -s en -l go -t daily
```

参数说明：

- `-s, --spoken`：指定口语语言代码（例如 `en` 表示英语）。
- `-l, --language`：指定编程语言（例如 `go` 表示 Go 语言）。
- `-t, --time`：指定时间范围（例如 `daily` 表示每日，`weekly` 表示每周，`monthly` 表示每月）。

## 贡献

欢迎贡献代码！请阅读以下指南：

1. Fork 本仓库
2. 创建一个新的分支 (`git checkout -b feature/AmazingFeature`)
3. 提交你的修改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开一个 Pull Request

## 许可证

该项目基于 MIT 许可证发布。详细信息请参阅 [LICENSE](LICENSE) 文件。
