.PHONY: build run tidy migrate-create login migrate-up migrate-down migrate-up-deploy migrate-down-deploy login build-ec2 init-deploy deploy-ec2-remote gen-abi

## run: Run the app.
run:
	@go run ./cmd/server

# Get list of .sql files from the up/ directory, sorted from old to new
MIGRATE_UP_FILES := $(shell ls migrations/up/*.sql | sort)

# Get list of .sql files from the down/ directory, sorted from new to old
MIGRATE_DOWN_FILES := $(shell ls migrations/down/*.sql | sort -r)

## migrate-up: Run all pending migrations to the database.
migrate-up:
	@echo "🚀 Starting UP migrations from /migrations/up..."
	@if [ -z "$(MIGRATE_UP_FILES)" ]; then \
		echo "No UP migration files found in migrations/up."; \
	else \
		for file in $(MIGRATE_UP_FILES); do \
			echo "--> Running UP: $$file"; \
			./bin/run_migration_file.sh $$file; \
		done; \
	fi
	@echo "✅ All UP migrations completed."

## migrate-down: Run all migrations to the database.
migrate-down:
	@echo "⏪ Starting DOWN migrations from /migrations/down..."
	@if [ -z "$(MIGRATE_DOWN_FILES)" ]; then \
		echo "No DOWN migration files found in migrations/down."; \
	else \
		for file in $(MIGRATE_DOWN_FILES); do \
			echo "--> Running DOWN: $$file"; \
			./bin/run_migration_file.sh $$file; \
		done; \
	fi
	@echo "✅ All DOWN migrations completed."


## create-migration NAME=<name>: Create a new migration file.
migrate-create:
	@if [ -z "$(NAME)" ]; then \
		echo "Usage: make create-migration NAME=<migration_name>"; \
		exit 1; \
	fi
	./bin/create_migration.sh $(NAME)


## migrate-up: Run all pending migrations to the database (aws).
migrate-up-deploy:
	@echo "🚀 Starting UP migrations from /migrations/up..."
	@if [ -z "$(MIGRATE_UP_FILES)" ]; then \
		echo "No UP migration files found in migrations/up."; \
	else \
		for file in $(MIGRATE_UP_FILES); do \
			echo "--> Running UP: $$file"; \
			./bin/run_migration_with_tunnel.sh $$file; \
		done; \
	fi
	@echo "✅ All UP migrations completed."

## migrate-down: Run all migrations to the database.
migrate-down-deploy:
	@echo "⏪ Starting DOWN migrations from /migrations/down..."
	@if [ -z "$(MIGRATE_DOWN_FILES)" ]; then \
		echo "No DOWN migration files found in migrations/down."; \
	else \
		for file in $(MIGRATE_DOWN_FILES); do \
			echo "--> Running DOWN: $$file"; \
			./bin/run_migration_with_tunnel.sh $$file; \
		done; \
	fi
	@echo "✅ All DOWN migrations completed."

# login to aws via ssh tunnel
login:
	@./bin/login.sh

# tidy go modules
tidy:
	go mod tidy

# build current or local machine
build: tidy
	go build -o dist/server ./cmd/server

# build AWS EC2 ARM64
build-ec2: tidy
	GOOS=linux GOARCH=arm64 go build -o dist/server ./cmd/server
	
# chmod +x bin/init-deploy
init-deploy:
	@echo "make[$@] init deploy from mac to ec2..."
	./bin/init_deploy.sh
	@echo "make[$@] done"

# chmod +x bin/remote-deploy
deploy-ec2-remote: build-ec2
	@echo "make[$@] build and deploy from mac to ec2..."
	./bin/remote_deploy.sh
	@echo "make[$@] done"

# generate Go bindings from ERC20 contract ABI
gen-erc20-abi:
	@echo "Generating Go bindings from ERC20 ABI..."
	@mkdir -p internal/driven-adapter/external/blockchain/gen
	@abigen --abi=contracts/erc20.abi --pkg=erc20 --type=ERC20 --out=internal/driven-adapter/external/blockchain/gen/erc20/erc20.go
	@echo "✅ Go bindings generated successfully!"

# generate Go bindings from swap contract ABI
gen-swap-abi:
	@echo "Generating Go bindings from Swap ABI..."
	@mkdir -p internal/driven-adapter/external/blockchain/gen
	@abigen --abi=contracts/swap.abi --pkg=swap --type=Swap --out=internal/driven-adapter/external/blockchain/gen/swap/swap.go
	@echo "✅ Go bindings generated successfully!"


