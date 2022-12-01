def get_calories(file_iter):
    elves = {}
    index = 0

    for line in file_iter:
        if line == "\n":
            index += 1
        else:
            if index not in elves:
                elves[index] = int(line)
            else:
                elves[index] += int(line)

    return elves
