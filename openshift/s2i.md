when running
`oc new-app php:7.3 --name=my-app https://github.com/alfredxiao/myphpapp.git#mybranch --context-dir myapp`

it uses php:7.3 image as builder and create resources
- imagestream
- builderconfig
- deployment
- service

`oc logs -f buildconfig/my-app` to track its progress
`oc expose service/mya-app` to expose the servide to outside world
