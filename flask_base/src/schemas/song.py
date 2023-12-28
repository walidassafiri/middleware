# schemas/song.py
from marshmallow import Schema, fields, validates_schema, ValidationError, validate


class SongSchema(Schema):
    id = fields.String(description="UUID")
    artist = fields.String(description="artist")
    title = fields.String(description="title")
    album = fields.String(description="album")
    content = fields.String(description="content")

    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("artist") or obj.get("artist") == "") and \
               (not obj.get("title") or obj.get("title") == "") and \
               (not obj.get("album") or obj.get("album") == "") and \
               (not obj.get("content") or obj.get("content") == "")  


class BaseSongSchema(Schema):
    artist = fields.String(description="artist")
    title = fields.String(description="title")
    album = fields.String(description="album")
    content = fields.String(description="content")


# Schéma pour la création d'une chanson
class CreateSongSchema(SongSchema):
    artist = fields.String(required=True, validate=validate.Length(max=255))
    title = fields.String(required=True, validate=validate.Length(max=255))
    album = fields.String(required=True, validate=validate.Length(max=255))
    content = fields.String(required=True, validate=validate.Length(max=255))


# Schéma pour la mise à jour d'une chanson
class UpdateSongSchema(SongSchema):
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("artist" in data and data["artist"] != "") or
                ("title" in data and data["title"] != "") or
                ("album" in data and data["album"] != "") or
                ("content" in data and data["content"] != "")):
            raise ValidationError("at least one of ['artist','title','album','content'] must be specified")


class GetSongSchema(SongSchema):
    id = fields.UUID(required=True)



class DeleteSongSchema(SongSchema):
    id = fields.UUID(required=True)