# Solution to day 20 of AoC



class Vector:
    def __init__(self,a,b,c):
        self.a = int(a)
        self.b = int(b)
        self.c = int(c)

    def __eq__(self, other):
        return isinstance(other, Vector) and (self.a == other.a and self.b == other.b and self.c == other.c)
    
    def __hash__(self):
        return hash((self.a,self.b,self.c))

    def __str__(self):
        return "<" + str(self.a) + ", " + str(self.b) + ", " + str(self.c) + ">"

class Particle:
    def __init__(self, position, velocity, acceleration):
        self.position = position
        self.velocity = velocity
        self.acceleration = acceleration


def getInput():
    f = open('input20.txt','r')
    lines = f.read()
    return lines.split('\n')[:-1]


def create_particles(data):
    particles = []
    for line in data:
        parts = line.split(' ')
        location_string = parts[0][3:-2]
        velocity_string = parts[1][3:-2]
        acceleration_string = parts[2][3:-1]
        location_parts = location_string.split(',')
        velocity_parts = velocity_string.split(',')
        acceleration_parts = acceleration_string.split(',')
        location = Vector(location_parts[0],location_parts[1],location_parts[2])
        velocity = Vector(velocity_parts[0],velocity_parts[1],velocity_parts[2])
        acceleration = Vector(acceleration_parts[0],acceleration_parts[1],acceleration_parts[2])
        particle = Particle(location,velocity,acceleration)
        particles.append(particle)
    return particles

def solve():
    inp = getInput()
    data = create_particles(inp)
    # Calculate the distance for each particle. First determine if it is going further away or
    # closer

    index = 0
    results = {}
    for particle in data:
        # First add acceleration to velocity, then velocity to location
        particle_closing = True # is the particle getting closer
        while particle_closing:
            p = particle.position
            a = particle.acceleration
            v = particle.velocity
            v.a += a.a
            v.b += a.b
            v.c += a.c
            particle.velocity = v
            old_distance = abs(p.a) + abs(p.b) + abs(p.c)
            p.a += v.a
            p.b += v.b
            p.c += v.c
            new_distance = abs(p.a) + abs(p.b) + abs(p.c)
            if new_distance > old_distance:
                results[index] = old_distance # The closest distance to 0
                particle_closing = False
        index += 1
    max_key = 0
    max_value = results[0] 
    print(results.keys())
    for key in results.keys():
        if results[key] < max_value:
            max_key = key
            max_value = results[key]
    print(max_key)



def filter_collisions(particles):
    positions = list(map(lambda k: k.position,particles))
    # Now count the occurances
    position_occurances = {}
    for position in positions:
        occurances = len(list(filter(lambda k: k == position, positions)))
        position_occurances[position] = occurances

    new_particles = []
    for particle in particles:
        if position_occurances[particle.position] == 1:
            new_particles.append(particle)
    return new_particles
    

def solve2():
    inp = getInput()
    data = create_particles(inp)
    # Calculate the distance for each particle. First determine if it is going further away or
    # closer


    for i in range(1000):
        for particle in data:
            # First add acceleration to velocity, then velocity to location
            p = particle.position
            a = particle.acceleration
            v = particle.velocity
            v.a += a.a
            v.b += a.b
            v.c += a.c
            particle.velocity = v
            p.a += v.a
            p.b += v.b
            p.c += v.c
        data = filter_collisions(data)
        print(len(data))

   
solve2()
