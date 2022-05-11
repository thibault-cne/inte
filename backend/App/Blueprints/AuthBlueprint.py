#  Copyright (c)
#  Project : backend
#
#  --
#
#  Author : thibault
#  File : AuthBlueprint.py
#  Description : *Enter description here*
#
#  --
#
#  Last modification : 2022/5/6

# Import needed packages
from typing import Union

from flask import Blueprint, request, jsonify, Response
from google.oauth2 import id_token
from google.auth.transport import requests
from dotenv import load_dotenv
from os import getenv
from flask_jwt_extended import create_access_token, create_refresh_token


# Import personal packages
from App.Database.AddFunctions.addUser import add_user


# Create the blueprint
auth_blueprint = Blueprint('auth', __name__)

# Load the dotenv file
load_dotenv()
CLIENT_ID = getenv('CLIENT_ID')


@auth_blueprint.route('/auth-api/<string:method>', methods=['GET', 'POST'])
def auth(method: str) -> Union[Response, str]:
    if request.method == 'POST':
        if method == 'login':
            # Create the return payload
            return_payload = {'status': 'success'}

            # Get the token from the request
            token = request.json['credential']
            try:
                # Specify the CLIENT_ID of the app that accesses the backend:
                id_info = id_token.verify_oauth2_token(token, requests.Request(), CLIENT_ID)

                # Or, if multiple clients access the backend server:
                # id_info = id_token.verify_oauth2_token(token, requests.Request())
                # if id_info['aud'] not in [CLIENT_ID_1, CLIENT_ID_2, CLIENT_ID_3]:
                #     raise ValueError('Could not verify audience.')

                # If auth request is from a G Suite domain:
                # if id_info['hd'] != GSUITE_DOMAIN_NAME:
                #     raise ValueError('Wrong hosted domain.')

                # ID token is valid. Get the user's Google Account ID from the decoded token.
                userid = id_info['sub']
                username = id_info['given_name'] + ' ' + id_info['family_name']
                email = id_info['email']

                # Add the user in the db if he doesn't exist
                # we also get the user's id
                user_id = add_user(email, username)

                # Create the access token
                access_token = create_access_token(identity=user_id)
                refresh_token = create_refresh_token(identity=user_id)

                # Add the access token to the return payload
                credentials = {'access_token': access_token, 'refresh_token': refresh_token}

                # Add the credentials to the return payload
                return_payload['credentials'] = credentials
            except ValueError:
                # Invalid token
                return_payload['status'] = 'error'
            return jsonify(return_payload)

    if request.method == 'GET':
        if method == 'login':
            return 'login'
