class mapper:
    def __init__(self, seed_to_soil, soil_to_fert, fert_to_water, water_to_light, light_to_temp, temp_to_hum, hum_to_loc):
        self.seed_to_soil = sorted(seed_to_soil)
        self.soil_to_fert = sorted(soil_to_fert)
        self.fert_to_water = sorted(fert_to_water)
        self.water_to_light = sorted(water_to_light)
        self.light_to_temp = sorted(light_to_temp)
        self.temp_to_hum = sorted(temp_to_hum)
        self.hum_to_loc = sorted(hum_to_loc)

        self.order = [self.seed_to_soil, self.soil_to_fert, self.fert_to_water, self.water_to_light, self.light_to_temp, self.temp_to_hum, self.hum_to_loc]

        self.reverse = list(reversed(self.order))



    ## binary lookup of destination to source
    def binary_lookup(self, m, key):

        m = sorted(m, key=lambda k: k.destination_start)
        l = 0
        r = len(m)-1

        while l <= r:
            mid = l + (r-l) // 2

            if m[mid].contains_destination(key):
                return m[mid].map_to_source(key)
            
            # key is greater than source start
            elif key < m[mid].destination_start:
                r = mid - 1
            else:
                l = l + 1

        return key

    def lookup(self, m, key):
        for _entry in m:
            if _entry.contains_source(key):
                return _entry.map_to_destination(key)
            elif key < _entry.source_start:
                return key
        return key


    def loc_to_seed(self, location):
        x = location
        for m in self.reverse:
            x = self.binary_lookup(m,x)
        return x

    def seed_to_loc(self, seed):
        x = seed
        for m in self.order:
            x = self.lookup(m, x)
        return x


class seed_entry:
    def __init__(self, seed_start, seed_range):
        self.seed_start = seed_start
        self.seed_range = seed_range

    def contains(self, seed):
        if seed >= self.seed_start and seed < self.seed_start + self.seed_range:
            return True
        return False
        

def solve2(lines):
    seeds, mapper = parse(lines)

    locations = sorted(mapper.hum_to_loc, key=lambda x: x.destination_start)

    i = 0

    seed_entries = []
    while i < len(seeds):
        start = seeds[i]
        length = seeds[i+1]
        seed_entries.append(seed_entry(start,length))
        i += 2

    seed_entries = sorted(seed_entries, key=lambda k: k.seed_start)


    for x in range(int(10e6),int(10e12)):
        seed = mapper.loc_to_seed(x)
        if x % 1000 == 0:
            print(x)
        if seed != -1:
            ## now we check if this seed exists
            for se in seed_entries:
                if se.contains(seed):
                    return x




    return -1

def solve1(lines):
    seeds, mapper = parse(lines)

    print('solving')


    lowest = 10e9
    for seed in seeds:
        loc = mapper.seed_to_loc(seed)
        if loc < lowest:
            lowest = loc


    return lowest


class entry:
    def __init__(self, destination_start, source_start, length):
        self.destination_start = destination_start
        self.source_start = source_start
        self.length = length

    def contains_source(self, x):
        if x >= self.source_start and x < self.source_start + self.length:
            return True
        return False

    def contains_destination(self, x):
        if x >= self.destination_start and x < self.destination_start + self.length:
            return True
        return False

    def map_to_source(self, destination):
        if not self.contains_destination(destination):
            exit("fubar")
        offset = destination - self.destination_start
        return self.source_start + offset


    def map_to_destination(self, x):
        if not self.contains_source(x):
            exit("fubar")
        offset = x - self.source_start
        return self.destination_start + offset

    def __lt__(self, other):
        return self.source_start < other.source_start

    def __str__(self):
        return f'{self.source_start} -> {self.destination_start}'



def parse(lines):
    seeds = list(map(lambda k: int(k), lines[0][7:].split(' ')))

    seed_to_soil = []
    soil_to_fert = []
    fert_to_water = []
    water_to_light = []
    light_to_temp = []
    temp_to_hum = [] 
    hum_to_loc = []

    txt_to_map = {
            'seed-to-soil': seed_to_soil,
            'soil-to-fertilizer': soil_to_fert,
            'fertilizer-to-water': fert_to_water,
            'water-to-light': water_to_light,
            'light-to-temperature': light_to_temp,
            'temperature-to-humidity': temp_to_hum,
            'humidity-to-location': hum_to_loc
            }


    expand = lambda x, r: [i for i in range(x, x+r)]


    last_seen = 'seed-to-soil'
    for line in lines[1:]:
        if line == '':
            continue

        if 'map' in line:
            last_seen = line.split(" ")[0]
            continue

        parts = line.split(' ')
        end, start, ranges = int(parts[0]), int(parts[1]), int(parts[2])
        e = entry(end,start,ranges)
        m = txt_to_map[last_seen]
        m.append(e)


    return (seeds, mapper(seed_to_soil, soil_to_fert, fert_to_water, water_to_light, light_to_temp, temp_to_hum, hum_to_loc))





if __name__ == '__main__':
    lines = open('input.txt').read().split("\n")
    #print(solve1(lines))
    print(solve2(lines))
