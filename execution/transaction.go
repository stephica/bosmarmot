package execution

type Transaction interface {
	Check()
	Deliver()
}

type TransactionDecoder interface {
	Decode(tx []byte) (Transaction, error)
}
