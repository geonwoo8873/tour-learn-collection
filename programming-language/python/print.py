# Default encoding for this Python file utf-8 -*-
# Other encodings can be specified if needed cp1252, latin-1, etc.
# -*- coding: <encoding name> -*-

######## Print Output Examples ########
# Single-line string output using print function
print("Default encoding for this Python file is utf-8")

# `"""..."""`multi-line string output using triple quotes
print("""
    Test Python print outputting from terminal
""")

# `\n` adds a new line before the text outputting to the terminal
print("\nTest")

# `(f"...")` is used for formatted string literals, allowing expressions to be embedded inside string output
name = "Alice"
age = 30
print(f"My name is {name} and I am {age} years old.")

# Using expressions inside f-strings
print(f"In five years, I will be {age + 5} years old.")

# Using expressions with more complex calculations inside f-strings
height = 1.75  # in meters
weight = 68  # in kilograms
print(f"My BMI is {weight / (height ** 2):.2f}")

# Using conditional expressions inside f-strings
is_adult = age >= 18
print(f"Is {name} an adult? {'Yes' if is_adult else 'No'}")