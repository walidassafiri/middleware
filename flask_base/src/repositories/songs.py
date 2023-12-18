from src.helpers import db
from src.models.user import User


def get_song(username):
    return db.session.query(User).filter(User.username == username).first()

