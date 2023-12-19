from marshmallow import Schema, fields, validates_schema, ValidationError


# Schéma utilisateur de sortie (renvoyé au front)
class SongSchema(Schema):
    id = fields.String(description="UUID")
    inscription_date = fields.DateTime(description="Inscription date")
    name = fields.String(description="Name")
    username = fields.String(description="Username")
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("name") or obj.get("name") == "") and \
               (not obj.get("username") or obj.get("username") == "") and \
               (not obj.get("inscription_date") or obj.get("inscription_date") == "")


class BaseRatingSchema(Schema):
    score = fields.String(description="score")
    content = fields.String(description="content")

class UpdateRatingSchema(Schema):
    score = fields.String(description="score")
    content = fields.String(description="content")
    idSong = fields.String(description="idSong")
    idUser = fields.String(description="idUser")

class SetRatingSchema(Schema):
    score = fields.String(description="score")
    content = fields.String(description="content")
    idSong = fields.String(description="idSong")


# Schéma utilisateur de modification (score, content)
class SongUpdateSchema(BaseRatingSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("score" in data and data["score"] != "") or
                ("content" in data and data["content"] != "")):
            raise ValidationError("at least one of ['score','content'] must be specified")

