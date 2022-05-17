#  Copyright (c)
#  Project : backend
#
#  --
#
#  Author : thibault
#  File : ProfileBlueprint.py
#  Description : *Enter description here*
#
#  --
#
#  Last modification : 2022/5/16

# Import needed packages
from flask import Blueprint, request, jsonify, Response
from flask_jwt_extended import jwt_required, get_jwt_identity
from typing import Tuple


# Import personal packages
from App.Core.retrieveUsersData import retrieve_users_data


# Define the blueprint
profile_blueprint = Blueprint('profile', __name__)


# Route to retrieve all users information
@profile_blueprint.route('/profile-api/<method>', methods=['GET'])
@jwt_required()
def get_profile_info(method: str) -> Tuple[Response, int]:
    """
    Get users info
    :param method:
    :return:
    """

    if request.method == 'GET':
        user_id = get_jwt_identity()
        if method == 'all':
            # Get user info
            data = retrieve_users_data(user_id)


            payload = {
                'status': 'success',
                'data': data
            }

            return jsonify(payload), 200

        else:
            return jsonify({'message': 'Method not found', 'status': 'error'}), 404

    else:
        return jsonify({'message': 'Method not allowed', 'status': 'error'}), 405
