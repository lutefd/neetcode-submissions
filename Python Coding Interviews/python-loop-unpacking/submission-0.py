from typing import List, Tuple


def best_student(scores: List[Tuple[str, int]]) -> str:
    result = ("", 0)
    for name, score in scores:
        if score > result[1]:
            result = (name, score)
    return result[0]


# do not modify below this line
print(best_student([("Alice", 90), ("Bob", 80), ("Charlie", 70)]))
print(best_student([("Alice", 90), ("Bob", 80), ("Charlie", 100)]))
print(best_student([("Alice", 90), ("Bob", 100), ("Charlie", 70)]))
print(best_student([("Alice", 90), ("Bob", 90), ("Charlie", 80), ("David", 100)]))
