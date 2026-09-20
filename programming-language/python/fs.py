""" 
'r' = Read the contents of a file and return it as a string. 
'a' = Append to the end of a file.
'w' = Write to a file, overwriting its contents.
'x' = Create a new file and write to it. Fails if the file already exists.
'b' = Binary mode. Used in conjunction with other modes (e.g., 'rb', 'wb').
't' = Text mode. Default mode for reading and writing text files.
"""

# open(<file path>, <fs mode>) as file:
with open("print.py", 'r') as file:
    content = file.read()
print(content)