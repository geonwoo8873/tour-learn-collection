# Default encoding for this Python file utf-8 -*-
# Other encodings can be specified if needed cp1252, latin-1, etc.
# -*- coding: <encoding name> -*-

from typing import Final

######## Print Output Examples ########
# Single-line string output using print function
print("Default encoding for this Python file is utf-8")

# `"""..."""`multi-line string output using triple quotes
print("""
    Test Python print outputting from terminal
""")

# `\n` adds a new line before the text outputting to the terminal
print("\nTest")

######## Variables Declaration ########
# 
temperature: Final[int] = -42

update_temperature: Final[float] = temperature

target: int = -20

difference = target-temperature
print(difference)