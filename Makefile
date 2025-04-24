ifndef version
override version = v1
endif

.PHONY: runTCP
runTCP:
	go run ./tcp/${version}/main.go

.PHONY: runDS
runDS:
	go run ./ds/${type}/${version}/main.go	