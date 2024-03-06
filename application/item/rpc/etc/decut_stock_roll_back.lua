local lockKey = KEYS[1]
local lockValue = ARGV[1]

local maxRetries = 10
local retryInterval = 1

local retries = 0
local acquired = 0
--检查锁是否存在
repeat
    acquired = redis.call("SETNX", lockKey, lockValue)
    if acquired == 1 then
        redis.call("EXPIRE", lockKey, 10)  -- 设置锁的过期时间
        local stock = redis.call("HGET",KEYS[2],KEYS[3])
        redis.call("HSET",KEYS[2],KEYS[3],stock + ARGV[2])
        redis.call("EXPIRE",KEYS[2],ARGV[3])
        redis.call("DEL", lockKey)
    else
        -- 未获取到锁，等待重试
        retries = retries + 1
        if retries > maxRetries then
            -- 达到最大重试次数，处理竞争情况
            print("达到最大重试次数,stock - ",ARGV[2])
        end
    end
until acquired == 1 or retries > maxRetries
