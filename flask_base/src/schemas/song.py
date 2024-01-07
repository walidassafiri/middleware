from marshmallow import Schema, fields, validates_schema, ValidationError


class BaseRatingSchema(Schema):
    score = fields.String(description="score")
    content = fields.String(description="content")

class UpdateRatingSchema(Schema):
    score = fields.String(description="score")
    content = fields.String(description="content")
    idSong = fields.String(description="idSong")
    idUser = fields.String(description="idUser")


# Schéma utilisateur de modification (score, content)
class SongUpdateSchema(BaseRatingSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("score" in data and data["score"] != "") or
                ("content" in data and data["content"] != "")):
            raise ValidationError("at least one of ['score','content'] must be specified")

