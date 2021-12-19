from datetime import datetime


def read_file(f):
    """Reads file"""
    with open(f, 'r') as file:
        return file.read().strip()


def timer(func):
    """Prints result and timer for result"""
    def wrapper(*args):
        start = datetime.now()
        f = func(*args)
        end = datetime.now()
        print(f)
        print(end - start)
        print()
        return f
    return wrapper