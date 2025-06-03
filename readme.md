# mini_service_citatnik

Микросервис для обработки цитат, написанный на Go.

## 📌 Основные функции
- Добавление новой цитаты
- Получение случайной цитаты
- Фильтрация по автору
- Получение всех цитат
- Удаление цитаты по ID

## 🚀 Технологии
```plaintext
Go 1.24.2+
mux
```

## ⚙️ Установка
### Через Go:
```bash
git clone https://github.com/Alexandrjob/mini_service_citatnik
cd mini_service_citatnik
go mod tidy
go run main.go
```

## 🫳🏻 Взаимодействие
### Curl команды:
#### POST запрос (одна строка)
```cmd
curl -X POST http://localhost:8080/quotes -H "Content-Type: application/json" -d "{\"author\":\"Confucius\", \"quote\":\"Life is simple, but we insist on making it complicated.\"}"
```

#### GET запросы
```

curl http://localhost:8080/quotes
curl http://localhost:8080/quotes/random
curl http://localhost:8080/quotes?author=Confucius
```

#### DELETE запрос
```
curl -X DELETE http://localhost:8080/quotes/1
```