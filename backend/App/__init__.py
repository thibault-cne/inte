# Import needed packages
from flask import Flask
from flask_cors import CORS
from flask_jwt_extended import JWTManager
from datetime import timedelta
from dotenv import load_dotenv
from os import getenv

# Import personal packages


# Load the dotenv file
load_dotenv()
APP_SECRET_KEY = getenv('APP_SECRET_KEY')
JWT_SECRET_KEY = getenv('JWT_SECRET_KEY')


# Definition of the app
def create_app() -> Flask:
    """
    Create the flask app
    :return: the flask app
    """
    app = Flask(__name__)

    app.config['SECRET_KEY'] = APP_SECRET_KEY

    # JWT Configuration
    app.config["JWT_SECRET_KEY"] = JWT_SECRET_KEY
    app.config["JWT_HEADER_TYPE"] = "Bearer"
    app.config["JWT_TOKEN_LOCATION"] = "headers"
    app.config["JWT_ACCESS_TOKEN_EXPIRES"] = timedelta(hours=1)
    app.config["JWT_REFRESH_TOKEN_EXPIRES"] = timedelta(days=30)

    # Import blueprints
    from App.Blueprints.AuthBlueprint import auth_blueprint
    app.register_blueprint(auth_blueprint)

    from App.Blueprints.StarsBlueprint import stars_blueprint
    app.register_blueprint(stars_blueprint)

    from App.Blueprints.DailyLotteryBluerint import daily_lottery_blueprint
    app.register_blueprint(daily_lottery_blueprint)

    # Handle JWT with the app to enable authentication
    JWTManager(app)

    # Handle CORS with the app
    CORS(app, ressources={r'/*': {'origins': '*'}})

    return app
