# CRUD_Go
Приложение с RestApi CRUD

## Методы взаимодействия:
```bash
GET     "/"      - Стартовая страница
GET     "/user"  - Получения данных пользователя
GET     "/users" - Получение всех пользователей
POST    "/user"  - Добавление пользователя
PUT     "/user"  - Обновление пользователя
DELETE  "/user"  - Удаление пользователя

```
## Все запросы на /user присылаются в таком формате Json
```json
{
    "Name": "",
    "LastName": "",
    "SurName": "",
    "Gender": "",
    "Status": ""
}
```
## Dockerfile Команды
```bash
docker build -t crud_go .
docker run -p 8090:8090 --name crud_go crud_go
```
