from src.helpers import db

class Song(db.Model):
    __tablename__ = 'song'

    id = db.Column(db.String(255), primary_key=True)
    artist = db.Column(db.String(255), nullable=False)
    title = db.Column(db.String(255), nullable=False)
    album = db.Column(db.String(255), nullable=False)
    content = db.Column(db.String(255), nullable=False)

    def __init__(self, song_id, artist, title, album, content):
        self.id = song_id
        self.artist = artist
        self.title = title
        self.album = album
        self.content = content

    def is_empty(self):
        return (not self.id or self.id == "") and \
               (not self.artist or self.artist == "") and \
               (not self.title or self.title == "") and \
               (not self.album or self.album == "") and \
               (not self.content or self.content == "")
