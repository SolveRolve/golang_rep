# для генерации файлов из .yaml

cd src
go mod init src
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
mkdir api
cd api


# из старой версии свагера надо перевести в новую на сайте https://editor.swagger.io/
# и положить в файл api.yaml

# создать файл для конфигурации примеры на сайте
# https://github.com/oapi-codegen/oapi-codegen?tab=readme-ov-file#impl-gorillamux

echo "package: api
       generate:
         std-http-server: true
         models: true
       output: gen.go" > cfg.yaml

# и выполнить
go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml api_new.yaml

