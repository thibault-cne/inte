#  Copyright (c)
#  Project : backend
#
#  --
#
#  Author : thibault
#  File : getProfileInfo.py
#  Description : *Enter description here*
#
#  --
#
#  Last modification : 2022/5/17

# Import needed packages


# Import personal packages
from .. import connect_to_database, close_database


def get_user_name(user_id: int) -> str:
    """
    Get user name
    :param user_id:
    :return:
    """

    # Connect to database
    db, cursor = connect_to_database()

    # Get username
    user_info = cursor.execute(
        "SELECT username FROM users WHERE id = :user_id",
        {'user_id': user_id}
    ).fetchone()

    # Close database
    close_database(db, cursor)

    # Return username
    return user_info[0]


def get_user_points(user_id: int) -> int:
    """
    Get user points
    :param user_id:
    :return:
    """

    # Connect to database
    db, cursor = connect_to_database()

    # Get user points
    user_points = cursor.execute(
        "SELECT points FROM users WHERE id = :user_id",
        {'user_id': user_id}
    ).fetchone()

    # Close database
    close_database(db, cursor)

    # Return user points
    return user_points[0]
