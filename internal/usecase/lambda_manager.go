package usecase

type LambdaManagerUsecase interface {
	LoadObject(ctx context.Context, parameter) (errResponse *port.ErrorResponse)
}

type LambdaManagerUsecase struct {
	LambdaManagerHttpRepository httpRepository.LambdaManagerHttpRepository
}

func NewLambdaManagerUsecase(lambdaManagerHttpRepository httpRepository.LambdaManagerHttpRepository) *LambdaManagerUsecase {
	return &LambdaManagerUsecase{
		LambdaManagerHttpRepository: lambdaManagerHttpRepository
	}
}

func (usecase *LambdaManagerUsecase) LoadObject(ctx context.Context, parameter) (errResponse *port.ErrorResponse) {
	// TODO: This layer will call lambda_manager repository layer
	response, err := usecase.LambdaManagerHttpRepository.LoadObject()
	if err != nil {
		return nil, err
	}

	return response, nil
}
