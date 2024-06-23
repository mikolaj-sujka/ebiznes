import os
from flask import Flask, request, jsonify
import openai
from openai import OpenAI
from dotenv import load_dotenv

app = Flask(__name__)

load_dotenv()

openai.api_key = os.getenv('OPENAI_API_KEY')

client = OpenAI()


@app.route('/gpt', methods=['POST'])
def gpt():
    data = request.json
    user_input = data.get('input')
    
    response = client.chat.completions.create(
        model="gpt-4",
        messages=[
            {"role": "system", "content": "You are a poetic assistant, skilled in explaining complex programming concepts with creative flair."},
            {"role": "user", "content": user_input}
        ]
    )
    # Access the response content correctly
    response_message = response.choices[0].message.content
    print(response_message )
    return jsonify({'response': response_message})


if __name__ == '__main__':
    app.run(port=5001)
