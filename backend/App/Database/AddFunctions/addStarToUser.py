#  Copyright (c)
#  Project : backend
#
#  --
#
#  Author : thibault
#  File : addStarToUser.py
#  Description : *Enter description here*
#
#  --
#
#  Last modification : 2022/5/10

# Import needed packages
from datetime import datetime


# Import personal packages
from .addPointsToUser import add_points_to_user
from .. import connect_to_database, close_database


STARS_POINTS = {
    'gold': 10,
    'silver': 7,
    'bronze': 5
}


def add_star_to_user(user_id: int, giver_id: int, message: str, star_type: str) -> None:
    """
    Add a star to a user
    :param user_id: the user's id
    :param giver_id: the giver's id
    :param message: the message
    :param star_type: the star type
    :return:
    """

    # Connect to the database
    db, cursor = connect_to_database()

    # Current date to format dd/mm/yyyy
    now = datetime.now()
    date = now.strftime("%d/%m/%Y")

    # Add the star
    query = """
        INSERT INTO stars (receiver_user_id, giver_user_id, message, star_rank, date) VALUES (?, ?, ?, ?, ?);
    """
    args = (user_id, giver_id, message, star_type, date)

    # Execute the query
    cursor.execute(query, args)
    db.commit()
    close_database(db, cursor)

    # Add the points to the user
    add_points_to_user(user_id, STARS_POINTS[star_type])
