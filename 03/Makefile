GENMOCK := make gen-mock

gen-mock-all:
	$(GENMOCK) source=ports/repo.go folder=mock_repository pkg=mockrepository


	$(GENMOCK) source=ports/service.go folder=mock_service pkg=mockservice

gen-mock:
	mockgen -source=$(source) -package=$(pkg) -destination=testings/$(folder)/$(pkg).go