package champetre

type Transaction struct {
	Collection    string
	UUId          string
	Parameter     map[string]any
	TransactionId string
	Kind          string
}