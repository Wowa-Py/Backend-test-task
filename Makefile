.PHONY: deps build run clean default parse

# Команда для установки зависимостей
deps:
	go get -v ./...

# Команда для сборки проекта main1
build-main1:
	go build -o ./bin/main1 ./1/main.go

# Команда для запуска проекта main1
run-main1: build-main1
	./bin/main1

# Команда для очистки собранных файлов main1
clean-main1:
	rm -f ./bin/main1

# Команда для сборки проекта парсера
build-parser:
	go build -o ./bin/parser ./2/parser.go

# Команда для запуска проекта парсера
run-parser: build-parser
	./bin/parser

# Команда для очистки собранных файлов парсера
clean-parser:
	rm -f ./bin/parser

# Команда по умолчанию (сборка и запуск проекта main1)
default: build-main1 run-main1

# Команда для сборки и запуска парсера
parse: build-parser run-parser

# Команда для очистки собранных файлов main1 и парсера
clean: clean-main1 clean-parser
