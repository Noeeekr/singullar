package transactions

type ResponseStatus string

const (
	StatusFailedTransactionRollback ResponseStatus = "StatusFailedTransactionRollback"
	StatusFailedTransactionStart    ResponseStatus = "StatusFailedTransactionStart"
	StatusSuccess                   ResponseStatus = "StatusSuccess"
	StatusFailedTransactionCommit   ResponseStatus = "StatusFailedTransactionCommit"
	StatusFailedTransactionScan     ResponseStatus = "StatusFailedTransactionScan"
	StatusFailedTransaction         ResponseStatus = "StatusFailedTransaction"
	StatusUnregisteredMigration     ResponseStatus = "StatusUnregisteredMigration"
	StatusUnregisteredMethod        ResponseStatus = "StatusUnregisteredMethod"
	StatusFound                     ResponseStatus = "StatusFound"
	StatusNotEqual                  ResponseStatus = "StatusNotEqual"
	StatusNotFound                  ResponseStatus = "StatusNotFound"
	StatusInvalidResponse           ResponseStatus = "StatusInvalidResponse"
	StatusInvalidSyntax             ResponseStatus = "StatusInvalidSyntax"
)
