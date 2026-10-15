from typing import List

def read_integers() -> List[int]:
    r_string = input("")
    string_list = r_string.split(",")
    num_list = []
    for i in string_list:
        num_list.append(int(i))
    return num_list
# do not modify the code below
print(read_integers())
print(read_integers())
print(read_integers())
