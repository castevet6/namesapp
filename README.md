# SOLITA DEV ACADEMY: Names app #
## Copyright (C) 2021 Antti Manninen ##

Names app is part of pre-task for 2021 Solita Dev Academy. This web app presents list of names and their amount in the company. Data is fetched directly from names.json data source. 

The names up is running in [Heroku instance](https://castevet6-namesapp.herokuapp.com)

### Tools used ###
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
├── ui-client
    └── src             := UI React source code
        ├── components  := React components
        ├── services    := Services to handle names data in front-end
        └── util        := Helper functions
```

### Setting up the app in Linux environment
1. You should have following installed:
    * __Go__ binaries (I used version 1.13)
    * __make__ binaries (not mandatory)
    * __npm__ package manager
2. Clone the repository: `git clone https://github.com/castevet6/namesapp.git`
3. install React dependencies: `cd ui-client && npm install'
4. build from app root folder: `make`
5. start the app: `make run`
    * alternatively: ./bin/namesapp
    * alternatively: ./go run ./cmd/namesapp
6. start the binaries with flag -devmode to enable CORS (I used this when running React locally)
    * also `make rundev` starts the app with -devmode set
