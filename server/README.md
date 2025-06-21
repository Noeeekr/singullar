# Singullar Backend Documentation

## About 

Singullar server is a "Modular Monolith" API. All services reside within a single Go module. This is viable since all services are more like logically related binaries.

### Singullar SAS

Singular **S**tatic **A**ssets **S**erver (SAS) is a small but powerfull assets server powered with GIN framework, focused on providing client static files.

### Singullar SMAPI

Singular **S**ervice **M**anagment **API** (SMAPI) provides many Enterprise Resource Planning features such as authorization logic, parallel product managment, information access and more.

### Singullar Clirp

Singular **CLI** for **R**everse **P**roxy (CLIRP) is a firewall sitting in front of the API. Additionally, it can start other services.

### Singullar PMI

Currently using vercel Postgres SQL database. Soon will be a local Singular **P**ostgreSQL **M**anagment **I**nterface (PMI)

# Todo list

[Done] Create a reverse proxy.
    -> Make it capable of starting Api.
        -> Make api startable.
    -> Make it capable of starting Static File Server.
        [Done] Make static file server startable.

-> Decouple postgres SQL code from api folder structure.
    -> Implement local postgres SQL.
    -> Implemenet it in an internal folder.
    
-> Conteinerize the application.
-> Implement cobra for CLI
    [Done] Clirp.
    [Done] Smapi.
    [Done] Sas.