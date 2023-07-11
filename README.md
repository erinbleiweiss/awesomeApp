Sample web app providing a basic REST API (Golang/Mux), DB (SQLite3), and frontend (Vue.JS)

# Demo Instructions
Web app can be viewed at http://localhost:8080/  
Button to access admin panel can be found at the top-right of the navigation  
An admin account is predefined, with credentials:
```
username: admin
password: admin
```

# Run
## Using Docker Compose
0. Prerequisites: Install Docker and Docker Compose
1. `docker-compose up`
2. View web app at http://localhost:8080/

## Manually
0. Prerequisites:
- go 1.20
- node 1.18

```
brew install go node
```

1. Run backend:
```
cd backend
go get ./...
go run main.go
```
2. Edit `frontend/view.config.js` (change `backend` to `localhost`):
```js
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: 'http://localhost:8888/',
  }
})
```
3. Run frontend:
```
cd frontend
npm install
npm run serve
```
4. View web app at http://localhost:8080/

