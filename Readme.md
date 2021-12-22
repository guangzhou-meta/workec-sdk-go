# workec-sdk-go

EC公共SDK，Golang版本

## 使用方式

### 1. 安装

```shell
go get -u github.com/guangzhou-meta/workec-sdk-go
```

### 2. 配置Api

```go
package main

import (
	"github.com/guangzhou-meta/workec-sdk-go/api"
	"github.com/guangzhou-meta/workec-sdk-go/common"
)

const (
	appId     = "123"
	appSecret = "abc"
	corpId    = "45678"
)

var apiConfig = api.NewConfig(appId, appSecret, corpId)

func main() {
	apiConfig.
		SetAppId(appId). // 设置EC服务AppId
		SetAppSecret(appSecret). // 设置EC服务AppSecret
		SetCorpId(corpId). // 设置EC服务企业Id
		SetServerBaseUrl(common.ServerBaseUrlV2) // 设置EC服务基础地址
}
```

### 3. 调用Api
```go
package main

import (
	"fmt"
)

import (
	"github.com/guangzhou-meta/workec-sdk-go/api"
)

const (
	appId     = "123"
	appSecret = "abc"
	corpId    = "45678"
)

var apiConfig = api.NewConfig(appId, appSecret, corpId)

func main() {
	res, err := apiConfig.OrgStructInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
```

## Api风格
例子:

| Api路由           | Api名称          |
|:----------------|:---------------|
| org/struct/info | OrgStructInfo  |
| org/dept/create | OrgDeptCreate  |
| role/changeName | RoleChangeName |
