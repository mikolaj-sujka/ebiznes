from flask import Flask, request, jsonify
import requests

app = Flask(__name__)

# Endpoint do wysyłania zapytań
@app.route('/query', methods=['POST'])
def query():
    data = request.json
    user_input = data.get('input')
    model_type = data.get('model_type', 'gpt')  # Domyślnie używamy ChatGPT

    if model_type == 'gpt':
        response = requests.post('http://localhost:5001/gpt', json={'input': user_input})
    else:
        response = requests.post('http://localhost:5002/llama2', json={'input': user_input})

    return jsonify(response.json())

if __name__ == '__main__':
    app.run(port=5000)
