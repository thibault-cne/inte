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
import os


# Import personal packages



def connect_to_database() -> Tuple[Connection, Cursor]:
    """Connect to the database.
    :return: the database connection and the cursor
    """
    path = os.path.dirname(os.path.abspath(__file__))
    dbPath = os.path.join(path, 'database.db')

    database = connect(dbPath)
    database_cursor = database.cursor()

    return database, database_cursor


def close_database(database: Connection, database_cursor: Cursor) -> None:
    """Close the database.
    :param database: the database connection
    :param database_cursor: the cursor
    """
    database_cursor.close()
    database.close()


def create_database() -> str:
    creation_query = """
        DROP TABLE IF EXISTS Users;
        DROP TABLE IF EXISTS Stars;
        DROP TABLE IF EXISTS TNder;
        DROP TABLE IF EXISTS Challenges;
        DROP TABLE IF EXISTS Calendar;
        DROP TABLE IF EXISTS DailyGame;
        DROP TABLE IF EXISTS Notifications;
        DROP TABLE IF EXISTS Suggestions;
        
        
        CREATE TABLE Users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            email TEXT NOT NULL,
            username TEXT NOT NULL,
            current_year INTEGER,
            promotion_year INTEGER,
            points INTEGER,
            godfather_id INTEGER,
            facebook TEXT,
            snapchat TEXT,
            instagram TEXT,
            google TEXT,
            hometown TEXT,
            studies TEXT,
            personal_message TEXT,
            last_login TEXT,
            color TEXT,
            FOREIGN KEY (godfather_id) REFERENCES Users(id)
        );
        
        CREATE TABLE Stars (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            giver_user_id INTEGER NOT NULL,
            receiver_user_id INTEGER NOT NULL,
            star_rank INTEGER NOT NULL,
            message TEXT NOT NULL,
            date TEXT NOT NULL,
            FOREIGN KEY (giver_user_id) REFERENCES Users(id),
            FOREIGN KEY (receiver_user_id) REFERENCES Users(id)
        );
        
        CREATE TABLE TNder (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            matching_user_id INTEGER NOT NULL,
            matched_user_id INTEGER NOT NULL,
            FOREIGN KEY (matching_user_id) REFERENCES Users(id),
            FOREIGN KEY (matched_user_id) REFERENCES Users(id)
        );
        
        CREATE TABLE Challenges (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            number INTEGER NOT NULL,
            title TEXT NOT NULL,
            description TEXT NOT NULL,
            winner_user_id INTEGER,
            begin_date TEXT NOT NULL,
            end_date TEXT NOT NULL,
            FOREIGN KEY (winner_user_id) REFERENCES Users(id)
        );
        
        CREATE TABLE Calendar (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            date TEXT NOT NULL,
            title TEXT NOT NULL,
            description TEXT NOT NULL
        );
        
        CREATE TABLE DailyGame (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            date TEXT NOT NULL,
            result INTEGER NOT NULL,
            FOREIGN KEY (user_id) REFERENCES Users(id)
        );
        
        CREATE TABLE Notifications (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            type TEXT NOT NULL,
            message TEXT NOT NULL,
            FOREIGN KEY (user_id) REFERENCES Users(id)
        );
        
        CREATE TABLE Suggestions (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            title TEXT NOT NULL,
            description TEXT NOT NULL,
            FOREIGN KEY (user_id) REFERENCES Users(id)
        );
    """

    return creation_query


if __name__ == '__main__':
    query = create_database()
    db, cursor = connect_to_database()

    cursor.executescript(query)
    db.commit()
    close_database(db, cursor)
