# Singullar Backend Documentation

## About  

Singullar is a Enterprise Resouce Planning focused on providing the best tools for schools. Singullar backend is composed of many services divided in core logic and utility. They act as a "Modular Monolith" interconnected to the main API. All services reside within a single Go module, this is viable since all services are all logically related.

### Singullar Core Services

#### Singullar SAS

Singular **S**tatic **A**ssets **S**erver (SAS) is a small but powerfull assets server powered with GIN framework, focused on providing client static files.

#### Singullar SMAPI

Singular **S**ervice **M**anagment **API** (SMAPI) provides many Enterprise Resource Planning features such as authorization logic, parallel product managment, information access and more.

#### Singullar Clirp

Singular **CLI** for **R**everse **P**roxy (CLIRP) is a firewall sitting in front of the API. Additionally, it can start other services.

#### Singullar PIM

Currently using vercel Postgres SQL database. Soon will be a local Singular **P**ostgreSQL **I**nterface **M**anagment (PIM)