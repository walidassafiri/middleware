import json
import yaml

from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from src.models.http_exceptions import *
from src.schemas.song import *
from src.schemas.errors import *
import src.services.songs as songs_service


# from routes import songs
songs = Blueprint(name="songs", import_name=__name__)



def isYaml(Accept,donnees):

  if 'application/json' in Accept:
    return donnees
  elif 'application/yaml' in Accept:
    return  yaml.dump(donnees)
  else:
    return donnees


@songs.route('/<id>/ratings', methods=['GET'])
@login_required
def get_ratingswithsong(id):
    """
    ---
    get:
      description: Get Ratings using idSong
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of Song id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/UpdateRatingSchema"
            application/yaml:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/UpdateRatingSchema"
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: UnauthorizedSchema
            application/yaml:
              schema: UnauthorizedSchema
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFoundSchema
            application/yaml:
              schema: NotFoundSchema
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntitySchema
            application/yaml:
              schema: UnprocessableEntitySchema
      tags:
          - ratings
    """
    accept_header = request.headers.get('Accept')
    try:

      reponse, status = songs_service.getSongRatings(id)

      return isYaml(accept_header,reponse), status
    except (NotFound):
      error = NotFoundSchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")
    except (UnprocessableEntity):
      error = UnprocessableEntitySchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")




@songs.route('/<id>', methods=['GET'])
@login_required
def get_song(id):
    """
    ---
    get:
      description: Getting a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/GetSong"
            application/yaml:
              schema: 
                 type: array
                 items:
                  $ref: "#/components/schemas/GetSong"
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
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntitySchema
            application/yaml:
              schema: UnprocessableEntitySchema
      tags:
          - songs
    """
    accept_header = request.headers.get('Accept')
    try:
      reponse, status = songs_service.get_song(id)

      return isYaml(accept_header,reponse), status
    except (NotFound):
      error = NotFoundSchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")
    except (UnprocessableEntity):
      error = UnprocessableEntitySchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")


@songs.route('/<id>/ratings', methods=['POST'])
@login_required
def addRatingbySong(id):
    """
    ---
    post:
      description: add Rating by Song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of Song id
      requestBody:
        required: true
        content:
            application/json:
              schema: SongUpdateSchema
            application/yaml:
              schema: SongUpdateSchema
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: UpdateRatingSchema
            application/yaml:
              schema: UpdateRatingSchema
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: UnauthorizedSchema
            application/yaml:
              schema: UnauthorizedSchema
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFoundSchema
            application/yaml:
              schema: NotFoundSchema
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntitySchema
            application/yaml:
              schema: UnprocessableEntitySchema
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrongSchema
            application/yaml:
              schema: SomethingWentWrongSchema
      tags:
          - Raiting
          - Song
    """
    accept_header = request.headers.get('Accept')
    # parser le body
    try:
        rating_add = SongUpdateSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return isYaml(accept_header,error), error.get("code")
    try:

      reponse , status =songs_service.addRatingSong(id,rating_add)

      return isYaml(accept_header,reponse),status

    except (NotFound):
      error = NotFoundSchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")
    except (SomethingWentWrong):
      error = SomethingWentWrongSchema().loads("{}")
      return isYaml(accept_header,error) , error.get("code")
    except (UnprocessableEntity):
      error = UnprocessableEntitySchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")


@songs.route('/<id>/ratings/<rating_id>', methods=['DELETE'])
@login_required
def DeleteRatingtoSong(id,rating_id):
    """
    ---
    delete:
      description: Delete Rating to Song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of Song id
        - in: path
          name: rating_id
          schema:
            type: uuidv4
          required: true
          description: UUID of Rating id
      responses:
        '204':
          description: No content
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: UnauthorizedSchema
            application/yaml:
              schema: UnauthorizedSchema
        '403':
          description: Forbidden, rating not theirs
          content:
            application/json:
              schema: ForbiddenSchema
            application/yaml:
              schema: ForbiddenSchema
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFoundSchema
            application/yaml:
              schema: NotFoundSchema
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrongSchema
            application/yaml:
              schema: SomethingWentWrongSchema
      tags:
          - Raiting
          - Song
    """
    accept_header = request.headers.get('Accept')
    # parser le body
    try:

      reponse, status =songs_service.deleteRatingtoSong(id,rating_id)

      return isYaml(accept_header,reponse), status
    except (NotFound):
      error = NotFoundSchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")
    except (Forbidden):
      error = ForbiddenSchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")
    except (SomethingWentWrong):
      error = SomethingWentWrongSchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")
    

@songs.route('/<id>/ratings/<rating_id>', methods=['GET'])
@login_required
def GetRatingtoSong(id,rating_id):
    """
    ---
    post:
      description: get rating
      parameters:
      - in: path
        name: id
        schema:
          type: uuidv4
        required: true
        description: UUID of Song id
      - in: path
        name: rating_id
        schema:
          type: uuidv4
        required: true
        description: UUID of Rating id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: UpdateRatingSchema
            application/yaml:
              schema: UpdateRatingSchema
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: UnauthorizedSchema
            application/yaml:
              schema: UnauthorizedSchema
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFoundSchema
            application/yaml:
              schema: NotFoundSchema
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntitySchema
            application/yaml:
              schema: UnprocessableEntitySchema
      tags:
          - Raiting
          - Song
    """
    accept_header = request.headers.get('Accept')
    # parser le body
    try: 

      reponse , status=songs_service.getRatingtoSong(id,rating_id)

      return isYaml(accept_header,reponse), status
    except (NotFound):
      error = NotFoundSchema().loads("{}")
      return isYaml(accept_header,error) , error.get("code")
    except (UnprocessableEntity):
      error = UnprocessableEntitySchema().loads("{}")
      return isYaml(accept_header,error) , error.get("code")

