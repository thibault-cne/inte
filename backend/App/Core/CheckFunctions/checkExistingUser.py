#  Copyright (c)
#  Project : backend
#
#  --
#
#  Author : thibault
#  File : checkExistingUser.py
#  Description : *Enter description here*
#
#  --
#
#  Last modification : 2022/5/7

# Import needed packages


# Import personal packages
from ...Database import connect_to_database, close_database


# Function to check if a user exists in the database
def check_existing_user(email: str) -> bool:
    # Connect to the database
    database, database_cursor = connect_to_database()

    # Query to check if the user exists
    query = """
        SELECT * FROM Users WHERE email = ?;
    """
    database_cursor.execute(query, (email,))
    user = database_cursor.fetchone()

    # Close the database
    close_database(database, database_cursor)

    # Return the result
    if user is None:
        return False
    else:
        return True
