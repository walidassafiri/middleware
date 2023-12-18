from src.helpers import db
from src.models.song import Song

def get_songs():
    songs = Song.query.all()
    songs_list = [{'id': song.id, 'Artist': song.Artist, 'Title': song.Title, 'Album': song.Album, 'Content': song.Content} for song in songs]
    return jsonify(songs_list)


def get_song(song_id):
    song = Song.query.get_or_404(song_id)
    return jsonify({'id': song.id, 'Artist': song.Artist, 'Title': song.Title, 'Album': song.Album, 'Content': song.Content})



def create_song():
    data = request.get_json()
    new_song = Song(Artist=data['Artist'], Title=data['Title'], Album=data['Album'], Content=data['Content'])
    db.session.add(new_song)
    db.session.commit()




def update_song(song_id):
    song = Song.query.get_or_404(song_id)
    data = request.get_json()
    song.Artist = data['Artist']
    song.Title = data['Title']
    song.Album = data['Album']
    song.Content = data['Content']
    db.session.commit()




def delete_song(song_id):
    song = Song.query.get_or_404(song_id)
    db.session.delete(song)
    db.session.commit()