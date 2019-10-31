# GO-SDK

go语言版本的中台服务调用的SDK

### 下载
* 使用git直接clone到本地的$GOPATH/src下

```bash
mkdir -p $GOPATH/src/gosdk
git clone git@github.com:tangshanshan1998/gosdk.git $GOPATH/src/gosdk

```

* 或使用go mod

```bash
# 在go.mod中添加一行
github.com/tangshanshan1998/gosdk latest

# 执行go mod vendor
go mod vendor
```

### 基本使用

```
// 获取对象，head是请求的HEAD字段，用来解析HEAD中的Authorization中的token
client, err:=gosdk.GetClientInstance(head)

// 对Authorization中的token解析，或对SetToken()中token解析，或SetAppInfo()
client, err = client.SetToken(token)
client, err = client.SetAppInfo(appid, appkey, channel, version)

// 可以使用SetServices()自定义服务地址，或通过serviceKey从环境变量中寻找服务地址（前者优先级高）
// services是map[string]string，key是serviceKey，value是服务地址
client = client.SetServices(services)

// 调用服务
// serviceKey对应服务地址；method是请求的方法，如post、get；api是具体请求的接口地址；params是要传递的参数，是map[string]interface{}的类型；
// alias是服务的别名；contentType是请求的格式，如application/x-www-form-urlencoded;file是上传文件时使用，一般为nil。
resp, err1 = client.Call(serviceKey, method, api, params, alias, contentType, file)

// resp是服务返回的结果，是[]byte数组，转化为string，优化内存
str := (*string)(unsafe.Pointer(&respBytes))

```
