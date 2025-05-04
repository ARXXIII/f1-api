.PHONY: run build tidy

# Название бинарника
APP_NAME=f1-api

# Путь до главного файла
MAIN=./cmd/main.go

# Команда для запуска сервера
run:
	go run $(MAIN)

# Команда для сборки бинарника
build:
	go build -o $(APP_NAME) $(MAIN)

# Команда для обновления зависимостей
tidy:
	go mod tidy
