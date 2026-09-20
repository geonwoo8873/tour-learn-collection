import mysql.connector

conn = mysql.connector.connect(
    host="localhost",
    user="yourusername",
    password="yourpassword",
    database="yourdatabase"
)

cursor = conn.cursor()

if conn.is_connected():
    print("Connected to the database")
else:
    print("Failed to connect to the database")
    conn.close()