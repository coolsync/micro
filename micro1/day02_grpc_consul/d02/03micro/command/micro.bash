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


?????????

micro????????????????????????

??????`micro[command]--help`???????????????????????????

?????????

micro[global options]??????[????????????][?????????]

?????????

??????3.0.0

?????????

auth????????????????????????????????????

?????????????????????micro call greeter Say.Hello'{name?????????John???}

cli???????????????cli

?????????????????????

env Get/set micro cli??????

gen????????????????????????????????????protobuf

init??????????????????????????????

???????????????micro kill[???]

????????????????????????

???????????????

????????????????????????????????????Micrologs helloworld

??????????????????????????????

??????????????????

?????????????????????micro run[???]

???????????????micro server

?????????????????????

?????????????????????????????????

???????????????

??????????????????????????????????????????????????????micro stats srv1 srv2 srv3

????????????????????????

?????????????????????????????????

???????????????????????????micro stream foo Bar.Baz'{key?????????value???}

?????????????????????micro update[???]

??????????????????????????????

?????????h??????????????????????????????????????????

???????????????

-c?????????????????????????????????~/.micro/config.json[$micro\u config\u file]

--env??????-e???????????????[$MICRO\u env]??????????????????

--?????????????????????micro server??????????????????????????????kubernetes[$micro\u profile]

--???????????????????????????micro??????[$micro\u????????????]???????????????????????????????????????

--auth\u address value???????????????auth????????????[$MICRO\u auth\u address]

--??????????????????????????????auth\u id?????????id[$MICRO\u auth\u id]

--??????????????????????????????auth\u?????????????????????[$MICRO\u auth\u secret]

--auth\u public\u key value JWT auth???base64??????PEM????????????????????????[$MICRO\u auth\u public\u key]

--auth\u private\u key value private key for JWT auth???base64??????PEM???[$MICRO\u auth\u private\u key]

--?????????\u?????????????????????????????????????????????[$MICRO\u registry\u address]

--?????????\u tls\u ca????????????????????????????????????????????????tls[$MICRO\u registry\u tls\u ca]

--?????????\u tls\u cert value Client cert for tls with registry[$MICRO\u registry\u tls\u cert]

--?????????\u tls\u????????????tls????????????????????????[$MICRO\u registry\u tls\u key]

--broker\u??????????????????????????????????????????[$MICRO\u broker\u??????]

--??????\u tls\u ca?????????????????????????????????????????????tls[$MICRO\u events\u tls\u ca]

--??????\u tls\u cert value Client cert for tls with events[$MICRO\u events\u tls\u cert]

--??????\u tls\u???????????????tls????????????????????????[$MICRO\u events\u tls\u key]

--broker\u tls\u ca????????????????????????????????????broker???tls[$MICRO\u broker\u tls\u ca]

--broker\u tls\u cert value Client cert for tls with broker[$MICRO\u broker\u tls\u cert]

--broker\u tls\u??????broker???tls????????????????????????[$MICRO\u broker\u tls\u key]

--????????????????????????????????????????????????[$MICRO\u store\u??????]

--proxy\u????????????????????????????????????HTTP??????[$MICRO\u proxy]

--?????????????????????????????????????????????????????????true???[$MICRO\u report\u??????]

--??????\u???????????????????????????[$micro\u service\u name]

--??????\u???????????????????????????[$micro\u service\u??????]

--??????\u????????????????????????[$MICRO\u service\u address]???????????????

--??????\u update??????????????????????????????????????????????????????????????????????????????(????????????true???[$MICRO\u?????????\u UPDATE]

--??????\u secret_ukey value key???????????????/???????????????????????????????????????????????????????????????????????????[$MICRO\u??????\u SECRET\u KEY]

--?????????-h show help???????????????false???

--?????????-v???????????????????????????false???