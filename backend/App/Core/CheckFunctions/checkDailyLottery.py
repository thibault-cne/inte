#  Copyright (c)
#  Project : backend
#
#  --
#
#  Author : thibault
#  File : checkDailyLottery.py
#  Description : *Enter description here*
#
#  --
#
#  Last modification : 2022/5/15

# Import needed packages
import datetime

# Import personal packages
from App.Database import connect_to_database, close_database


def check_daily_lottery(user_id: int) -> bool:
    """
    Check if the user has already played the daily lottery
    :param user_id: the user id
    :return:
    """
    # Connect to the database
    db, cursor = connect_to_database()

    # Get the today's date to format dd/mm/yyyy
    today = datetime.datetime.today()
    today_date = today.strftime('%d/%m/%Y')

    # Check if the user has already played the daily lottery
    query = """SELECT * FROM DailyGame WHERE user_id = ? AND date = ?"""
    cursor.execute(query, (user_id, today_date))

    # If the user has already played the daily lottery, return True
    if cursor.fetchone():
        close_database(db, cursor)
        return True
    else:
        close_database(db, cursor)
        return False
