



class mapper:
    def __init__(self, seed_to_soil, soil_to_fert, fert_to_water, water_to_light, light_to_temp, temp_to_hum, hum_to_loc):
        self.seed_to_soil = seed_to_soil 
        self.soil_to_fert = soil_to_fert 
        self.fert_to_water = fert_to_water
        self.water_to_light = water_to_light 
        self.light_to_temp = light_to_temp 
        self.temp_to_hum = temp_to_hum 
        self.hum_to_loc = hum_to_loc 

        self.order = [self.seed_to_soil, self.soil_to_fert, self.fert_to_water, self.water_to_light, self.light_to_temp, self.temp_to_hum, self.hum_to_loc]

    def lookup(self, m, key):
        for _entry in m:
            if _entry.contains_source(key):
                return _entry.map_to_destination(key)
        return key

    def seed_to_loc(self, seed):
        x = seed
        for m in self.order:
            x = self.lookup(m, x)
        return x
        


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
        if x >= self.source_start and x <= self.source_start + self.length:
            return True

    def map_to_destination(self, x):
        if not self.contains_source(x):
            exit("fubar")
        offset = x - self.source_start
        return self.destination_start + offset



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
            print(last_seen)
            continue

        parts = line.split(' ')
        end, start, ranges = int(parts[0]), int(parts[1]), int(parts[2])
        e = entry(end,start,ranges)
        m = txt_to_map[last_seen]
        m.append(e)


    return (seeds, mapper(seed_to_soil, soil_to_fert, fert_to_water, water_to_light, light_to_temp, temp_to_hum, hum_to_loc))





if __name__ == '__main__':
    lines = open('input.txt').read().split("\n")
    print(solve1(lines))
