# Todo list

* Implement tools to check test coverage.

* Database
    * `DONE` Cobra command line interface.
    * `DONE` `UPGRADABLE` Implement local containers for testing. 
    * `DONE` `UPGRADABLE` Implement migrations  
    * `DONE` `UPGRADABLE` Implement dependency managment  
    * `DONE` Move to internal folder for security.
    * `DONE` Command line interface migrations command.
    * Implement --ignore-existing --recreate-exiting for database migrate commands
    * Implement seeding. 
    * Update database/scan to have => .IgnoreScanErrors() .ScanErrorOn() => With accepted errors: Not found, Found 
    * Implement command line backup utilities. (from a different repository)  

* API
    * `DONE` Cobra command line interface.
    * `DONE` Decouple logistic services from API.
    * `DONE` Refactor API to implement the decoupled services.  