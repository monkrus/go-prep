# go-prep
Interview preparation tasks

Steps:
Win10, go 1.14.4 
1. go mod init github.com/monkrus/go-prep.git

2. Create basic web server with net/http
- `log.Fatal` checks the connection status

3. Run `go run main.go` to test

4. Run `go get -u github.com/gorilla/mux` 



- Перепиши все на Gorilla/Mux (это просто самый популярный шаблонизатор, я с ним работал на уроке)

- Далее, создай параметр запроса "name", чтобы в зависимости от твоего параметра текст менялся, например
/hello?name=sergei
<b>Hello, Sergei</b>

- Потом создай ошибку, если заместо Нормального имени будет число
Например : /hello?name=123

- И проведи юнит-тестирование этого метода

- Естественно, тут мы переходим в QA - поэтому определи классы эквивалентности и уже позже только используй значения под тесты