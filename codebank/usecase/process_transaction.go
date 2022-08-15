package usecase

type UseCaseTransaction struct {
	TransactionRepository domain.TransactionRepository
}

func NewUseCaseTransaction(TransactionRepository domain.TransactionRepository) UseCaseTransaction {
	return UseCaseTransaction(TransactionRepository: TransactionRepository)
}

func (u UseCaseTransaction) ProcessTransaction (transactionDto dto.Transaction) (domain.Transaction, error) {
	
}