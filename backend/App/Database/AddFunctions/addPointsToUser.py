#  Copyright (c)
#  Project : backend
#
#  --
#
#  Author : thibault
#  File : addPointsToUser.py
#  Description : *Enter description here*
#
#  --
#
#  Last modification : 2022/5/10

# Import needed packages


# Import personal packages
from .. import connect_to_database, close_database


def add_points_to_user(user_id: int, points: int) -> None:
    """
    Add points to a user
    :param user_id: the user's id
    :param points: the points to add
    :return:
    """

    # Connect to the database
    db, cursor = connect_to_database()

    # Add points to the user
    query = """
        UPDATE users SET points = points + ? WHERE id = ?
    """
    args = (points, user_id)

    # Execute the query
    cursor.execute(query, args)
    db.commit()
    close_database(db, cursor)
