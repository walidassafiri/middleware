# flask/services/songs.py
import json
import requests
import uuid
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

from src.schemas.user import UserSchema
from src.schemas.song import *
from src.models.song import Song as SongModel
from src.models.http_exceptions import *
import src.repositories.songs as songs_repository
import src.repositories.users as users_repository
from src.schemas.errors import *

# URL de l'API songs (golang)
songs_url = "http://localhost:8092/songs/"



def isSongIdValid(id):
    verif = requests.request(method="GET", url=songs_url+(id))
    if verif.status_code == 200:
        return True
    else:
        return False



def isUuidValid(id):
    try:
        uuid_obj = uuid.UUID(id)
        return uuid_obj.version == 4
    except ValueError:
        # La conversion en UUID a échoué
        return False



def isUserIdValid(id):
    verif = requests.request(method="GET", url=songs_url+id)
    if verif.json()["idUser"] == current_user.id:
        return True
    else:
        return False


def get_song(id):

    if not isUuidValid(id):
        raise UnprocessableEntity
    if not isSongIdValid(id):
        raise NotFound

    response = requests.get(songs_url + id)

    if response.status_code == 500:
        raise SomethingWentWrong  
    return response.json(), response.status_code


def create_song(song_data):
    
    song_schema = CreateSongSchema().loads(json.dumps(song_data), unknown=EXCLUDE)
    response = requests.request(method="POST", url=songs_url, json=song_schema)
    
    if response.status_code == 500:
        raise SomethingWentWrong

    return "", response.status_code


def get_song_from_db(song_id):
    return songs_repository.get_song(song_id)


def song_exists(song_id):
    return get_song_from_db(song_id) is not None

def update_song(id, song_update):
    if not isUuidValid(id):
        raise UnprocessableEntity
    if not isSongIdValid(id):
        raise NotFound

    response = requests.request(method="PUT", url=songs_url+id,json=song_update)

    if response.status_code == 500:
        raise SomethingWentWrong
    return response.json(), response.status_code



def get_songs():
    response = requests.get(songs_url)
    return response.json(), response.status_code


def delete_song(id):
    if not isSongIdValid(id):
        raise NotFound
   
    response=requests.request(method="DELETE", url=songs_url + id)

    if response.status_code == 500:
        raise SomethingWentWrong

    return "", response.status_code
