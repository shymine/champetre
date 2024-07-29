package champetre

import (
	"log"
	"os"

	"github.com/fatih/structs"
	collectionsutils "github.com/shymine/collectionsutils"
	"golang.org/x/exp/maps"
)

// Is used to build the registry for the database
type RegistryBuilder struct {
	registeredType map[Model]string
	// should finish by a "/"
	databasePath string 
	databaseName string
}

// Used to register the collections used for the registry.
// model should be a default object of the concerned collection.
// All Model structs should be registered
func (rb *RegistryBuilder) Register(model Model) {
	_, ok := rb.registeredType[model]
	if ok {
		return 
	}
	rb.registeredType[model] = ""
}

// Set the path to the database
func (rb *RegistryBuilder) Database(dbPath string, dbName string) {
	rb.databasePath = dbPath
	rb.databaseName = dbName
}

func (rb *RegistryBuilder) Compile() registry {
	childObjects := map[Model][]Model{}
	repositories := map[Model]repository{}

	// for each registered type
	for k := range rb.registeredType {
		values := structs.Values(k)
		// for each of its attributes
		for _, att := range(values) {
			// if the type of the attribute is registered then 
			// we add the type k to the list of children for the type of the attribute
			m, is_model := att.(Model)
			if is_model {
				def := m.Default()
				_, exist := childObjects[def]
				if exist {
					childObjects[def] = append(childObjects[def], k.Default())
				} else {
					childObjects[def] = []Model{k.Default()}
				}
			}
		}
		// we create a repository for the type
		repositories[k.Default()] = repository{elements: []Model{}}
	}

	// setup the database, if it returns true, the the database already exists
	exists := rb.setupDatabase()

	// if it exists, then load the database
	if exists {
		rb.loadDatabase(repositories)
	}

	channel := make(chan Transaction, 20)
	// TODO change chan to file writing in transaction.json instead
	trHandler := transactionHandler{
			database: rb.databaseName,
			databasePath: rb.databasePath,
			transactions: []Transaction{},
			channel: channel,
		}
	registry := registry{
		transactionHandler: trHandler,
		repositories: repositories,
		childObjects: childObjects,
		channel: channel,
	}
	return registry
}

// setup the database folder system and files
// even if the database exist, if there are new collections in the registeredType, they re created
// else create the database folder, the transactions save file and logs
func (rb *RegistryBuilder) setupDatabase() bool {
	_, err := os.Stat(rb.databasePath+rb.databaseName)
	exist := err != nil
	// get Models Name
	registered := collectionsutils.Map(
		maps.Keys(rb.registeredType), 
		func(l Model) string {return l.Kind()},
	)

	if exist {
		registeredModelsToInitialize := rb.getModelsNotInitialized(rb.getLocalCollections(), registered)
		rb.initializeCollections(registeredModelsToInitialize)
	} else {
		if err := os.Mkdir(rb.databasePath+rb.databaseName, os.ModePerm); err != nil {
			log.Fatal(err)
		} else {
			rb.createFiles()
			rb.initializeCollections(registered)
		}
	}
	return exist
}

// load the elements from the database and populate the repositories
func (rb *RegistryBuilder) loadDatabase(repositories map[Model]repository) {
	// TODO loading of the database
	// The representation where the referenced documents are not taken in account
	intermediateRep := rb.setIntermediateRepresentation(repositories)

}
