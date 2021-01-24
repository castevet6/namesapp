package jsonreader

import (
    "encoding/json"
    "io/ioutil"

    "github.com/castevet6/dev-academy-names-go/pkg/models"
)

// internal package variant for holding the Name structs in array
var names []models.Name

// wrapper struct for providing methods
// no direct access to names array exposed
type NameModel struct {}

// create new NameModel struct
func New(fpath string) *NameModel {
    raw, err := ioutil.ReadFile(fpath)
    if err != nil {
        panic(err)
    }

    var namestruct models.Names

    err = json.Unmarshal(raw, &namestruct)
    if err != nil {
        panic(err)
    }

    // get the array from inside "names" wrapper attribute in JSON
    namearray := namestruct.Names
    // set the package internal variant value
    names = namearray

    return &NameModel{}
}

// return array of all Name structs
func (n *NameModel) GetAll() []models.Name {
    return names
}

// find Name struct by name, in case not found return empty struct Name{}
func (n *NameModel) GetName(name string) models.Name {
    for _, v := range names {
        if v.Name == name {
            return v
        }
    }
    return models.Name{}
}

func (n *NameModel) IsEmpty(name models.Name) bool {
    return name == (models.Name{})
}
