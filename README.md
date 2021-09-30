This repo contains code which was written while learning Golang from the materials on the https://metanit.com and couple
from https://tour.golang.org.

Some exercises will need database (Docker used instead) and driver for it. Also used HTTP router and URL matcher
[mux](https://github.com/gorilla/mux) from [Gorilla web toolkit](https://www.gorillatoolkit.org/). To get all this use  
``go get github.com/go-sql-driver/mysql``  
``go get github.com/gorilla/mux``

For convenience of using, along with creating and running container also created table ``products`` with a couple of
records. SQL-script for creating table and initialization data stored in
[.docker](.docker/setup.sql) folder.

For creating and running MySQL Docker-container use:  
``docker-compose up --force-recreate -d``

Of course, it requires to be installed Docker and Docker Compose.
