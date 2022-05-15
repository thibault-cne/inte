#  Copyright (c)
#  Project : backend
#
#  --
#
#  Author : thibault
#  File : getAllStars.py
#  Description : *Enter description here*
#
#  --
#
#  Last modification : 2022/5/15

# Import needed packages


# Import personal packages
from .. import connect_to_database, close_database


# Create a function to retrieve all stars info ordered by date
def get_all_stars() -> list:
    """
    Get all stars info ordered by date
    :return: a list of dicts
    """
    # Connect to the database
    db, cursor = connect_to_database()

    # Get all stars info ordered by date
    query = "SELECT * FROM Stars ORDER BY date DESC"
    cursor.execute(query)

    # Get the result
    stars = cursor.fetchall()

    # Close the database
    close_database(db, cursor)

    # Return the result
    return stars
