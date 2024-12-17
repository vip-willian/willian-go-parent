// 声明模块自身的路径
module willian-go-parent

// 声明这个模块所用的Go版本
go 1.23

// 声明这个模块依赖的模块
require github.com/vmihailenco/msgpack v4.0.4+incompatible

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/bytedance/sonic v1.12.6 // indirect
	github.com/bytedance/sonic/loader v0.2.1 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.7 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.10.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.23.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/goccy/go-json v0.10.4 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmoiron/sqlx v1.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.12.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/protobuf v1.36.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/mysql v1.5.7 // indirect
	gorm.io/gorm v1.25.12 // indirect
)

// exclude 让go mod不要加载这个模块
// replace 声明替换，=> 前是代码中的路径，后是要替换成的路径
// retract 声明这个模块中，损坏的版本，或者要撤回，要大家别用的版本。作用是，设置了的版本，使用者碰到这个 版本的时候，就会看到对应的提示

// go mod init：初始化go mod， 生成go.mod文件，后可接参数指定 module 名，上面已经演示过。
// go mod download：手动触发下载依赖包到本地cache（默认为$GOPATH/pkg/mod目录）
// go mod graph：打印项目的模块依赖结构
// go mod tidy ：添加缺少的包，且删除无用的包
// go mod verify ：校验模块是否被篡改过
// go mod why：查看为什么需要依赖
// go mod vendor ：导出项目所有依赖到vendor下
// go build 构建二进制
// go get 添加依赖
// go get -u 更新依赖
