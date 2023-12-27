from src.helpers import db
from src.models.song import Song

def get_songs():
    songs = Song.query.all()
    songs_list = [{'id': song.id, 'artist': song.artist, 'title': song.title, 'album': song.album, 'content': song.content} for song in songs]
    return jsonify(songs_list)


def get_song(song_id):
    song = Song.query.get_or_404(song_id)
    return jsonify({'id': song.id, 'artist': song.artist, 'title': song.title, 'album': song.album, 'content': song.content})



def create_song(song_model):
    """
    Ajoute une nouvelle chanson à la base de données.
    """
    db.session.add(song_model)
    try:
        # Tentative de commit des changements
        db.session.commit()
        print("Chanson ajoutée avec succès à la base de données.")
    except Exception as e:
        # En cas d'erreur, annuler les modifications
        db.session.rollback()
        print(f"Erreur lors de l'ajout de la chanson : {e}")



def update_song(updated_song_model):
    """
    Met à jour une chanson dans la base de données.
    """
    song = get_song_from_id(updated_song_model.id)
    if song:
        song.artist = updated_song_model.artist
        song.title = updated_song_model.title
        song.album = updated_song_model.album
        song.content = updated_song_model.content
        db.session.commit()
    else:
        raise NotFound("Chanson non trouvée")




def delete_song(song_id):
    song = Song.query.get_or_404(song_id)
    db.session.delete(song)
    db.session.commit()