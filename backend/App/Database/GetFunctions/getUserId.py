#  Copyright (c)
#  Project : backend
#
#  --
#
#  Author : thibault
#  File : getUserId.py
#  Description : *Enter description here*
#
#  --
#
#  Last modification : 2022/5/7

# Import needed packages


# Import personal packages
from .. import connect_to_database, close_database


# Function to get the user's id
def get_user_id(email: str) -> int:
    # Connect to the database
    database, database_cursor = connect_to_database()

    # Query to get the user's id
    query = """
        SELECT id FROM Users WHERE email = ?;
    """
    database_cursor.execute(query, (email,))
    user_id = database_cursor.fetchone()

    # Close the database
    close_database(database, database_cursor)

    # Return the user's id
    return user_id[0]
