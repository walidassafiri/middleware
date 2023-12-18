from src.helpers import db

class Song(db.Model):
    __tablename__ = 'song'

    id = db.Column(db.String(255), primary_key=True)
    Artist = db.Column(db.String(255), nullable=False)
    Title = db.Column(db.String(255), nullable=False)
    Album = db.Column(db.String(255), nullable=False)
    Content = db.Column(db.String(255), nullable=False)

    def __init__(self, song_id, Artist, Title, Album, Content):
        self.id = song_id
        self.Artist = Artist
        self.Title = Title
        self.Album = Album
        self.Content = Content

    def is_empty(self):
        return (not self.id or self.id == "") and \
               (not self.Artist or self.Artist == "") and \
               (not self.Title or self.Title == "") and \
               (not self.Album or self.Album == "") and \
               (not self.Content or self.Content == "")
