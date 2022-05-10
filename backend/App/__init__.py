# Import needed packages
from flask import Flask
from flask_cors import CORS
from flask_jwt_extended import JWTManager
from datetime import timedelta

# Import personal packages


# Definition of the app
def create_app() -> Flask:
    """
    Create the flask app
    :return: the flask app
    """
    app = Flask(__name__)

    app.config['SECRET_KEY'] = 'jOQiU7q9I2l4eycgKM3hh6oTy7CEaX6j'

    # Config upload folder
    app.config['UPLOAD_FOLDER'] = "./Data/ProfilPictures/"

    # JWT Configuration
    app.config["JWT_SECRET_KEY"] = "VSfa4AP9UHqdM00NWXs0N3KgxoAeOZXK"
    app.config["JWT_HEADER_TYPE"] = "Bearer"
    app.config["JWT_TOKEN_LOCATION"] = "headers"
    app.config["JWT_ACCESS_TOKEN_EXPIRES"] = timedelta(hours=1)
    app.config["JWT_REFRESH_TOKEN_EXPIRES"] = timedelta(days=30)

    # Import blueprints
    from App.Blueprints.AuthBlueprint import auth_blueprint
    app.register_blueprint(auth_blueprint)

    # Handle JWT with the app to enable authentication
    JWTManager(app)

    # Handle CORS with the app
    CORS(app, ressources={r'/*': {'origins': '*'}})

    return app
