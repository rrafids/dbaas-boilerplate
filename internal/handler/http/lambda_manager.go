package http

type LambdaManagerHandler struct {
	LambdaManagerUsecase usecase.LambdaManagerUsecase
}

func NewLambdaManagerHandler(LambdaManagerUsecase usecase.LambdaManagerUsecase) *LambdaManagerHandler {
	return &LambdaManagerHandler{
		LambdaManagerUsecase,
	}
}

func (handler *LambdaManagerHandler) LoadObject(c fiber.Context) error {
	// TODO: This layer will call lambda_manager usecase layer
	response, err := handler.LambdaManagerUsecase.LoadObject()
	if err != nil {
		return nil, err
	}

	return response, nil
}
