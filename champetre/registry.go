package champetre

import (
	"errors"
	"fmt"
)

// Contains the various functions to deal with the database and
// Index the collections
type registry struct {
	transactionHandler transactionHandler
	repositories map[Model]repository
	childObjects map[Model][]Model
	channel chan Transaction
}

func (rg *registry) Get(collection Model, filter func(Model) bool) ([]Model, error) {
	rep, ok := rg.repositories[collection]
	if !ok {
		return nil, errors.New("collection not found " + fmt.Sprint(collection.Default()))
	}
	results := []Model{}
	for _, el := range rep.elements {
		if filter(el) {
			results = append(results, el)
		}
	}
	return results, nil
}

func (rg *registry) Save(elem Model) error {
	// check if the Model belong correctly to a repository
	_, ok := rg.repositories[elem.Default()]
	if !ok {
		return errors.New("no collection found " + fmt.Sprint(elem.Default()))
	}
	// flatten the models present in the saved element
	elems := decompose(elem)
	// save all elements
	for _, el := range elems {
		localRep, ok := rg.repositories[el.Default()]
		if !ok {
			return errors.New("collection not found " + fmt.Sprint(el.Default()))
		}
		// if uuid is the default one then it is a new element to save
		if isUUIdDefault(el.Id()) {
			el.SetId(getUUId())
			localRep.Add(el)
			transaction := Transaction{
				Collection: el.Kind(),
				UUId: el.Id(),
				Parameter: Serialize(el),
				TransactionId: getUUId(),
				Kind: "create",
			}
			rg.channel <- transaction
		} else {
			// check if the element already exist
			matches, err := rg.Get(el.Default(), func(m Model) bool {return m.Id() == el.Id()})
			if err == nil || len(matches) > 1 {
				return errors.Join(err, errors.New("matches size " + fmt.Sprint(len(matches))))
			} else {
				// the document has been found
				// check if there have been changes, if yes, create the UpdateTransaction
				match := matches[0]
				if match != el {
					localRep.Replace(el)
					transaction := Transaction{
						Collection: el.Kind(),
						UUId: el.Id(),
						Parameter: Serialize(el),
						TransactionId: getUUId(),
						Kind: "update",
					}
					rg.channel <- transaction
				}
			}
		}
	}
	return nil
}

func (rg *registry) Delete(collection Model, filter func(Model) bool) ([]Model, error) {
	rep, ok := rg.repositories[collection]
	if !ok  {
		return nil, errors.New("collection not found " + fmt.Sprint(collection.Default()))
	}
	removedElements := []Model{}
	repositoryElements := []Model{}
	for _, el := range rep.elements {
		if filter(el) {
			removedElements = append(removedElements, el)
		} else {
			repositoryElements = append(repositoryElements, el)
		}
	}
	rep.elements = repositoryElements
	for _, el := range removedElements {
		transaction := Transaction{
			Collection: el.Kind(),
			UUId: el.Id(),
			Parameter: nil,
			TransactionId: getUUId(),
			Kind: "delete",
		}
		rg.channel <- transaction
	}
	return removedElements, nil
}

// return the set of registered type elements composing the parameter
func decompose(elem Model) []Model {
	// TODO decompose element
	return []Model{}
}