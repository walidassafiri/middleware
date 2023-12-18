import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from src.models.http_exceptions import *
from src.schemas.user import UserUpdateSchema
from src.schemas.errors import *
import src.services.songs as songs_service

# from routes import users
songs = Blueprint(name="songs", import_name=__name__)


@songs.route('/<id>/ratings', methods=['GET'])
@login_required
def get_ratingswithsong(id):
    """
    ---
    get:
      description: Getting a user
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of user id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: User
            application/yaml:
              schema: User
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - users
    """
    return songs_service.getAllRatingsbyIdSong(id)
