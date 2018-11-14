# sample_go_webapp

This web app is a very simple application that will connect to a mysql DB and print out some DB information.

The DB connection string can set in either the config file, or as a system variable.
* System Variable
  * export Connection_String="id:password@tcp(your-amazonaws-uri.com:3306)/dbname"
* config.json
  * {"Connection_String": "id:password@tcp(your-amazonaws-uri.com:3306)/dbname"}
