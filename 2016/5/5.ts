import {Md5} from 'ts-md5/dist/md5';


function solve1(): string {
    const input = "wtnhxymk"

    let password: string = "";
    
    let i = 0;
    while (password.length != 8) {

        let foundHash = false;
        while (!foundHash) {
            const candidate = input + i;
            const hashed = Md5.hashStr(candidate);
            if ( hashed.startsWith("00000") ){
                const char = hashed.toString()[5];
                foundHash = true;
                password += char;
            }
            i++;
        }

    }

    return password;
}


function solve2(): string {
    const input = "wtnhxymk"

    let password: string[] = new Array(8);
    
    let i = 0;
    let foundChars = 0;
    while (foundChars != 8) {
        let foundHash = false;
        while (!foundHash) {
            const candidate = input + i;
            i++;
            const hashed = Md5.hashStr(candidate);
            if ( hashed.startsWith("00000") ){
                if ("01234567".indexOf(hashed.toString()[5]) < 0) {
                    continue;
                }
                const  idx = + hashed.toString()[5];
                if (password[idx] !== undefined || idx > 7) {
                    continue;
                }
                const char = hashed.toString()[6];
                foundHash = true;
                foundChars++;
                password[idx] = char;
                console.log(password.join(''));
            }
        }

    }

    return password.join('');
}

console.log(solve2());
