# Generic-Go

在 Go 语言早期版本中，由于缺乏泛型，开发者在处理集合类型如切片、列表、队列等时，往往需要编写大量类型特定的代码。这不仅增加了开发工作量，也使得代码难以维护。Go 1.18 的泛型特性解决了这一问题，`Generic-Go` 正是基于这一特性，旨在为 Go 开发者提供一个现代、高效的工具库。

`Generic-Go` 是一个为 Go 语言设计的泛型工具库，它充分利用了 Go 1.18 版本引入的泛型特性，为开发者提供了一套类型安全、灵活且易于使用的编程接口。泛型是编程语言中一种强大的工具，它允许编写更通用、更灵活的代码，减少样板代码，并提高代码的复用性。



## 目标

- [x]  **切片**
    - [x]  添加
    - [x]  删除
    - [x]  缩容
    - [x]  聚合运算
        - [x]  求最大值
        - [x]  求最小值
        - [x]  求和
    - [x]  逆转
    - [x]  查找
    - [x]  索引
    - [x]  映射
    - [x]  包含
    - [x]  集合运算
        - [x]  交集
        - [x]  并集
        - [x]  差集
        - [x]  对称差集
- [ ]  **List**
    - [x]  ArrayList
    - [x]  LinkedList
    - [x]  ConcurrentList
    - [ ]  SkipList
- [ ]  **队列**
    - [ ]  基于 ArrayList
    - [ ]  基于 LinkedList
    - [ ]  优先级队列
- [ ]  **栈**
- [ ]  **堆**
- [ ]  **Map**
    - [ ]  基于 map 的 HashMap 封装
    - [ ]  LinkedMap
- [ ]  **树**
    - [ ]  红黑树
    - [ ]  基于红黑树的 TreeMap 和 TreeSet
- [ ]  **Set**
    - [ ]  HashSet
    - [ ]  TreeSet
- [ ]  **跳表**
    - [ ]  基于跳表的有序 SortedSet
- [ ]  **并发队列**
    - [ ]  并发队列
    - [ ]  并发阻塞队列
    - [ ]  并发阻塞优先级队列
- [ ]  **利用 `Generic-Go` 实现本地缓存**
    - [ ]  适配 Redis



## 安装

`Generic-Go` 作为库使用，可以通过 Go Modules 直接引入到您的项目中。使用以下命令：

```bash
go get github.com/HJH0924/Generic-Go
```



## 开发

拉取代码

```shell
git clone https://github.com/HJH0924/Generic-Go.git
```

配置环境

```shell
make setup
```



## 使用方法

以下是如何使用 `Generic-Go` 中的切片添加元素的功能示例，更多示例代码可查看 `examples` ：

```go
package main

import (
	"fmt"
	"github.com/HJH0924/Generic-Go/internal/slice"
)

func ExampleSliceAdd() {
	src := []int{3, 5, 7}
    // 使用 Generic-Go 的 Add 函数向切片添加元素
	src, err := slice.Add[int](src, 0, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(src)
}

func main() {
	ExampleSliceAdd()
}
```



## 构建和测试

`Makefile` 提供了多个目标来简化开发流程：

- `make bench`：运行基准测试。
- `make ut`：执行单元测试。
- `make setup`：执行项目设置脚本。
- `make fmt`：格式化代码。
- `make lint`：运行静态代码分析。
- `make tidy`：整理模块依赖。
- `make check`：执行代码检查流程。

运行 `make` 命令加上目标名称来执行相应任务。



## 补充

>   如果执行 `make check` 提示找不到 `goimports` 命令，则可能需要添加环境变量
>
>   执行下面安装 `goimports` 命令可能会将 `goimports` 安装在你的家目录下
>
>   ```shell
>   go install golang.org/x/tools/cmd/goimports@latest
>   ```
>
>   则需要添加环境变量
>
>   ```shell
>   export PATH=$PATH:$HOME/go/bin
>   ```
