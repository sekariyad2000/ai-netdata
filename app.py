from flask import Flask, render_template, jsonify, request
from flask_socketio import SocketIO
import requests
import json
import os
from dotenv import load_dotenv

# Load environment variables
load_dotenv()

app = Flask(__name__)
socketio = SocketIO(app)

# Configuration
NETDATA_URL = os.getenv('NETDATA_URL', 'http://localhost:19999')
OLLAMA_URL = os.getenv('OLLAMA_URL', 'http://localhost:11434')

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/api/metrics')
def get_metrics():
    try:
        # Get system metrics from Netdata
        response = requests.get(f"{NETDATA_URL}/api/v1/data?chart=system.cpu")
        return jsonify(response.json())
    except Exception as e:
        return jsonify({"error": str(e)}), 500

@app.route('/api/chat', methods=['POST'])
def chat():
    try:
        data = request.json
        prompt = data.get('prompt')
        
        # Call Ollama API
        response = requests.post(
            f"{OLLAMA_URL}/api/generate",
            json={
                "model": "llama2",
                "prompt": prompt,
                "stream": False
            }
        )
        
        return jsonify(response.json())
    except Exception as e:
        return jsonify({"error": str(e)}), 500

if __name__ == '__main__':
    socketio.run(app, debug=True, host='0.0.0.0', port=5000) 