# randStr

`randStr` 函数是这个随机字符串生成器的核心，它负责根据给定的字符集、所需字符串的长度以及掩码位数来生成随机字符串。下面详细解释这个函数的工作原理：



## 函数签名

```go
func randStr(length int, charset string) string
```
- `length`：生成的随机字符串的长度。
- `charset`：字符集，一个字符串，从中选择字符来构建随机字符串。



## 工作原理

1. **掩码计算**：
    - `idxBits` 表示随机数中用于索引字符集的位数。例如，如果 `idxBits` 是 6，那么掩码就是 `0111111`（二进制），这意味着每次从随机数的低 6 位中提取一个索引值。
    - `idxMask` 是通过 `1<<idxBits - 1` 计算得到的掩码，用于从随机数中提取出用于索引的位。

2. **随机数缓存**：
    - 使用 `rand.Int63()` 生成一个 64 位的随机整数 `cache`，这个整数将被用来生成多个索引值。
    - `remain` 用于跟踪 `cache` 中还剩下多少组 `idxBits` 位的随机数可以被提取。

3. **生成随机字符串**：
    - 创建一个长度为 `length` 的字节切片 `result`，用于存储生成的随机字符串。
    - 循环 `length` 次，每次循环生成一个随机字符：
        - 如果 `remain` 为 0，表示 `cache` 中的随机位数已经用完，需要重新生成一个新的 64 位随机整数。
        - 使用掩码 `idxMask` 从 `cache` 的低 `idxBits` 位中提取出一个索引值 `randIndex`。
        - 检查 `randIndex` 是否在字符集 `charset` 的长度范围内，如果是，则从 `charset` 中取出对应的字符作为随机字符串的一部分。
        - 将 `cache` 向右移动 `idxBits` 位，以便在下一次循环中使用下一组随机位。
        - 每次循环后，`remain` 减 1，直到 `cache` 中的所有随机位数都被用完。

4. **返回结果**：
    - 循环结束后，`result` 切片中存储了生成的随机字符串，将其转换为字符串并返回。



## 代码示例

```go
func randStr(length int, charset string) string {
	idxBits := 0
	charsetSize := len(charset)
	for charsetSize > (1<<idxBits)-1 {
		idxBits++
	}

	idxMask := (1 << idxBits) - 1
	remain := 63 / idxBits
	cache := rand.Int63()
	res := make([]byte, length)

	for i := 0; i < length; {
		if remain == 0 {
			cache, remain = rand.Int63(), 63/idxBits
		}

		if randIdx := int(cache & int64(idxMask)); randIdx < charsetSize {
			res[i] = charset[randIdx]
			i++
		}

		cache >>= idxBits
		remain--
	}

	return string(res)
}
```



## 关键点

- **性能优化**：通过缓存随机数 `cache`，减少了调用随机数生成函数的次数，提高了性能。
- **随机性**：使用 `rand.Int63()` 生成的随机数保证了随机性。
- **灵活性**：通过调整 `idxBits` 的值，可以控制随机索引的分布，从而影响随机字符串的生成。

这种方法有效地利用了有限的随机数资源，生成了指定长度的随机字符串，同时保持了较高的性能和良好的随机性。
