import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user


from flask import jsonify
from src.schemas.user import UserSchema
from src.schemas.song import *
from src.models.http_exceptions import *
import src.repositories.users as users_repository
from src.schemas.errors import *

songs_url = "http://localhost:8092/songs/"  # URL de l'API users (golang)
ratings_url = "http://localhost:8089/ratings/"

def getAllRatingsbyIdSong(idSong):

    if not isSongIdValid(idSong):
        raise NotFound
    response_ratings = requests.request(method="GET", url=ratings_url)

# Filtré 
    resultat =[] 
    for item in response_ratings.json():
        if item["idSong"] == idSong:
            resultat.append(item)

    return resultat, response_ratings.status_code

def isSongIdValid(idSong):
    verif = requests.request(method="GET", url=songs_url+idSong)
    if verif.status_code == 200:
        return True
    else:
        return False
def isRatingIdValid(idRating):
    verif = requests.request(method="GET", url=ratings_url+idRating)
    if verif.status_code == 200:
        return True
    else:
        return False
def isUserIdtoRatingValid(idRating):
    verif = requests.request(method="GET", url=ratings_url+idRating)
    if verif.json()["idUser"] == current_user.id:
        return True
    else:
        return False

def addRatingbySong(idSong,rating_add):

    if not isSongIdValid(idSong):
        raise NotFound

    rating_add["idSong"] = idSong
    rating_add["idUser"] = current_user.id

    UpdateRating_schema = UpdateRatingSchema().loads(json.dumps(rating_add), unknown=EXCLUDE)

    response_ratings = requests.request(method="POST", url=ratings_url,json=UpdateRating_schema)

    return "", response_ratings.status_code

def deleteRatingtoSong(idSong,idRating):

    if not isSongIdValid(idSong) or not isRatingIdValid(idRating) or  not isUserIdtoRatingValid(idRating):
        raise NotFound

    response_ratings = requests.request(method="DELETE", url=ratings_url+idRating)

    return "", response_ratings.status_code

def getRatingtoSong(idSong,idRating):

    if not isSongIdValid(idSong) or not isRatingIdValid(idRating):
        raise NotFound
    
    response_ratings = requests.request(method="GET", url=ratings_url+idRating)

    return response_ratings.json(), response_ratings.status_code

def setRatingtoSong(idSong,idRating,rating_upt):

    if not isSongIdValid(idSong) or not isRatingIdValid(idRating) or  not isUserIdtoRatingValid(idRating):
        raise NotFound
    
    response_ratings = requests.request(method="PUT", url=ratings_url+idRating,json=rating_upt)

    return "", response_ratings.status_code