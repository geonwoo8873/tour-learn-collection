import mysql.connector
# `dataclass` is used to create class that primarily store data
from dataclasses import dataclass

# `@dataclass` is used to automatic genrate special methods
@dataclass
class DatabaseConfig:
    host: str = "localhost"
    port: int = 3306
    user: str = "yourusername"
    password: str = "yourpassword"
    database: str = "yourdatabase"

_config = DatabaseConfig()

conn = mysql.connector.connect(
    host=_config.host,
    port=_config.port,
    user=_config.user,
    password=_config.password,
    database=_config.database
)

cursor = conn.cursor()

if conn.is_connected():
    print("Connected to the database")
else:
    print("Failed to connect to the database")
    conn.close()