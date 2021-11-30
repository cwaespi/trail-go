tidy:
	go mod tidy

client-generate:
	go run github.com/prisma/prisma-client-go generate

db-sync:
	go run github.com/prisma/prisma-client-go migrate dev --name init
