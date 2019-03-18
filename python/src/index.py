import os

from flask import jsonify
from flask import session

from src import app

@app.route('/', methods=['GET'])
def index():
  if 'count' not in session:
    session['count'] = 0
  session['count'] += 1
  data = {
    'message': 'hello python, flask',
    'count': session['count']
  }
  return jsonify(data)
