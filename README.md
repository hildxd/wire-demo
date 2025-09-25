# wire 使用方式

## 1. 安装 wire

```bash
go install github.com/google/wire/cmd/wire@latest
```

安装完后执行 `wire` 如果是命令没找到需要添加 gobin 到环境变量内

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

## 2. 项目结构

```
go-wire/
├── .git/
├── .gitignore
├── README.md           # 项目说明文档
├── go.mod             # Go 模块定义
├── go.sum             # Go 模块依赖校验和
├── main.go            # 程序入口文件
├── wire_gen.go        # Wire 自动生成的依赖注入代码
├── wrie.go            # Wire 配置文件
└── pkg/               # 包目录
    ├── event/
    │   └── event.go   # 事件相关功能
    ├── greeting/
    │   └── greeting.go # 问候语相关功能
    └── message/
        └── message.go # 消息相关功能
```

- 项目里有三个包 `event`, `greeting`, `message`, 依赖关系是 event -> greeting -> message，如果不用依赖注入就会很麻烦, 增加一些初始化的冗余代码
- 使用 wire， 需要你在根目录创建个 `wire.go` 文件  
   a. package 一定是 main  
   b. 顶部需要有以下代码注释  
   `golang
      //go:build wireinject
      // +build wireinject
      `
  b. 暴露的初始化方法需要是大写字母开头  
   c. 引入相关包
  d. 执行 `wire` 即可，生成了 `wire_go.go` 就成功了
