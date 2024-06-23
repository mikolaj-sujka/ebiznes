from flask import Flask, render_template, request, jsonify
import requests

app = Flask(__name__)

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/send_query', methods=['POST'])
def send_query():
    user_input = request.form['input']
    model_type = request.form['model_type']
    
    response = requests.post('http://localhost:5000/query', json={'input': user_input, 'model_type': model_type})
    
    return jsonify(response.json())

if __name__ == '__main__':
    app.run(port=5003)
