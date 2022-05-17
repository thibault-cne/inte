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
    query = "SELECT * FROM Stars WHERE moderation_status = 1 ORDER BY date DESC"
    cursor.execute(query)

    # Get the result
    stars = cursor.fetchall()

    # Close the database
    close_database(db, cursor)

    # Change the format 
    stars = [starMeilleurFormat(star) for star in stars]
    
    # Return the result
    return stars


# Function to get all stars of a user
def get_all_stars_of_user(user_id: int) -> list:
    """
    Get all stars of a user
    :param user_id: the user id
    :return: a list of dicts
    """
    # Connect to the database
    db, cursor = connect_to_database()

    # Get all stars of a user
    query = "SELECT * FROM Stars WHERE receiver_user_id = ? AND moderation_status = 1 ORDER BY date DESC"
    cursor.execute(query, (user_id,))

    # Get the result
    stars = cursor.fetchall()

    # Close the database
    close_database(db, cursor)

    # Return the result
    return stars


def starMeilleurFormat(star):
    star1 = {}
    star1["id"] = star[0]
    star1["giver_user_id"] = star[1]
    star1["receiver_user_id"] = star[2]
    star1["star_rank"] = "d'or" if star[3] == 0 else "d'argent" if star[3] == 1 else "de bronze"
    star1["message"] = star[4]
    star1["date"] = star[5]
    star1["moderation_status"] = star[6]
    return star1