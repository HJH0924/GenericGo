--[[
    检查给定时间窗口内的请求数量是否超过了阈值
    输入：
    KEYS[1] - 限流对象
    ARGV[1] - 时间窗口大小（毫秒数）
    ARGV[2] - 阈值（例如：100个请求）
    ARGV[3] - 当前时间（Unix时间戳 毫秒数）
    输出：
    true  - 触发限流
    false - 没有触发限流
--]]

local key = KEYS[1]

local interval = tonumber(ARGV[1])
local threshold = tonumber(ARGV[2])
local now = tonumber(ARGV[3])

-- 计算时间窗口的起始时间
local min = now - interval

-- 移除时间窗口之外的旧请求
redis.call('ZREMRANGEBYSCORE', key, '-inf', min)

-- 计算当前时间窗口中的请求数量
local cnt = redis.call('ZCOUNT', key, '-inf', '+inf')

if cnt >= threshold then
    -- 请求数量超过阈值，触发限流
    return "true"
else
    -- 没有触发限流，将当前请求加入到有序集合中
    -- ZADD key score member
    redis.call('ZADD', key, now, now)
    -- 设置过期时间
    redis.call('PEXPIRE', key, interval)
    return "false"
end