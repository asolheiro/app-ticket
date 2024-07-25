# NestJs commands
run-dev:
	clear
	make db
	@cd partners-api-nestjs && npm run start:dev

gen-module:
	clear
	@cd partners-api-nestjs && nest g module ${NAME}

gen-resource:
	clear
	@cd partners-api-nestjs && nest g resource

gen-service:
	clear
	@cd partners-api-nestjs && nest g service ${NAME}

# Prisma commands
prisma-init:
	clear 
	make db
	@cd partners-api-nestjs && npx prisma init

prisma-mig-dev:
	clear
	make db 
	@cd partners-api-nestjs && npx prisma migrate dev

# Database commands
db:
	docker compose up -d

db-down:
	docker compose down
