from collections import *
from itertools import *
from functools import *

import numpy as np


def sq_dist(a, b):
    return np.sum((a - b) ** 2)

def all_triangles(beacon_list):
    result = set()
    triangle_to_points = {}
    for triple in combinations(list(range(len(beacon_list))), 3):
        dists = tuple(sorted(np.sum((beacon_list[a] - beacon_list[b])**2) for a, b in combinations(triple, 2)))
        result.add(dists)
        triangle_to_points[dists] = triple
    return result, triangle_to_points


def diffs_from_triple(triple):
    for p1, p2 in pairwise(islice(cycle(triple), 4)):
        yield p1 - p2

def order(points, triangles):
    for triple in permutations(points):
        if all(d == np.sum(diff ** 2) for d, diff in zip(triangles, diffs_from_triple(triple))):
            return triple


A = np.array([])
def iter_aligned_triangles(first, first_map, second, second_map, overlaps):
    for triangle in overlaps:
        if len(set(triangle)) != 3:
            continue

        first_points = order([first[i] for i in first_map[triangle]], triangle)
        second_points = order([second[i] for i in second_map[triangle]], triangle)
        yield first_points, second_points


def match(first, first_map, second, second_map, overlaps):
    coord_perm = np.full(3, -1, dtype=np.int32)
    for first_points, second_points in iter_aligned_triangles(first, first_map, second, second_map, overlaps):
        for diff1, diff2 in zip(diffs_from_triple(first_points), diffs_from_triple(second_points)):
            for i, d in enumerate(diff1):
                if np.sum(diff1 == d) > 1:
                    continue
                coord_perm[i] = np.argmax(np.abs(diff2) == np.abs(d))

        if np.all(coord_perm >= 0):
            break
    else:
        return None

    transformed = [p[coord_perm] for p in second]
    flip = np.zeros(3, np.int32)
    for first_points, second_points in iter_aligned_triangles(first, first_map, transformed, second_map, overlaps):
        for diff1, diff2 in zip(diffs_from_triple(first_points), diffs_from_triple(second_points)):
            assert np.all(np.abs(diff1) == np.abs(diff2))
            for i, (d1, d2) in enumerate(zip(diff1, diff2)):
                if not d1:
                    continue
                flip[i] = np.sign(d1) * np.sign(d2)

        if np.all(flip):
            break
    else:
            return None

    for first_points, second_points in iter_aligned_triangles(first, first_map, transformed, second_map, overlaps):
        shift = second_points[0] * flip - first_points[0]
        break;
    return [p * flip - shift for p in transformed], shift


def main():
    fo = open("2021_19.txt", "r+")
    lines = fo.readlines()

    beacons = []

    curr = None
    for line in lines:
        if line.startswith('---'):
            if curr:
                beacons.append(curr)
            curr = []
        elif line.strip():
            curr.append(np.array(list(map(int, line.split(',')))))
    beacons.append(curr)


    triangles = [all_triangles(s) for s in beacons]

    positions = [np.zeros(3, dtype=np.int32)]
    req_match = 220
    stack = [0]
    shifts = {0: np.zeros(3, dtype=np.int32)}
    while stack:
        i = stack.pop()
        reference_triangles, reference_map = triangles[i]
        for j, (other_triangles, other_map) in enumerate(triangles):
            if j in shifts:
                continue
            matching_triangles = reference_triangles & other_triangles
            if len(matching_triangles) >= req_match:
                print(i, j)
                transformed, shift = match(beacons[i], reference_map, beacons[j], other_map, matching_triangles)
                if transformed:
                    shifts[j] = shift
                    stack.append(j)
                    beacons[j] = transformed
    print(len(set(tuple(p) for scanner_list in beacons for p in scanner_list)))
    print(max(np.sum(np.abs(s1 - s2)) for s1, s2 in combinations(shifts.values(), 2)))

def pairwise(iterable):
    "s -> (s0,s1), (s1,s2), (s2, s3), ..."
    a, b = tee(iterable)
    next(b, None)
    return zip(a, b)


if __name__ == '__main__':
    main()