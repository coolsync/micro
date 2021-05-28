NAME:
   micro - A framework for cloud native development

   Use `micro [command] --help` to see command specific help.

USAGE:
   micro [global options] command [command options] [arguments...]

VERSION:
   v3.0.0

COMMANDS:
   auth      Manage authentication, accounts and rules
   call      Call a service e.g micro call greeter Say.Hello '{"name": "John"}'
   cli       Run the interactive CLI
   config    Manage configuration values
   env       Get/set micro cli environment
   gen       Generate a micro related dependencies e.g protobuf
   init      Generate a profile for micro plugins
   kill      Kill a service: micro kill [source]
   login     Interactive login flow.
   logout    Logout.
   logs      Get logs for a service e.g. micro logs helloworld
   network   Manage the micro service network
   new       Create a service template
   run       Run a service: micro run [source]
   server    Run the micro server
   service   Run a micro service
   services  List services in the registry
   signup    Signup to the Micro Platform
   stats     Query the stats of specified service(s), e.g micro stats srv1 srv2 srv3
   status    Get the status of services
   store     Commands for accessing the store
   stream    Create a service stream e.g. micro stream foo Bar.Baz '{"key": "value"}'
   update    Update a service: micro update [source]
   user      Print the current logged in user
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -c value                   Set the config file: Defaults to ~/.micro/config.json [$MICRO_CONFIG_FILE]
   --env value, -e value      Set the environment to operate in [$MICRO_ENV]
   --profile value            Set the micro server profile: e.g. local or kubernetes [$MICRO_PROFILE]
   --namespace value          Namespace the service is operating in (default: "micro") [$MICRO_NAMESPACE]
   --auth_address value       Comma-separated list of auth addresses [$MICRO_AUTH_ADDRESS]
   --auth_id value            Account ID used for client authentication [$MICRO_AUTH_ID]
   --auth_secret value        Account secret used for client authentication [$MICRO_AUTH_SECRET]
   --auth_public_key value    Public key for JWT auth (base64 encoded PEM) [$MICRO_AUTH_PUBLIC_KEY]
   --auth_private_key value   Private key for JWT auth (base64 encoded PEM) [$MICRO_AUTH_PRIVATE_KEY]
   --registry_address value   Comma-separated list of registry addresses [$MICRO_REGISTRY_ADDRESS]
   --registry_tls_ca value    Certificate authority for TLS with registry [$MICRO_REGISTRY_TLS_CA]
   --registry_tls_cert value  Client cert for TLS with registry [$MICRO_REGISTRY_TLS_CERT]
   --registry_tls_key value   Client key for TLS with registry [$MICRO_REGISTRY_TLS_KEY]
   --broker_address value     Comma-separated list of broker addresses [$MICRO_BROKER_ADDRESS]
   --events_tls_ca value      Certificate authority for TLS with events [$MICRO_EVENTS_TLS_CA]
   --events_tls_cert value    Client cert for TLS with events [$MICRO_EVENTS_TLS_CERT]
   --events_tls_key value     Client key for TLS with events [$MICRO_EVENTS_TLS_KEY]
   --broker_tls_ca value      Certificate authority for TLS with broker [$MICRO_BROKER_TLS_CA]
   --broker_tls_cert value    Client cert for TLS with broker [$MICRO_BROKER_TLS_CERT]
   --broker_tls_key value     Client key for TLS with broker [$MICRO_BROKER_TLS_KEY]
   --store_address value      Comma-separated list of store addresses [$MICRO_STORE_ADDRESS]
   --proxy_address value      Proxy requests via the HTTP address specified [$MICRO_PROXY]
   --report_usage             Report usage statistics (default: true) [$MICRO_REPORT_USAGE]
   --service_name value       Name of the micro service [$MICRO_SERVICE_NAME]
   --service_version value    Version of the micro service [$MICRO_SERVICE_VERSION]
   --service_address value    Address to run the service on [$MICRO_SERVICE_ADDRESS]
   --prompt_update            Provide an update prompt when a new binary is available. Enabled for release binaries only. (default: true) [$MICRO_PROMPT_UPDATE]
   --config_secret_key value  Key to use when encoding/decoding secret config values. Will be generated and saved to file if not provided. [$MICRO_CONFIG_SECRET_KEY]
   --help, -h                 show help (default: false)
   --version, -v              print the version (default: false)


