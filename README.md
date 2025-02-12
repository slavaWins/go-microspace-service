<div align="center">

<h1> go-microspace-service </h1>
 <b>gin</b> | <b>swaggerUi</b> | <b>jwt auth</b> | <b>Docker</b>
</div>
 
## About

Базовый микросервис. С готовой jwt авторизацией и мидлвейром для неё вынесенным в отдельную библиотеку


## Cmd 

Запуск  сервиса 

    go run main.go 


Генерация Свагера


    swag init --parseDependency --parseInternal --parseDepth 2




Свагер будет досутпен по юрл

http://localhost:8081/swagger/index.html
