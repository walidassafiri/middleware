# flask/services/songs.py
import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

from src.schemas.song import SongSchema
from src.models.song import Song as SongModel
from src.models.http_exceptions import *
import src.repositories.songs as songs_repository

# URL de l'API songs (golang)
songs_url = "http://localhost:8092/songs/"


def get_song(id):
    response = requests.get(songs_url + id)
    return response.json(), response.status_code


def create_song(song_data):
    # Créer un modèle Song à partir des données reçues
    song_model = SongModel.from_dict(song_data)
    # Créer un schéma Song pour la requête vers l'API songs
    song_schema = SongSchema().loads(json.dumps(song_data), unknown=EXCLUDE)

    # Créer la chanson du côté de l'API songs
    response = requests.post(songs_url, json=song_schema)
    if response.status_code != 201:
        return response.json(), response.status_code

    # Ajouter la chanson dans la base de données pour maintenir la cohérence
    try:
        song_model.id = response.json()["id"]
        songs_repository.add_song(song_model)
    except Exception:
        raise SomethingWentWrong

    return response.json(), response.status_code


def update_song(id, song_data):
    if id != current_user.id:
        raise Forbidden

    # Créer un schéma Song pour la requête vers l'API songs
    song_schema = SongSchema().loads(json.dumps(song_data), unknown=EXCLUDE)
    response = None

    # Si des données doivent être modifiées du côté de l'API, lancer la requête de modification
    if not SongSchema.is_empty(song_schema):
        response = requests.put(songs_url + id, json=song_schema)
        if response.status_code != 200:
            return response.json(), response.status_code

    # Si des données doivent être modifiées du côté de la base de données
    song_model = SongModel.from_dict(song_data)
    if not song_model.is_empty():
        song_model.id = id
        found_song = songs_repository.get_song_from_id(id)
        # Assurez-vous que les champs non modifiés ne sont pas écrasés
        if not song_model.title:
            song_model.title = found_song.title
        # Ajoutez d'autres champs si nécessaire

        try:
            songs_repository.update_song(song_model)
        except exc.IntegrityError as e:
            if "NOT NULL" in e.orig.args[0]:
                raise UnprocessableEntity
            raise Conflict

    return (response.json(), response.status_code) if response else get_song(id)


def get_songs():
    response = requests.get(songs_url)
    return response.json(), response.status_code


def delete_song(id):
    if id != current_user.id:
        raise Forbidden

    # Lancer la requête de suppression du côté de l'API songs
    response = requests.delete(songs_url + id)
    if response.status_code != 204:
        return response.json(), response.status_code

    # Supprimer la chanson de la base de données
    try:
        songs_repository.delete_song(id)
    except Exception:
        raise SomethingWentWrong

    return "", 204
