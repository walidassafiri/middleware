# schemas/song.py
from marshmallow import Schema, fields, validates_schema, ValidationError


class SongSchema(Schema):
    id = fields.String(description="UUID")
    Artist = fields.String(description="Artist")
    Title = fields.String(description="Title")
    Album = fields.String(description="Album")
    Content = fields.String(description="Content")

    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("Artist") or obj.get("Artist") == "") and \
               (not obj.get("Title") or obj.get("Title") == "") and \
               (not obj.get("Album") or obj.get("Album") == "") and \
               (not obj.get("Content") or obj.get("Content") == "")  


class BaseSongSchema(Schema):
    Artist = fields.String(description="Artist")
    Title = fields.String(description="Title")
    Album = fields.String(description="Album")
    Content = fields.String(description="Content")


# Schéma pour la création d'une chanson
class SongCreateSchema(SongSchema):
    pass


# Schéma pour la mise à jour d'une chanson
class SongUpdateSchema(SongSchema):
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("Artist" in data and data["Artist"] != "") or
                ("Title" in data and data["Title"] != "") or
                ("Album" in data and data["Album"] != "") or
                ("Content" in data and data["Content"] != "")):
            raise ValidationError("at least one of ['Artist','Title','Album','Content'] must be specified")



