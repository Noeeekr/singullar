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

IN ORDER: 
-> Decouple postgres SQL code from SMAPI structure.
    -> Implement local postgres SQL.
    -> Implement it in an internal folder.
    -> Make Postgres Dockerfile
    -> Make Migrations
    -> Make Backup utilities
    -> Make Seeds

-> Make SMAPI implementation functional
    -> Connect to PMI
    -> Make startable

-> Implement Clirp service statup features
    -> Make it capable of starting Smapi.
    -> Make it capable of starting Pmi.
    -> Make it capable of starting Sas
    
-> Implement cobra for CLI
    [Done] Clirp.
    [Done] Smapi.
    [Done] Sas.