build: 
	@go build -o bin/exe ./cmd/Saigon-Inventory
run: build
	@./bin/exe
clean:
	@rm ./bin/*