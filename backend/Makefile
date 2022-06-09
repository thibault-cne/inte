NAME=go-backend-server
VERSION=0.0.1

.PHONY: build
## build: Compile the packages.
build:
	@echo "${SBLUE}Building ${IGREY}${NAME} ${VERSION}${S}${S} ⚙"
	@go build -o $(NAME)
	@echo ""
	@echo "${GREEN}Building Completed ${IGREY}${NAME} ${VERSION}${S}${S} 🎉"
	@echo ""

.PHONY: run
## run: Build and Run in development mode.
run: build
	@echo "${RED}Running ${IGREY}${NAME} ${VERSION}${S} in development mode${S} 🚀"
	@./$(NAME) -e development

.PHONY: run-prod
## run-prod: Build and Run in production mode.
run-prod: build
	@echo "Running ${NAME} ${VERSION} in production mode"
	@./$(NAME) -e production

.PHONY: clean
## clean: Clean project and previous builds.
clean:
	@echo "Cleaning project..."
	@rm -f $(NAME)

.PHONY: deps
## deps: Download modules
deps:
	@echo "Downloading dependencies..."
	@go mod download

.PHONY: test
## test: Run tests with verbose mode
test:
	@echo "${GREEN}Running tests...{S}"
	@go test -v ./tests/*

.PHONY: help
all: help
# help: show this help message
help: Makefile
	@echo
	@echo " Choose a command to run in "$(APP_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo


# Colors and styles
S 		=		\033[0m
BOLD 	= 		\033[1m
ITALIC 	= 		\033[3m
UNDER 	= 		\033[4m
REV 	= 		\033[7m

# Colors
GREY 	= 		\033[30m
RED 	= 		\033[31m
GREEN	=		\033[32m
YELLOW	=		\033[33m
BLUE	=		\033[34m
PURPLE	=		\033[35m
CYAN	=		\033[36m
WHITE	=		\033[37m

SGREY	=		\033[90m
SRED	=		\033[91m
SGREEN	=		\033[92m
SYELLOW	=		\033[93m
SBLUE	=		\033[94m
SPURPLE	=		\033[95m
SCYAN	=		\033[96m
SWHITE	=		\033[97m

# Colored backgrounds

IGREY	=		\033[40m
IRED	=		\033[41m
IGREEN	=		\033[42m
IYELLOW	=		\033[43m
IBLUE	=		\033[44m
IPURPLE	=		\033[45m
ICYAN	=		\033[46m
IWHITE	=		\033[47m