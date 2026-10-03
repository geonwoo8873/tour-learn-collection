import os

while True:
    print(f"""
    'r', 'read' = Read the contents of a file and return it as a string. 
    'a', 'append' = Append to the end of a file.
    'w', 'write' = Write to a file, overwriting its contents.
    'x', 'create' = Create a new file and write to it. Fails if the file already exists.
    'b', 'binary' = Binary mode. Used in conjunction with other modes (e.g., 'rb', 'wb').
    't', 'text' = Text mode. Default mode for reading and writing text files.
    
    {os.listdir(
        os.getcwd()
    )}
    """)
    
    read_file = input("Enter the file operation (r, a, w, x, b, t): ")

    match read_file:
        case "r", "read":
            read_file = input("Enter the file path: ")
            with open(read_file, 'r') as file:
                content = file.read()
            print(content)
        case "a", "append":
            with open(read_file, 'a') as file:
                content = input("Enter the content to append: ")
                file.write(content)
        case "w", "write":
            with open(read_file, 'w') as file:
                content = input("Enter the content to write: ")
                file.write(content)
        case "x", "create":
            with open(read_file, 'x') as file:
                if os.path.exists(read_file):
                    print("File already exists.")
                else:
                    content = input("Enter the content to write: ")
                    file.write(content)
        case "b", "binary":
            read_file = input("Enter the file path: ")
            with open(read_file, 'rb') as file:
                content = file.read()
            print(content)
        case "t", "text":
            read_file = input("Enter the file path: ")
            with open(read_file, 'rt') as file:
                content = file.read()
            print(content)
        case _:
            print("Invalid option")