@songs.route('/<id>/ratings/<rating_id>', methods=['PUT'])
@login_required
def SetRatingtoSong(id,rating_id):
    """
    ---
    post:
      description: Register
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of Song id
        - in: path
          name: rating_id
          schema:
            type: uuidv4
          required: true
          description: UUID of Rating id
      requestBody:
        required: true
        content:
            application/json:
                schema: SongUpdateSchema
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: UpdateRatingSchema
            application/yaml:
              schema: UpdateRatingSchema
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: UnauthorizedSchema
            application/yaml:
              schema: UnauthorizedSchema
        '403':
          description: Already logged in
          content:
            application/json:
              schema: ForbiddenSchema
            application/yaml:
              schema: ForbiddenSchema
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFoundSchema
            application/yaml:
              schema: NotFoundSchema
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntitySchema
            application/yaml:
              schema: UnprocessableEntitySchema
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrongSchema
            application/yaml:
              schema: SomethingWentWrongSchema
      tags:
          - Raiting
          - Song
    """
    accept_header = request.headers.get('Accept')
    # parser le body
    try:
      rating_upt = SongUpdateSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
      error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
      return isYaml(accept_header,error), error.get("code")

    try:

      reponse, status = songs_service.setRatingtoSong(id,rating_id,rating_upt)
      
      return isYaml(accept_header,reponse), status
    except (NotFound):
      error = NotFoundSchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")
    except (Forbidden):
      error = ForbiddenSchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")
    except (SomethingWentWrong):
      error = SomethingWentWrongSchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")
    except (UnprocessableEntity):
      error = UnprocessableEntitySchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")



@songs.route('', methods=['GET'])
@login_required
def getAllSongs():
    """
    ---
    get:
      description: All song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true

          description: UUID of song id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/UpdateRatingSchema"
            application/yaml:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/UpdateRatingSchema"
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
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntitySchema
            application/yaml:
              schema: UnprocessableEntitySchema
      tags:
          - songs
    """
    accept_header = request.headers.get('Accept')
    try:

      reponse, status = songs_service.get_songs()

      return isYaml(accept_header,reponse), status
    except (NotFound):
      error = NotFoundSchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")
    except (UnprocessableEntity):
      error = UnprocessableEntitySchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")



@songs.route('/<id>', methods=['PUT'])
@login_required
def update_song(id):
    """
    ---
    put:
      description: Updating a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      requestBody:
        required: true
        content:
            application/json:
                schema: UpdateSong
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
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
              schema: ForbiddenSchema
            application/yaml:
              schema: ForbiddenSchema
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
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
              schema: SomethingWentWrongSchema
            application/yaml:
              schema: SomethingWentWrongSchema
      tags:
          - songs
    """
    accept_header = request.headers.get('Accept')

    # parser le body
    try:
      song_data = UpdateSongSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
      error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
      return isYaml(accept_header,error), error.get("code")

    try:

      reponse, status = songs_service.update_song(id, song_data)
      
      return isYaml(accept_header,reponse), status
    except (NotFound):
      error = NotFoundSchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")
    except (Forbidden):
      error = ForbiddenSchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")
    except (SomethingWentWrong):
      error = SomethingWentWrongSchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")
    except (UnprocessableEntity):
      error = UnprocessableEntitySchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")





@songs.route('/<id>', methods=['DELETE'])
@login_required
def delete_song(id):
    """
    ---
    delete:
      description: Supprimer une chanson
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID de l'identifiant de la chanson
      requestBody:
        required: true
        content:
            application/json:
                schema: DeleteSong  # Utilisez le schéma de suppression
      responses:
        '204':
          description: Pas de contenu
        '401':
          description: Non autorisé
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '403':
          description: Forbidden, rating not theirs
          content:
            application/json:
              schema: ForbiddenSchema
            application/yaml:
              schema: ForbiddenSchema
        '404':
          description: Non trouvé
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrongSchema
            application/yaml:
              schema: SomethingWentWrongSchema
      tags:
          - songs
    """
    accept_header = request.headers.get('Accept')
    try:

      reponse, status = songs_service.delete_song(id)

      return isYaml(accept_header,reponse), status
    except (NotFound):
      error = NotFoundSchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")
    except (Forbidden):
      error = ForbiddenSchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")
    except (SomethingWentWrong):
      error = SomethingWentWrongSchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")





@songs.route('', methods=['POST'])
@login_required
def create_song():
    """
    ---
    post:
      description: Creating a new song
      requestBody:
        required: true
        content:
            application/json:
              schema: CreateSongSchema
            application/yaml:
              schema: CreateSongSchema
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: SongSchema
            application/yaml:
              schema: SongSchema
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: UnauthorizedSchema
            application/yaml:
              schema: UnauthorizedSchema
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntitySchema
            application/yaml:
              schema: UnprocessableEntitySchema
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrongSchema
            application/yaml:
              schema: SomethingWentWrongSchema
      tags:
          - Song
    """
    accept_header = request.headers.get('Accept')

    # parser le body
    try:
      song_data = UpdateSongSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
      error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
      return isYaml(accept_header,error), error.get("code")


    try:
      reponse, status = songs_service.create_song(song_data)
      return isYaml(accept_header,reponse), status
    except (NotFound):
      error = NotFoundSchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")
    except (Forbidden):
      error = ForbiddenSchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")
    except (SomethingWentWrong):
      error = SomethingWentWrongSchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")
    except (UnprocessableEntity):
      error = UnprocessableEntitySchema().loads("{}")
      return isYaml(accept_header,error), error.get("code")
