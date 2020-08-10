# task
Тестовое задание на позицию Junior Go Developer

Запуск проекта
```sh
make run
```
Параметры задаются в файле docker-compose.yml. 
* UPD_INTERVAL_FOOTBALL (указывается в секундах)
* UPD_INTERVAL_SOCCER (указывается в секундах)
* UPD_INTERVAL_BASEBALL (указывается в секундах)
* SERVER_HOST
* SERVER_PORT
* RPC_SERVER_HOST
* RPC_SERVER_PORT
* DB_USERNAME
* DB_PASSWORD
* DB_HOST
* DB_PORT
* LOG_LEVEL (logrus levels: Trace, Debug, Info, Warn, Error, Fatal, Panic)
* LINE_PROVIDER_API_URL: http://lineprovider:8000/api/v1/lines
