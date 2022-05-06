#  Copyright (c)
#  Project : backend
#
#  --
#
#  Author : thibault
#  File : __init__.py
#  Description : *Enter description here*
#
#  --
#
#  Last modification : 2022/5/6

# Import needed packages
from sqlite3 import connect, Connection, Cursor
from typing import Tuple


# Import personal packages



def connect_to_database() -> Tuple[Connection, Cursor]:
    """Connect to the database.
    :return: the database connection and the cursor
    """
    db = connect('database.db')
    cursor = db.cursor()

    return db, cursor


def close_database(db: Connection, cursor: Cursor) -> None:
    """Close the database.
    :param db: the database connection
    :param cursor: the cursor
    """
    cursor.close()
    db.close()



