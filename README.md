### ABOUT

The project is api allow user follow posts of blogs posts. For do that I use RSS feed url where I have a scraper application to extract the new posts and save in database. 

### MOTIVATION 

The project has main goal improve my knowledge about golang and learn new things.

### TECHNOLOGIES

- Golang
- Postgres(Database)
- Docker
- Docker compose
- Github actions(pipeline CD)
- Kubernetes(to run all containers in production)

### INSTRUCTIONS TO RUN PROJECT LOCALLY

- Clone
- Create **.env** file based **.env.example** file
- Execute command **docker-compose up -d** to execute the follow containers: api, scraper, postgres and pgadmin(postgres client).
- Execute command **cd migrations && goose postgres "db_url_connection" up** to execute migrations
  
### EXTRA 

- Directory named **infra** where has all kubernetes script to deploy application in production.
- To test the api I created **endpoints.http** file where have all endpoints of api. Tip: you need to install **REST CLIENT** plugin in Vscode.