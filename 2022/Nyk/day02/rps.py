winning_states = {
    "X": {"winning_state": "C", "score": 1, "equal_state": "A"},
    "Y": {"winning_state": "A", "score": 2, "equal_state": "B"},
    "Z": {"winning_state": "B", "score": 3, "equal_state": "C"},
}

existing_state = {
    "A": {"states": {"Z": "B", "Y": "A", "X": "C"}, "score": 1},
    "B": {"states": {"Z": "C", "Y": "B", "X": "A"}, "score": 2},
    "C": {"states": {"Z": "A", "Y": "C", "X": "B"}, "score": 3},
}
