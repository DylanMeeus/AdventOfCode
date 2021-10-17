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


console.log(solve1());
