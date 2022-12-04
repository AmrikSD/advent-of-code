def get_elf_sections(line):
    range_1, range_2 = line.split(",")

    elf_1_1, elf_1_2 = [int(value) for value in range_1.split("-")]
    elf_2_1, elf_2_2 = [int(value) for value in range_2.split("-")]

    return elf_1_1, elf_1_2, elf_2_1, elf_2_2
