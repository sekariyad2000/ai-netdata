.PHONY: setup run clean test install-deps

# Python virtual environment name
VENV = venv

# Default Python interpreter
PYTHON = python3

# Default port
PORT = 5000

setup: install-deps
	@echo "Setup complete! Run 'make run' to start the application."

install-deps:
	@echo "Creating virtual environment..."
	$(PYTHON) -m venv $(VENV)
	@echo "Installing dependencies..."
	. $(VENV)/bin/activate && pip install -r requirements.txt

run:
	@echo "Starting the application..."
	. $(VENV)/bin/activate && python app.py

clean:
	@echo "Cleaning up..."
	rm -rf $(VENV)
	find . -type d -name "__pycache__" -exec rm -r {} +
	find . -type f -name "*.pyc" -delete

test:
	@echo "Running tests..."
	. $(VENV)/bin/activate && python -m pytest

# Help command
help:
	@echo "Available commands:"
	@echo "  make setup        - Set up the virtual environment and install dependencies"
	@echo "  make run         - Run the application"
	@echo "  make clean       - Clean up virtual environment and cache files"
	@echo "  make test        - Run tests"
	@echo "  make help        - Show this help message" 