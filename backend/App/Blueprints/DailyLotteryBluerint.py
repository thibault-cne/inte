#  Copyright (c)
#  Project : backend
#
#  --
#
#  Author : thibault
#  File : DailyLotteryBluerint.py
#  Description : *Enter description here*
#
#  --
#
#  Last modification : 2022/5/15

# Import needed packages
from flask import Blueprint, request, jsonify
from flask_jwt_extended import jwt_required, get_jwt_identity


# Import personal packages
from App.Core.CheckFunctions.checkDailyLottery import check_daily_lottery


# Create the blueprint
daily_lottery_blueprint = Blueprint('daily_lottery', __name__)


@daily_lottery_blueprint.route('/check_daily_lottery', methods=['GET'])
@jwt_required
def check_daily_lottery():
    """
    Check if the user has already played the daily lottery
    :return:
    """
    if request.method == 'GET':
        user_id = get_jwt_identity()

        # Check if the user has already played the daily lottery
        if check_daily_lottery(user_id):
            payload = {
                'message': 'You have already played the daily lottery',
                'status': 'success',
                'data': {
                    'has_played': True
                }
            }
        else:
            payload = {
                'message': 'You have not played the daily lottery yet',
                'status': 'success',
                'data': {
                    'has_played': False
                }
            }

        return jsonify(payload), 200
