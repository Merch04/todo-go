# Todo Application

![Go](https://img.shields.io/badge/Go-1.18+-blue.svg) ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13.0+-blue.svg)

(Todo) на языке Go. Проект разработан с использованием PostgreSQL в качестве бд.

## Установка

Перед началом работы убедитесь, что у вас установлены следующие зависимости:

- [Go](https://golang.org/dl/) (1.18 или выше)
- [PostgreSQL](https://www.postgresql.org/download/)

## Билд проекта

1. Клонируйте репозиторий:

   ```bash
   git clone https://github.com/Merch04/todo-go.git
   cd todo-go
   ```
2. Подтяните зависимости:
   ```bash
   go tidy
   ```
3. Сбилдите приложение:
    
    ```bash
    go build .\cmd\api\
    ```
4. Настройте конфигурацию:
   
   Приложение использует файл конфигурации для настройки параметров (yaml). Пример конфигурации:

    ```yaml
    port: 8000
    
    auth:
      hash_salt: "hash_salt"
      signing_key: "signing_key"
      token_ttl: 86400
    
    db:
      host: "localhost"
      port: 5432
      user: "postgres"
      password: "postgres"
      dbname: "todo"
      sslmode: disable
    ```
      ### Параметры:
      + **port**: Порт, на котором будет запущен сервер.
      + **auth**: Настройки аутентификации.
        - **hash_salt**: Соль для хеширования паролей.
        - **signing_key**: Ключ для подписи токенов.
        - **token_ttl**: Время жизни токена в секундах.
      + **db**: Настройки подключения к базе данных PostgreSQL.
        - **host**: Адрес хоста базы данных.
        - **port**: Порт базы данных.
        - **user**: Имя пользователя базы данных.
        - **password**: Пароль к базе данных.
        - **dbname**: Название базы данных.
        - **sslmode**: Режим SSL-соединения.
        - 
      ### Размещение
      > файл с конфигом хранить в ./config/
 
  5. Запустите:
  
     ```bash 
      .\api.exe
     ```
