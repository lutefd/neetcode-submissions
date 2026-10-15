def add_two_numbers() -> int:
    u_input = input("")
    strings = u_input.split(",")
    ints = []
    for s in strings:
        ints.append(int(s))
    return sum(ints)



# do not modify below this line
print(add_two_numbers())
print(add_two_numbers())
print(add_two_numbers())
print(add_two_numbers())
