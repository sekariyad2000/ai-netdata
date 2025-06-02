# Netdata + Ollama Dashboard

A web application that combines Netdata system metrics with Ollama LLM capabilities, providing an interactive dashboard for system monitoring and AI-powered assistance.

## Prerequisites

- Python 3.8 or higher
- Netdata running locally (default: http://localhost:19999)
- Ollama running locally (default: http://localhost:11434)
- The llama2 model pulled in Ollama (`ollama pull llama2`)
- Make (for using Makefile commands)

## Installation

### Using Make (Recommended)

1. Clone this repository:
```bash
git clone <repository-url>
cd <repository-name>
```

2. Set up the environment and install dependencies:
```bash
make setup
```

3. Run the application:
```bash
make run
```

### Manual Installation

1. Clone this repository:
```bash
git clone <repository-url>
cd <repository-name>
```

2. Create a virtual environment and activate it:
```bash
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
```

3. Install the required packages:
```bash
pip install -r requirements.txt
```

4. Create a `.env` file (optional, for custom configuration):
```bash
NETDATA_URL=http://localhost:19999
OLLAMA_URL=http://localhost:11434
```

## Running the Application

1. Make sure Netdata and Ollama are running on your system.

2. Start the Flask application:
```bash
make run
```

3. Open your web browser and navigate to:
```
http://localhost:5000
```

## Available Make Commands

- `make setup` - Set up the virtual environment and install dependencies
- `make run` - Run the application
- `make clean` - Clean up virtual environment and cache files
- `make test` - Run tests
- `make help` - Show all available commands

## Features

- Real-time system metrics from Netdata
- Interactive chat interface with Ollama LLM
- Modern Bootstrap-based UI
- Responsive design for all devices
- Real-time updates of system metrics

## Usage

1. The dashboard shows system metrics on the left side, updating every 5 seconds.
2. Use the chat interface on the right to ask questions about your system metrics.
3. The AI assistant can help interpret the metrics and provide insights.

## Troubleshooting

- If you can't connect to Netdata, ensure it's running and accessible at the configured URL.
- If the chat doesn't work, verify that Ollama is running and the llama2 model is installed.
- Check the browser console for any JavaScript errors.
- Verify that all required ports (5000, 19999, 11434) are available and not blocked by your firewall.
- If Make commands fail, ensure you have Make installed on your system. 