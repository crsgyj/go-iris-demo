# go-iris 实践 - 货品录入系统

I wrote nodejs before, now try go-iris.
我之前写nodejs, 现在尝试go-iris。

### 

\> before - what's you need: 

* iris  -  "github.com/kataras/iris"
* redis  -  "github.com/go-redis/redis"
* go-validator  -  "gopkg.in/go-playground/validator.v9"
* toml  -   "github.com/pelletier/go-toml"

### RUN

database:

 I use docker:

```bash
$ docker run -p 6379:6379 --name redis redis@latest
```

if you not ,  visit  https://redis.io/download to download redis is also a good choice.

server:

```go
$ go run main.go
```

client: 

```javascript
$ cd client && yarn && yarn serve
```



### Preview

the interface is like pic below:

![image](https://crsgyj.oss-cn-shanghai.aliyuncs.com/pictures/%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20190625024042.png?Expires=1561405603&OSSAccessKeyId=TMP.AgFMO66N34CNAFHBZM4wrhVY6yuyOEFdwvaP0bWtPmJX6Nu2i543HnWsVJb_AAAwLAIUIKnDrfagudCT2b73U1mx-O1ijncCFCn8OS6cE_We4uDyAFouTsqBFvVT&Signature=dpNaFgpJxbrti4gjAczGOmPlQa8%3D)

![img](https://crsgyj.oss-cn-shanghai.aliyuncs.com/pictures/%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20190625024141.png?Expires=1561405351&OSSAccessKeyId=TMP.AgFMO66N34CNAFHBZM4wrhVY6yuyOEFdwvaP0bWtPmJX6Nu2i543HnWsVJb_AAAwLAIUIKnDrfagudCT2b73U1mx-O1ijncCFCn8OS6cE_We4uDyAFouTsqBFvVT&Signature=ec9xHOpNiGEfB6TyKuuELRm%2BIWk%3D)