姓名：

micro—云本地开发框架

使用`micro[command]--help`查看命令特定帮助。

用法：

micro[global options]命令[命令选项][参数…]

版本：

版本3.0.0

命令：

auth管理身份验证、帐户和规则

呼叫服务，例如micro call greeter Say.Hello'{name”：“John”}

cli运行交互式cli

配置管理配置值

env Get/set micro cli环境

gen生成微相关依赖关系，例如protobuf

init为微插件生成配置文件

杀死服务：micro kill[源]

登录交互登录流。

注销注销。

日志获取服务的日志，例如Micrologs helloworld

微业务网络的网络管理

新建服务模板

运行运行服务：micro run[源]

服务器运行micro server

服务运行微服务

注册表中的服务列表服务

注册微平台

统计信息查询指定服务的统计信息，例如micro stats srv1 srv2 srv3

状态获取服务状态

存储用于访问存储的命令

流创建服务流，例如micro stream foo Bar.Baz'{key”：“value”}

更新更新服务：micro update[源]

用户打印当前登录用户

帮助，h显示一个命令的命令列表或帮助

全局选项：

-c值设置配置文件：默认为~/.micro/config.json[$micro\u config\u file]

--env值，-e值设置要在[$MICRO\u env]中操作的环境

--配置文件值设置micro server配置文件：例如本地或kubernetes[$micro\u profile]

--服务正在（默认为“micro”）[$micro\u命名空间]中操作的命名空间值命名空间

--auth\u address value逗号分隔的auth地址列表[$MICRO\u auth\u address]

--用于客户端身份验证的auth\u id值帐户id[$MICRO\u auth\u id]

--用于客户端身份验证的auth\u秘密值帐户机密[$MICRO\u auth\u secret]

--auth\u public\u key value JWT auth（base64编码PEM）的密钥值公钥：[$MICRO\u auth\u public\u key]

--auth\u private\u key value private key for JWT auth（base64编码PEM）[$MICRO\u auth\u private\u key]

--注册表\u地址值逗号分隔的注册表地址列表[$MICRO\u registry\u address]

--注册表\u tls\u ca值证书颁发机构，用于具有注册表的tls[$MICRO\u registry\u tls\u ca]

--注册表\u tls\u cert value Client cert for tls with registry[$MICRO\u registry\u tls\u cert]

--注册表\u tls\u注册表为tls的键值客户端密钥[$MICRO\u registry\u tls\u key]

--broker\u地址值逗号分隔的代理地址列表[$MICRO\u broker\u地址]

--事件\u tls\u ca值证书颁发机构，用于具有事件的tls[$MICRO\u events\u tls\u ca]

--事件\u tls\u cert value Client cert for tls with events[$MICRO\u events\u tls\u cert]

--事件\u tls\u具有事件的tls的键值客户端密钥[$MICRO\u events\u tls\u key]

--broker\u tls\u ca值证书颁发机构，用于带有broker的tls[$MICRO\u broker\u tls\u ca]

--broker\u tls\u cert value Client cert for tls with broker[$MICRO\u broker\u tls\u cert]

--broker\u tls\u带有broker的tls的密钥客户端密钥[$MICRO\u broker\u tls\u key]

--存储地址值逗号分隔的存储地址列表[$MICRO\u store\u地址]

--proxy\u地址值代理请求通过指定的HTTP地址[$MICRO\u proxy]

--报表使用情况报告使用统计信息（默认值：true）[$MICRO\u report\u用法]

--服务\u微服务的名称值名称[$micro\u service\u name]

--服务\u微服务的版本值版本[$micro\u service\u版本]

--服务\u地址值地址，以在[$MICRO\u service\u address]上运行服务

--提示\u update在新二进制可用时提供更新提示。仅为发布二进制文件启用(默认值：true）[$MICRO\u提示符\u UPDATE]

--配置\u secret_ukey value key，用于编码/解码机密配置值。如果没有提供，将生成并保存到文件中[$MICRO\u配置\u SECRET\u KEY]

--帮助，-h show help（默认值：false）

--版本，-v打印版本（默认值：false）