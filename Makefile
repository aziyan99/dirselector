APP_NAME := dirselector
SRC_DIR := ./cmd/$(APP_NAME)
OUTPUT_DIR := ./build
MODE ?= prod

ifeq ($(MODE), prod)
    BUILD_FLAGS := -ldflags "-s -w" # Strips debug info for smaller binary
    OUTPUT := $(OUTPUT_DIR)/$(APP_NAME).exe
else ifeq ($(MODE), debug)
    BUILD_FLAGS := -gcflags "all=-N -l" # Disables optimizations for debugging
    OUTPUT := $(OUTPUT_DIR)/$(APP_NAME)-debug.exe
else ifeq ($(MODE), dev)
    BUILD_FLAGS :=
    OUTPUT := $(OUTPUT_DIR)/$(APP_NAME)-dev.exe
else
    $(error Invalid MODE: $(MODE))
endif

.PHONY: all prod debug dev clean run

all: prod

prod:
	@echo "Building for production..."
	go build -mod=readonly $(BUILD_FLAGS) -o $(OUTPUT) $(SRC_DIR)
	@echo "Production build output: $(OUTPUT)"

debug:
	@echo "Building for debug..."
	go build -mod=readonly $(BUILD_FLAGS) -o $(OUTPUT) $(SRC_DIR)
	@echo "Debug build output: $(OUTPUT)"

dev:
	@echo "Building for development..."
	go build -mod=readonly $(BUILD_FLAGS) -o $(OUTPUT) $(SRC_DIR)
	@echo "Development build output: $(OUTPUT)"

run:
	@echo "Running in $(MODE) mode..."
	$(OUTPUT)

clean:
	@echo "Cleaning up build artifacts..."
	del /Q $(OUTPUT_DIR)\*
	@echo "Clean complete."
