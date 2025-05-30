package RLock

// 获取读锁 Lua 脚本
// KEYS[1]: 读锁键
// KEYS[2]: 写锁键
// ARGV[1]: 客户端 ID
// ARGV[2]: 超时时间（毫秒）
const acquireReadLockScript = `
if redis.call('exists', KEYS[2]) == 1 then
    return 0
end
local count = redis.call('hincrby', KEYS[1], ARGV[1], 1)
redis.call('pexpire', KEYS[1], ARGV[2])
return count
`

// 释放读锁 Lua 脚本
// KEYS[1]: 读锁键
// ARGV[1]: 客户端 ID
const releaseReadLockScript = `
local count = redis.call('hget', KEYS[1], ARGV[1])
if count == nil then
    return 0
end
count = tonumber(count) - 1
if count > 0 then
    redis.call('hset', KEYS[1], ARGV[1], count)
    return count
else
    redis.call('hdel', KEYS[1], ARGV[1])
    if redis.call('hlen', KEYS[1]) == 0 then
        redis.call('del', KEYS[1])
    end
    return 0
end
`

// 获取写锁 Lua 脚本
// KEYS[1]: 写锁键
// ARGV[1]: 客户端 ID
// ARGV[2]: 超时时间（毫秒）
const acquireWriteLockScript = `
if redis.call('exists', KEYS[1]) == 1 then
    return 0
end
redis.call('set', KEYS[1], ARGV[1], 'PX', ARGV[2])
return 1
`

// 释放写锁 Lua 脚本
// KEYS[1]: 写锁键
// ARGV[1]: 客户端 ID
const releaseWriteLockScript = `
if redis.call('get', KEYS[1]) == ARGV[1] then
    return redis.call('del', KEYS[1])
else
    return 0
end
`

// 获取读锁 TTL 的 Lua 脚本
const getReadLockTTLScript = `
-- 检查读锁是否存在且由当前客户端持有
if redis.call('HEXISTS', KEYS[1], ARGV[1]) == 0 then
	return -2
end

-- 返回读锁剩余时间
return redis.call('PTTL', KEYS[1])
`

// 续期读锁的 Lua 脚本
const renewReadLockScript = `
-- 检查读锁是否存在且由当前客户端持有
if redis.call('HEXISTS', KEYS[1], ARGV[1]) == 0 then
	return 0
end

-- 续期读锁
redis.call('PEXPIRE', KEYS[1], ARGV[2])
return 1
`

// 续期写锁的 Lua 脚本
const renewWriteLockScript = `
-- 检查写锁是否存在且由当前客户端持有
if redis.call('GET', KEYS[1]) ~= ARGV[1] then
	return 0
end

-- 续期写锁
redis.call('PEXPIRE', KEYS[1], ARGV[2])
return 1
`

// 获取写锁 TTL 的 Lua 脚本
const getWriteLockTTLScript = `
-- 检查写锁是否存在且由当前客户端持有
if redis.call('GET', KEYS[1]) ~= ARGV[1] then
	return -2
end

-- 返回写锁剩余时间
return redis.call('PTTL', KEYS[1])
`
