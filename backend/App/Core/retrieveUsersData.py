#  Copyright (c)
#  Project : backend
#
#  --
#
#  Author : thibault
#  File : retrieveUsersData.py
#  Description : *Enter description here*
#
#  --
#
#  Last modification : 2022/5/16

# Import needed packages


# Import personal packages
from App.Database.GetFunctions.getAllStars import get_all_stars_of_user
from App.Database.GetFunctions.getProfileInfo import get_user_name, get_user_points


# Create a function to retrieve all users information and format it to be used in the profile api
def retrieve_users_data(user_id: int) -> dict:
    """
    Get users info
    :param user_id:
    :return: dict
    """

    # Get user stars
    stars = get_all_stars_of_user(user_id)

    # Get username
    user_name = get_user_name(user_id)

    # Get user's points
    user_points = get_user_points(user_id)

    # Format data to be returned
    data = {
        'stars': stars,
        'user_name': user_name,
        'user_points': user_points
    }

    return data
