# Todo list

* Implement tools to check test coverage.

* Refactor database features from API to PIM service  
    * `DONE` Implement local postgres container for test purposes. 
    * `IN PROGRESS` Make Migrations in PIM  
    * Make Seeds in PIM  
    * Move PIM to an internal folder for code security.      
    * Make Backup utilities in PIM CLI  
    * Implement Dependency Managment in PIM  
  
* Decouple all possible services from main API.
    * Refactor main API logic using the new services.  

* Implement startup features in Clirp to start SMAPI, PIM, SAS and others.  
    * Enable Smapi Startup.  
    * Enable PIM Startup.  
    * Enable SAS Startup.    
      
* Implement cobra for CLI  
    `DONE` Clirp.  
    `DONE` Sas.  