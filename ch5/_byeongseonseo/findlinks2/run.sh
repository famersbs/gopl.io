echo "================================"
echo "https://golang.org"
go run main.go https://golang.org 

echo "================================"
echo "https://golang.org/404"
go run main.go https://golang.org/404 

echo "================================"
echo "https://ssl.google-analytics.com/ga.js"       # javascript 코드도 html 파싱에서 오류가 나지 않음
go run main.go https://ssl.google-analytics.com/ga.js

echo "================================"
echo "https://no_such_host"
go run main.go https://no_such_host