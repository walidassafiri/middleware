import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

from flask import jsonify
from src.schemas.user import UserSchema
from src.models.user import User as UserModel
from src.models.http_exceptions import *
import src.repositories.users as users_repository
from src.schemas.errors import *

songs_url = "http://localhost:8092/songs/"  # URL de l'API users (golang)
ratings_url = "http://localhost:8089/ratings/"

def getAllRatingsbyIdSong(idSong):
    response_ratings = requests.request(method="GET", url=ratings_url)

# Filtré 
    resultat =[] 
    for item in response_ratings.json():
        if item["idSong"] == idSong:
            resultat.append(item)

    return resultat, response_ratings.status_code

