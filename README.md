# go-prep
Interview preparation tasks
- Win10, go 1.14.4 

Steps:
1. go mod init github.com/monkrus/go-prep.git

2. Create basic web server with net/http
- `log.Fatal` checks the connection status

3. Run `go run main.go` to test

4. Run `go get -u github.com/gorilla/mux` 

5. Replace `net/http` with `Gorilla/Mux` 

6. Create a query parameter `name` 

7. Check open/closed ports

8. Install `go get github.com/anvie/port-scanner`

8. Create an intentional error

9. Write unit test for this method

- Потом создай ошибку, если заместо Нормального имени будет число
Например : /hello?name=123

- И проведи юнит-тестирование этого метода

- Естественно, тут мы переходим в QA - поэтому определи классы эквивалентности и уже позже только используй значения под тесты
