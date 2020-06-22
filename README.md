# go-todo

A responsive web app written in Vue.js with a Golang api backend. Todo items are stored in a mysql database. Frontend communicates with the backend via the exposed API functions.

<p align="center">
  <img src="https://github.com/hbostann/go-todo/blob/master/gif/demo.gif" alt="DEMO">
</p>

## Installation and Running

1. Clone the repo
2. Run backend with `go run gotodo.go`. Make sure that you have a mysql server running (program uses port 4211 for mysql connection)
3. Serve the Vue.js frontend. For simplicity you can use `npm run serve` in frontend folder (Don't do this if you are actually publising this webapp).

## License
[MIT](https://choosealicense.com/licenses/mit/)