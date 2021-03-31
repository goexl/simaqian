# wukong（行者，悟空）

悟空，Golang版本的缓存框架，存在的唯一目的就是加速系统性能

# 为什么要叫悟空

《西游记》里记载，孙悟空一个筋斗云十万八千里，和系统缓存加载异曲同工

## 功能

- 支持本地内存缓存BigCache
- 支持Redis
- 支持Memcache
- 统一的缓存接口
- 方便集成自己的缓存
- 更友好的API设计

## 内置存储

- Redis
- Memcache
- BigCache

## 内置序列化

- Msgpack
- Json
- Xml

## 为什么要写这个缓存

最近一直在寻找一个统一的Golang版本的缓存框架，无奈于Golang的生态确实不如Java，各自为政，许久寻得一个框架基本满足要求https://github.com/eko/gocache
但是其接口设计使用使用起来特别麻烦，每次存放数据，都得传Options参数，还不能省略，十分讨厌，所以在其基础之上，增加更易使用的方法形成此框架

## 样例代码

### 简单缓存

简单缓存只提供了缓存的基本功能，不包括链式调用、监控等功能

#### 使用Redis

```go
store := wukong.NewMemcache(redis.NewClient(&redis.Options{
	Addr: "127.0.0.1:6379",
}))

cache := wukong.New(store)
err := cache.Set("my-key", "my-value", WithExpiration(15*time.Second))
if err != nil {
    panic(err)
}

value := cache.Get("my-key")
```

#### 使用Memcache

```go
store := wukong.NewRedis(memcache.New(
	"10.0.0.1:11211", 
	"10.0.0.2:11211", 
	"10.0.0.3:11212",
))

cache := wukong.New(store)
err := cache.Set("my-key", "my-value", WithExpiration(15*time.Second))
if err != nil {
    panic(err)
}

value := cache.Get("my-key")
```

#### 使用BigCache

```go
bigcache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(5*time.Minute))
store := store.NewBigcache(bigcache)

cache := wukong.New(store)
err := cache.Set("my-key", "my-value", WithExpiration(15*time.Second))
if err != nil {
    panic(err)
}

value := cache.Get("my-key")
```

### 定制缓存

#### 序列化器

序列化器可以定制你存放在缓存中的数据序列化方式（从结构体到二进制数据以及从二进制数据到结构体），wukong提供了方便使用的序列化器，且内置常用的序列化器

##### 代码

```go
store := wukong.NewRedis(redis.NewClient(&redis.Options{
	Addr: "127.0.0.1:6379",
}), WithSerializer(wukong.SerializerJson{}))

cache := wukong.New(store)
err := cache.Set("my-key", "my-value", WithExpiration()15*time.Second)
if err != nil {
    panic(err)
}

value := cache.Get("my-key")
```

##### 支持的序列化器有

- wukong.SerializerJson
- wukong.SerializerXml
- wukong.SerializerMsgpack

##### 增加自己的序列化器

实现序列化接口，就可以方便的实现自己的序列化器

```go
type Serializer interface {
	// Marshal 将结构体编码成二进制数组
	Marshal(obj interface{}) ([]byte, error)
	// Unmarshal 将二进制数据解码成结构体
	Unmarshal(data []byte, obj interface{}) error
}
```

#### 缓存数据存放

可以很方便的增加自己的缓存存储方案，只需要实现一个特定的接口即可

##### 代码

```go
type Store interface {
	// Get 取得缓存值
	Get(key string) ([]byte, error)
	// Set 设置缓存值
	Set(key string, data []byte, options ...option) error
	// Delete 删除缓存值
	Delete(key string) error
	// Invalidate 设置缓存失效
	Invalidate(options ...invalidateOption) error
	// Clear 清除缓存
	Clear() error
	// Type 缓存类型
	Type() string
}
```

##### 内置的数据存放方案

- wukong.NewRedis
- wukong.NewBigCache
