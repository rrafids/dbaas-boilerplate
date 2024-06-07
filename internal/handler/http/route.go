package http

func Route() {
	// Repository
	lambdaManagerHttpRepo := lambdaManagerHttpRepository.NewLambdaManagerHttpRepository()

	// Usecase
	lambdaManagerUsecase := usecase.NewAuthUsecase(lambdaManagerHttpRepo)

	// Handler
	lambdaManagerHandler := NewLambdaManagerHandler(lambdaManagerUsecase)

	// Route
	lambdaManagerV1Route := e.Group("/v1/lambda-managers")

	lambdaManagerV1Route.POST("/load-object", lambdaManagerHandler.LoadObject)
}