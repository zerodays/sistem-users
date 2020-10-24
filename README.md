# Sistem Users
Users microservice for sistem.

## Running
1. Start postgresql database
    ```
   docker run -p 5432:5432 -e POSTGRES_PASSWORD=postgres -v sistem_users:/var/lib/postgresql/data -d postgres
   ```
2. Build the executable 
    ```
   make build
   ```
3. Start the server
    ```
   ./users serve
   ```

For other options see help with
```
./users help
```

## Developing
For development postgresql database should be started as before
```
   docker run -p 5432:5432 -e POSTGRES_PASSWORD=postgres -v sistem_users:/var/lib/postgresql/data -d postgres
```

Instead of building the executable by hand every time you change the code, you can
run the server with hot reload using [air](https://github.com/cosmtrek/air).
After installing [air](https://github.com/cosmtrek/air) run `air` to do that.

When changing default configuration or database migrations in `configs` directory, you should embed the data
in binary executable using the tool [go-bindata](https://github.com/go-bindata/go-bindata). To do that run
```
make bindata
```
