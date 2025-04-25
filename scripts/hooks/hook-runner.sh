#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

run_checks() {
    local check_type=$1
    echo "Running ${check_type} checks..."

    # Run tests
    echo "Running tests..."
    make test
    if [ $? -ne 0 ]; then
        echo -e "${RED}❌ Tests failed. Please fix the failing tests before ${check_type}.${NC}"
        exit 1
    fi

    # Run linter
    echo "Running linter..."
    make lint
    if [ $? -ne 0 ]; then
        echo -e "${RED}❌ Linting failed. Please fix the issues before ${check_type}.${NC}"
        exit 1
    fi

    echo -e "${GREEN}✅ All ${check_type} checks passed!${NC}"
    exit 0
}