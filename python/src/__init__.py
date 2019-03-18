import os
from flask import Flask

app = Flask(__name__)
app.secret_key = "YL,hJDFCgb6E[.E&M4xEQrKuU9kYmbufT%qcxn#4LHqQFF{;"

from src import index
