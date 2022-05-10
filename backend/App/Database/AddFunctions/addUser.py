#  Copyright (c)
#  Project : backend
#
#  --
#
#  Author : thibault
#  File : addUser.py
#  Description : *Enter description here*
#
#  --
#
#  Last modification : 2022/5/7

# Import needed packages


# Import personal packages
from .. import connect_to_database, close_database
from ..GetFunctions.getUserId import get_user_id
from ...Core.CheckFunctions.checkExistingUser import check_existing_user


# Function to register a user in the database and return the user's id
def add_user(email: str, username: str) -> int:
    # check if the user already exists
    if check_existing_user(email):
        return get_user_id(email)

    # Connect to the database
    database, database_cursor = connect_to_database()

    # Query to add the user
    query = """
        INSERT INTO Users (email, username) VALUES (?, ?);
    """
    database_cursor.execute(query, (email, username))
    database.commit()

    # Close the database
    close_database(database, database_cursor)

    # Return the user's id
    return get_user_id(email)
