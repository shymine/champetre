package champetre

import "errors"

// TODO make it so insterad of reading from chan, it read every few minutes from transactions.json
type transactionHandler struct {
	databasePath string
	database     string
	transactions []Transaction
	channel      chan Transaction
}

func (th *transactionHandler) Push(transaction Transaction) {
	// TODO push transaction
}

func (th *transactionHandler) Pop(trId string) error {
	// parcours la liste & trouve l id correspondant
	// retire l indice et récupère la transaction
	// th.save()
	// th.log(transaction)
	// retourne une erreur si aucune transaction ne match trId
	// TODO Pop transaction
	return errors.New("trId not found" + trId)
}

// save the current list of transaction to do
func (th *transactionHandler) save() {
	// TODO save function
}

// add the current transaction to the history of done transactions
func (th *transactionHandler) log(transaction Transaction) {
	// TODO log function
}

// load the transaction history and recreate the transactions
func (th *transactionHandler) load() {
	// TODO load function
}

// resolve the different transactions in the base files depending on its kind
func (th *transactionHandler) resolveTransaction() {
	// TODO resolve transaction
}

// the run function that push and pop the channels and resolve the transactions
func (th *transactionHandler) Run() { 
	// TODO run function
}