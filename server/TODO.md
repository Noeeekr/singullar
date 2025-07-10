# Todo list (In order)

* Implement tools to check test coverage.

* Database
    * `DONE` Cobra command line interface.
    * `DONE` `UPGRADABLE` Implement local containers for testing. 
    * `DONE` `UPGRADABLE` Implement migrations  
    * `DONE` `UPGRADABLE` Implement dependency managment  
    * `DONE` Move to internal folder for security.
    * `DONE` Command line interface migrations command.
    * Implement schema configuration for user, database and table migrations
    * `DONE` Implement --ignore-existing --recreate-exiting for database migrate commands
    * `PARTIAL` Implement seeding. 
    * `DONE` Implement database/scan to have a switch on errors RowFound and RowNotFound.
    * Implement command line backup utilities. (in a different project a import and use here).  
    * Change operations and migrations to never commit by default so .Commit() must be called.
    
* API
    * `DONE` Cobra command line interface.
    * `DONE` Decouple logistic services from API.
    * `DONE` Refactor API to implement the decoupled services.  