package champetre

import (
	"io/fs"
	"log"
	"os"
	"reflect"
	"slices"

	collectionutils "github.com/shymine/collectionsutils"
)

func (rb *RegistryBuilder) setIntermediateRepresentation(repositories map[Model]repository) map[Model][]Model {
	intermediateRep := map[Model][]Model{}
	// For each registered Models
	for model := range repositories {
		// Load the corresponding collection in the database
		loadedElems := rb.loadCollection(model.Kind())
		tmpRepo := []Model{}
		modelParams := model.Parameters()
		// For each loaded element
		for _, el := range loadedElems {
			tmpModel := model.Default()
			// For each parameter
			for name, v := range modelParams {
				vof := reflect.ValueOf(v)
				switch vof.Kind() {
				case reflect.Struct:
					t := el[name].(map[string]any)
					loadStruct(t, tmpModel, name)
				case reflect.Slice, reflect.Array:
					t := el[name].([]any)
					loadSlice(t, tmpModel, name)
				case reflect.Map:
					t := el[name].(map[string]any)
					loadMap(t, tmpModel, name)
				default:
					loadBaseType(el[name], tmpModel, name)
				}
			}
			tmpRepo = append(tmpRepo, tmpModel)
		}
		intermediateRep[model] = tmpRepo
	}
	return intermediateRep
}

func loadStruct(elem map[string]any, object any, fieldname string) {
	// TODO load struct from json
}

func loadSlice(elem []any, object any, fieldname string) {
	// TODO load slice from json
}

func loadMap(elem map[string]any, object any, fieldname string) {
	// TODO load map from json
}

func loadBaseType(elem any, object any, fieldname string) {
	// TODO load bool, int, float, string from json
}

func (rb *RegistryBuilder) getLocalCollections() []string {
	entries, err := os.ReadDir(rb.databasePath + rb.databaseName + "/collections")
	if err != nil {
		log.Fatal(err)
	}
	return collectionutils.Map(
		entries,
		func(l fs.DirEntry) string { return l.Name()},
	)
}

func (rb *RegistryBuilder) getModelsName() []string {
	res := make([]string, len(rb.registeredType))
	
	i := 0
	for k, _ := range rb.registeredType {
		res[i] = k.Kind()
		i++
	}
	return res
}

// get the elements from registered that are not in collections
func (rb *RegistryBuilder) getModelsNotInitialized(collections []string, registered []string) []string {
	res := []string{}
	for _, el := range registered {
		if !slices.Contains(collections, el) {
			res = append(res, el)
		}
	}
	return res
}

func (rb *RegistryBuilder) initializeCollections(collectionNames []string) {
	for _, el := range collectionNames {
		if err := os.Mkdir(rb.databasePath+rb.databaseName+"/collections/"+el, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}

func (rb *RegistryBuilder) createFiles() {
	_, err := os.Create(rb.databasePath + rb.databaseName + "transactions.log")
	if err != nil {
		log.Fatal(err)
	}
	trJson, err := os.Create(rb.databasePath + rb.databaseName + "transactions.json")
	if err != nil {
		log.Fatal(err)
	}
	_, err = trJson.WriteString("[]")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(rb.databasePath+rb.databaseName+"/collections", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

// Load the jsonify elements from the database
func (rb *RegistryBuilder) loadCollection(kind string) []map[string]any {
	// TODO load a collection from the database
	return []map[string]any{}
}