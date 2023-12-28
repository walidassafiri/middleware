# flask/services/songs.py
import json
import requests
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


def get_song(id):
    response = requests.get(songs_url + id)
    return response.json(), response.status_code


def create_song(song_data):
    song_schema = CreateSongSchema().loads(json.dumps(song_data), unknown=EXCLUDE)
    response = requests.request(method="POST", url=songs_url, json=song_schema)

    if response.status_code != 201:
        return response.json(), response.status_code

    # Ajouter la chanson dans la base de données
    try:
        song_model = SongSchema().load(response.json())
        songs_repository.create_song(song_model)
    except Exception:
        raise SomethingWentWrong

    return response.json(), response.status_code



def update_song(id, updated_data):
    song_schema = UpdateSongSchema().loads(json.dumps(updated_data), unknown=EXCLUDE)
    response = requests.request(method="PUT", url=songs_url + id, json=song_schema)

    if response.status_code != 200:
        return response.json(), response.status_code

    # Mettre à jour la chanson dans la base de données
    try:
        updated_song_model = SongSchema().load(response.json())
        songs_repository.update_song(updated_song_model)
    except Exception:
        raise SomethingWentWrong

    return response.json(), response.status_code



def get_songs():
    response = requests.get(songs_url)
    return response.json(), response.status_code


def delete_song(id):
    response=requests.request(method="DELETE", url=songs_url + id)

    try:
        songs_repository.delete_song(id)
    except Exception:
        error = NotFoundSchema().loads("{}")
        return error, error.get("code")

    return response.json(), 200
