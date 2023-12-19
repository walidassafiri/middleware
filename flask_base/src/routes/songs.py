import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from src.models.http_exceptions import *
from src.schemas.song import *
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
    try:
      return songs_service.getAllRatingsbyIdSong(id)
    except (NotFound):
      error = UnauthorizedSchema().loads("{}")
      return error, error.get("code")

    

@songs.route('/<id>/ratings', methods=['POST'])
@login_required
def addRatingbySong(id):
    """
    ---
    post:
      description: Register
      requestBody:
        required: true
        content:
            application/json:
                schema: UserRegister
      responses:
        '201':
          description: Created
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
        '403':
          description: Already logged in
          content:
            application/json:
              schema: Forbidden
            application/yaml:
              schema: Forbidden
        '409':
          description: User already exists
          content:
            application/json:
              schema: Conflict
            application/yaml:
              schema: Conflict
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - auth
          - users
    """
    # parser le body
    try:
        rating_add = SongUpdateSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")
    try:
      return songs_service.addRatingbySong(id,rating_add)
    except (NotFound):
      error = UnauthorizedSchema().loads("{}")
      return error, error.get("code")
