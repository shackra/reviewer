##
# reviewer
#
# @file
# @version 0.1

PROJECT_DIR := $(CURDIR)

build:
	@go build -o $(PROJECT_DIR)/main $(PROJECT_DIR)/cmd/reviewer

format:
	@gofumpt -extra -w $(PROJECT_DIR)
	@golines --chain-split-dots --ignore-generated \
		--reformat-tags --shorten-comments -w $(PROJECT_DIR)

.PHONY: build format

# end
