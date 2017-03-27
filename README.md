# goconfiger
golang实现简单的ini配置文件解析

- 关于goconfiger
  - goconfiger是一个使用golang开发的配置加载简易packet，旨在使用golang程序开发时加载类似ini类型的文件。支持多个配置文件和单一配置文件两种模式下的使用。详见示例。
- 如何使用
  - 在golang中使用import 导入goconfiger。
  - 在你的golang工程中使用goconfiger暴露的function进行配置文件相关的操作。
- 示例：
```
# example.ini
        
[listener]
server_addr=0.0.0.0
server_port=10929
        
[update]
update_addr = 127.0.0.1
update_port = 10931
        
[task]
max_delivery_time = 10
       
[file]
node_info_file = ipinfo.xml
protected_node_list = protected.list
cache_dir = /cache1/
cache_dir = /cache2/
cache_dir = /cache3/
```



    
```
// 多文件加载， 每个config file加载完后都会形成一个实例对象，各自管理
import (
	"goconfiger"
	"fmt"
)

var configFile string = "/Users/src/example/example.ini"

func main() {

	cfg, err := goconfiger.LoadAndGetConfiger(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}

fmt.Println(cfg.GetValueByString("listener", "server_addr"))
fmt.Println(cfg.GetValueByInt("listener", "server_port"))
fmt.Println(cfg.GetValueByString("update", "update_addr"))
fmt.Println(cfg.GetValueByFloat64("update", "update_port"))
fmt.Println(cfg.GetValueByString("task", "max_delivery_time"))
fmt.Println(cfg.GetValueByString("file","protected_node_list"))
fmt.Println(cfg.GetValueByInt("file", "node_info_file"))
fmt.Println(cfg.GetValueByStringList("file", "cache_dir"))
 ```

```
// 单配置加载，goconfiger内部管理实例对象
import (
	"goconfiger"
	"fmt"
)

var configFile string = "/Users/src/example/example.ini"

func main() {

	if err := goconfiger.LoadConfig(configFile);err != nil{
		fmt.Println(err)
		return
	}

fmt.Println(goconfiger.GetValueByString("listener", "server_addr"))
fmt.Println(goconfiger.GetValueByInt("listener", "server_port"))
fmt.Println(goconfiger.GetValueByString("update", "update_addr"))
fmt.Println(goconfiger.GetValueByFloat64("update", "update_port"))
fmt.Println(goconfiger.GetValueByString("task", "max_delivery_time"))
fmt.Println(goconfiger.GetValueByString("file","protected_node_list"))
fmt.Println(goconfiger.GetValueByInt("file", "node_info_file"))
fmt.Println(goconfiger.GetValueByStringList("file", "cache_dir"))
 ```


