package transactions

type ResponseStatus string

const (
	StatusFailedTransactionRollback ResponseStatus = "Status failed transaction rollback"
	StatusFailedTransactionStart    ResponseStatus = "Status failed transaction start"
	StatusSuccess                   ResponseStatus = "Status success"
	StatusFailedTransactionCommit   ResponseStatus = "Status failed transaction commit"
	StatusFailedTransactionScan     ResponseStatus = "Status failed transaction scan"
	StatusFailedTransaction         ResponseStatus = "Status failed transaction"
	StatusUnregisteredMigration     ResponseStatus = "Status unregistered migration"
	StatusUnregisteredMethod        ResponseStatus = "Status unregistered method"
	StatusFound                     ResponseStatus = "Status found"
	StatusNotEqual                  ResponseStatus = "Status not equal"
	StatusNotFound                  ResponseStatus = "Status not found"
	StatusInvalidResponse           ResponseStatus = "Status invalid response"
	StatusInvalidSyntax             ResponseStatus = "Status invalid syntax"
)
