from advent import read_file, timer
from collections import defaultdict
from itertools import permutations


def orientation():
    """Yields each of 8 possible orientations"""
    for i in range(8):
        yield (-1) ** (i % 2), (-1) ** ((i // 2) % 2), (-1) ** ((i // 4) % 2)


class Scanner:
    def __init__(self, num, pos=None):
        self.beacons = set()
        self.num = num
        self.pos = pos

    def add_beacon(self, beacon):
        self.beacons.add(beacon)

    def compare(self, other):
        # Possible positions from self
        c = dict()

        # For each orientation and permutation (?? 48 possible orientations ??)
        for orient in orientation():
            for perm in permutations((0, 1, 2)):
                c[orient, perm] = defaultdict(int)

                # Compare each beacon from each scanner
                for a in self.beacons:
                    for b in other.beacons:
                        # Assume a is absolute, find other.pos
                        c[orient, perm][tuple(da+o*b[dp] for da, o, dp in zip(a, orient, perm))] += 1

                # If 12 beacons have the same comparison, k is the absolute position of the other scanner
                for k, v in c[orient, perm].items():
                    if v >= 12:
                        other.pos = k

                        # Reorient all beacons based on absolute position
                        other.reorient(orient, perm)
                        return other.pos

    def __str__(self):
        return '--- scanner ' + str(self.num) + ' ---\nPosition: ' \
               + str(self.pos) + '\n' + '\n'.join(str(b) for b in self.beacons)

    def __repr__(self):
        return str(self)

    def reorient(self, orient, perm):
        """Reorient the scanner by the permutation and orientation"""
        new_s = set()
        for b in self.beacons:
            new_b = tuple(self.pos[order] - o*b[p] for o, p, order in zip(orient, perm, (0, 1, 2)))
            new_s.add(new_b)

        self.beacons = new_s


def prep_scanners(ins):
    s = []

    # Read each scanner in
    for i, scanner in enumerate(f):
        s.append(Scanner(i))
        for line in scanner.split('\n')[1:]:
            s[i].add_beacon(tuple(int(i) for i in line.split(',')))

    # Give scanner 0 the origin position
    known = {0}
    s[0].pos = (0, 0, 0)

    # Keep comparing scanners until all absolute positions are known
    while set(range(len(s))) - known != set():
        new = set()
        for k in known:
            for i in range(len(s)):
                if i in known:
                    continue
                if s[k].compare(s[i]):
                    new.add(i)
        known |= new
        print("Total Mapped:", len(known))

    return s


def manhattan(a, b):
    return sum(abs(da - db) for da, db in zip(a, b))


@timer
def p1(s):
    beacons = set.union(*(scan.beacons for scan in s))
    return len(beacons)

f = read_file("2021_19.txt").split('\n\n')
s = prep_scanners(f)
print('\n')
p1(s)
