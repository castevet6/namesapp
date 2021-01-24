# SOLITA DEV ACADEMY: Names app #
## Copyright (C) 2021 Antti Manninen ##

### Tools used
* __Go__ for Web API and data handling
* __React__ for UI structure and functionality
* __React Bootstrap__ for UI styling
* __make__ for building the app
* __Heroku__ for deployment

### Folder structure ###
```
├── bin                 := folder Go app binaries
├── cmd               
│   └── namesapp        := Go names app
├── data                := Directory containing the data source file `names.json`
├── pkg       
│   └── models          := Data model and service layer for names list
├── ui                  := Compiled static files from React app
    └── src             := UI React source code
        ├── components  := React components
        ├── services    := Services to handle names data in front-end
        └── util        := Helper functions
```
