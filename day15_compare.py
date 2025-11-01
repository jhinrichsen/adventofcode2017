#!/usr/bin/env python3
import time

def day15_part1(start_a, start_b):
    FACTOR_A = 16807
    FACTOR_B = 48271
    MODULO = 2147483647
    PAIRS = 40000000

    a, b = start_a, start_b
    matches = 0

    for _ in range(PAIRS):
        a = (a * FACTOR_A) % MODULO
        b = (b * FACTOR_B) % MODULO
        if (a & 0xFFFF) == (b & 0xFFFF):
            matches += 1

    return matches

def day15_part2(start_a, start_b):
    FACTOR_A = 16807
    FACTOR_B = 48271
    MODULO = 2147483647
    PAIRS = 5000000

    a, b = start_a, start_b
    matches = 0

    for _ in range(PAIRS):
        while True:
            a = (a * FACTOR_A) % MODULO
            if a % 4 == 0:
                break
        while True:
            b = (b * FACTOR_B) % MODULO
            if b % 8 == 0:
                break
        if (a & 0xFFFF) == (b & 0xFFFF):
            matches += 1

    return matches

if __name__ == "__main__":
    # Example input
    start_a = 65
    start_b = 8921

    print("Python implementation benchmark:")
    print()

    # Part 1
    start = time.time()
    result1 = day15_part1(start_a, start_b)
    elapsed1 = time.time() - start
    print(f"Part 1: {result1} matches")
    print(f"Time: {elapsed1:.3f}s ({elapsed1*1000:.0f}ms)")
    print()

    # Part 2
    start = time.time()
    result2 = day15_part2(start_a, start_b)
    elapsed2 = time.time() - start
    print(f"Part 2: {result2} matches")
    print(f"Time: {elapsed2:.3f}s ({elapsed2*1000:.0f}ms)")
    print()

    print(f"Go is ~{elapsed1/0.109:.1f}x faster than Python for Part 1")
    print(f"Go is ~{elapsed2/0.196:.1f}x faster than Python for Part 2")
