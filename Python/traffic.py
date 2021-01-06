from flask import Flask
from flask_restful import Api,Resource
from random import seed
from random import choice
import random
import json
import requests
app = Flask(__name__)

api = Api(app)
API_URI = 'http://dummy.sopes1grupo9vd20.tk/'
class HelloWorld(Resource):
    def get(self):
        f = open("./datos.json","r")
        listaCasos = json.load(f)
        #value = randint(0,len(listaCasos)-1)
        caso = json.dumps(choice(listaCasos))
        x = requests.post(API_URI,data = caso)
        return x.text
    

api.add_resource(HelloWorld,"/")
if __name__ == "__main__":
    app.run(debug=True)