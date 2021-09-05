## Setup MySQL Docker
- fetch n run mysql image
```bash
docker run \
--name mysql-instance \
--mount type=bind,src=/Users/mjmhtjain/repo/vaccine-alert/scripts/db/schema.sql,dst=/docker-entrypoint-initdb.d/1.sql \
-e MYSQL_ROOT_PASSWORD=root \
-d mysql/mysql-server:latest
```

- run PhpMyAdmin Client
```bash
docker run --name myadmin -d --link mysql-instance:db -p 8080:80 phpmyadmin
```

- run web client on this url `http://localhost:8080/`

### To create a user in MySQL
- Enter CommandLine Client
```bash
docker exec -it mysql-instance mysql -u root -p
```

- Create User without password
```bash
CREATE USER 'root'@'%';
```

- Create User with password
```bash
CREATE USER 'root'@'%' IDENTIFIED BY 'password';
```

- Grant user privileges
```bash
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
```

docker run \
--name mysql-instance \
--mount type=bind,src=/Users/mjmhtjain/repo/vaccine-alert/scripts/db/schema.sql,dst=/docker-entrypoint-initdb.d/1.sql \
-e MYSQL_ROOT_PASSWORD=root \
-d mysql/mysql-server:latest