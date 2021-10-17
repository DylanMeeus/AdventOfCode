import * as fs from 'fs';


type Room = {
    name: string;
    sector: number;
    checksum: string;
}


function calculateChecksum(input: string): string {
    const counts: Map<string, number> = new Map<string, number>();
    for (const str of input.split('')) {
        if (counts.get(str) !== undefined ) {
            counts.set(str, counts.get(str)! + 1); // compiler is being stupid, so I have to add `!` after map access.
        } else {
            counts.set(str, 1);
        }
    }

    const sortedMap: Map<string, number> = new Map([...counts.entries()].sort( (e1,e2) => {
            if (e2[1] == e1[1]) {
                return e2[0] < e1[0] ? 1 : -1;
            }
            return e2[1] - e1[1];
    }));


    return Array.from(sortedMap.keys()).slice(0,5).join('');
}



function convertInputToStruct(lines: string[]): Room[] {
        // aaaaa-bbb-z-y-x-123[abxyz]
        // turn into 3 parts like below 
        // { name }       {id} {checksum}
        const mapToRoom = (input: string): Room => {
            const parts = input.split("[");
            const nameAndSector = parts[0];
            const checkSum: string = parts[1].replace("]","");
            const nameSectorSplit: string[] = nameAndSector.split("-");
            const sectorID = nameSectorSplit[nameSectorSplit.length-1];
            const sanitizedName = nameSectorSplit.slice(0,nameSectorSplit.length - 1).join('');

            return {
                name: sanitizedName,
                sector: + sectorID,
                checksum: checkSum,
            }
        };

        return lines.map(mapToRoom);
}

function solve1() {
    fs.readFile('input.txt', 'utf8', (error, data) => {
        if (error != null) {
            console.log("could not read file, sorry!");
            return;
        }
        const lines: string[] = data.split("\n").filter(l => l != "");

        const validRooms = convertInputToStruct(lines).filter(room => calculateChecksum(room.name) == room.checksum);
        const result: number = validRooms.map((r: Room) => r.sector)
                                             .reduce((prev: number, curr: number) => prev + curr);
        console.log(result);
    });
}


// caesarCipher shifts each character in `str` by `rot`
function caesarCipher(str: string, rot: number): string {
    const normalized_rot = rot % 26;

    const char_codes = str.split('').map(char => char.charCodeAt(0) - 97); // 97 == 'a'

    const result = char_codes.map(code => (code + normalized_rot) % 26).
        map(rot_code => String.fromCharCode(rot_code + 97));

    return result.join('');
}

function solve2() {
    fs.readFile('input.txt', 'utf8', (error, data) => {
        if (error != null) {
            console.log("could not read file, sorry!");
            return;
        }
        const lines: string[] = data.split("\n").filter(l => l != "");
        const rooms = convertInputToStruct(lines);

        const deciphered = rooms.map(r => {
            return {
                name: caesarCipher(r.name, r.sector),
                sector: r.sector
            };
        });
        const result = deciphered.filter(r => r.name.indexOf('obj') > 0)
        console.log(result);
    });
}

//solve1();
solve2();
