#  Copyright (c)
#  Project : backend
#
#  --
#
#  Author : thibault
#  File : StarsBlueprint.py
#  Description : *Enter description here*
#
#  --
#
#  Last modification : 2022/5/15

# Import needed packages
from flask import Blueprint, request, jsonify, Response
from flask_jwt_extended import jwt_required, get_jwt_identity
from typing import Tuple


# Import personal packages
from App.Database.GetFunctions.getAllStars import get_all_stars

# Create the blueprint
stars_blueprint = Blueprint('stars', __name__)


# Create a function to retrieve all stars info ordered by date
@stars_blueprint.route('/stars-info/<method>', methods=['GET'])
@jwt_required(optional=True)
def get_all_stars_info(method) -> Tuple[Response, int]:
    """
    Get all stars info ordered by date
    :return: a list of dicts
    """
    if request.method == 'GET':
        # Check if the user has already played the daily lottery
        if method == 'all':
            stars = get_all_stars()
        else:
            return jsonify({'message': 'Method not found', 'status': 'error'}), 404

        # Return the result
        return jsonify(stars), 200
