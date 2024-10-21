# 基于 Redis + 滑动窗口 实现的限流器

## Redis 命令解释

### ZADD
Redis Zadd 命令用于将一个或多个成员元素及其分数值加入到有序集当中。

如果某个成员已经是有序集的成员，那么更新这个成员的分数值，并通过重新插入这个成员元素，来保证该成员在正确的位置上。

分数值可以是整数值或双精度浮点数。

如果有序集合 key 不存在，则创建一个空的有序集并执行 ZADD 操作。

当 key 存在但不是有序集类型时，返回一个错误。

注意： 在 Redis 2.4 版本以前， ZADD 每次只能添加一个元素。

#### 语法
```shell
ZADD KEY_NAME SCORE1 VALUE1.. SCOREN VALUEN
```

#### 返回值
被成功添加的新成员的数量，不包括那些被更新的、已经存在的成员。

#### 示例
```shell
redis 127.0.0.1:6379> ZADD salary 2000 "tom"
(integer) 1
redis 127.0.0.1:6379> ZADD salary 3500 "peter"
(integer) 1
redis 127.0.0.1:6379> ZADD salary 5000 "jack" 3000 "Tvux"
(integer) 2
redis 127.0.0.1:6379> zrange salary 0 -1 withscores # 显示有序集内所有成员及其 score 值
1) "tom"
2) "2000"
3) "Tvux"
4) "3000"
5) "peter"
6) "3500"
7) "jack"
8) "5000"
```



### ZCOUNT

Redis Zcount 命令用于计算有序集合中**指定分数区间**的成员数量。

#### 语法
```shell
ZCOUNT KEY_NAME min max
```

#### 返回值
分数值在 min 和 max 之间的成员的数量。

#### 示例
```shell
redis 127.0.0.1:6379> ZCOUNT salary 2000 5000
(integer) 4
```



### ZCARD

Redis Zcard 命令用于计算集合中元素的数量。

#### 语法

```shell
ZCARD KEY_NAME
```

#### 返回值

当 key 存在且是有序集类型时，返回有序集的基数。 当 key 不存在时，返回 0 。

#### 示例

```shell
redis 127.0.0.1:6379> ZCARD salary
(integer) 4
```

>    `ZCOUNT key -inf +inf` 和 `ZCARD key` 在某种程度上是等价的，但它们在行为上有细微的差别。
>
>   1.  **`ZCOUNT key -inf +inf`**:
>       -   `ZCOUNT` 命令用于计算有序集合中，分数（score）在给定区间内的成员数量。
>       -   当你使用 `-inf +inf` 作为区间时，它意味着你正在计算集合中所有成员的数量，因为 `-inf` 表示负无穷大，`+inf` 表示正无穷大，所以这个区间涵盖了所有可能的分数。
>   2.  **`ZCARD key`**:
>       -   `ZCARD` 命令用于返回有序集合中的成员数量。
>       -   它直接返回集合中的成员总数，而不需要指定分数区间。
>
>   **尽管两者都返回有序集合中的成员数量，但 `ZCOUNT` 命令在执行时会检查分数区间，而 `ZCARD` 则直接返回计数。在实际使用中，如果你只是想知道集合中的成员数量，使用 `ZCARD` 会更高效，因为它不需要进行分数区间的检查。**
>
>   在限流器的上下文中，如果你的目的是检查某个 IP 在当前时间窗口内的请求次数是否超过了阈值，并且你已经确保了时间窗口内的请求都被记录在了有序集合中，那么使用 `ZCARD` 会更合适，因为它直接返回了集合的大小，而不需要额外的区间检查。
>
>   总结来说，如果你不需要考虑分数区间，只是想要得到有序集合的成员数量，那么 `ZCARD` 是一个更简单、更高效的选择。



### ZREMRANGEBYSCORE

> `Z` - 有序集合
> `REM` - remove 移除
> `RANGE` - range 排序
> `BYSCORE` - by score 按照分数

Redis Zremrangebyscore 命令用于移除有序集中，指定分数（score）区间内的所有成员。

#### 语法
```shell
ZREMRANGEBYSCORE KEY_NAME min max
```

#### 返回值
被移除成员（member）的数量。

#### 示例
```shell
redis 127.0.0.1:6379> ZREMRANGEBYSCORE salary 1500 3500 # 移除所有薪水在 1500 到 3500 内的员工
(integer) 3

redis> ZRANGE salary 0 -1 WITHSCORES  # 剩下的有序集成员
1) "jack"
2) "5000"
```



### PEXPIRE

Redis PEXPIRE 命令和 EXPIRE 命令的作用类似，但是它**以毫秒为单位**设置 key 的生存时间，而不像 EXPIRE 命令那样，以秒为单位。

#### 语法

```shell
PEXPIRE KEY_NAME milliseconds
```

#### 返回值

设置成功，返回 1

key 不存在或设置失败，返回 0

#### 示例

```shell
redis 127.0.0.1:6379> SET mykey "Hello"
"OK"
redis 127.0.0.1:6379> PEXPIRE mykey 1500
(integer) 1
redis 127.0.0.1:6379> TTL mykey
(integer) 1
redis 127.0.0.1:6379> PTTL mykey
(integer) 1498
```



## 实现 - IP限流器

基于 Redis 和 滑动窗口 实现的 IP 限流器通常使用 Redis 的有序集合（sorted set）来存储每个 IP 地址的请求时间戳。每个请求都会被添加到这个集合中，并且**使用当前时间戳作为分数**。通过这种方式，可以快速地检查一个 IP 在给定的时间窗口内是否超过了请求次数的限制。



假设现在要限制每个IP地址每秒最多100个请求，则

-   限流对象：IP地址
-   时间窗口：1s
-   请求阈值：100
-   分数：当前时间戳
-   成员：当前时间戳



Lua 脚本用于原子性地执行Redis的以下操作：

-   移除时间窗口之外的旧请求。
-   检查当前时间窗口内的请求数量是否超过了阈值。
-   如果没有超过阈值，记录新的请求。

