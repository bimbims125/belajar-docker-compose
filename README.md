# Go web API with MySQL

### Clone Project
```
git clone https://github.com/bimbims125/belajar-docker-compose.git
```

### Run Project
```
go run main.go
```

## Using docker compose
You can use docker compose for run multicontainer

### Setup docker compose

#### Setup MySQL container
Change ```container_name``` with your preferences name

**Example :**
```
container_name=mysql-container
```

Change ```MYSQL_ROOT_PASSWORD``` value with your database password

**Example :**
```
MYSQL_ROOT_PASSWORD=your_db_password
```
Change ```MYSQL_DATABASE``` value with your database name

**Example**
```
MYSQL_DATABASE=your_db_name
```

#### Setup go application
Change go application ```ports```

**Example :**
```
ports="8080:8080"
```

Change password, hostname, database name ```db/setup.go``` file

```
dsn := "root:{your_db_password}@tcp({mysql_services_name}:3306)/{your_db_name}"
```
**Notes** : You can see ```mysql_services_name``` in compose file

### Run docker compose

You can run docker compose like this

```
docker compose up --build
```

## Enjoy!


