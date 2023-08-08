package main

import (
    "fmt" 
    "os"
    "bufio"
)

func current_command(s string, ip int) (rune, int) {
    if ip >= len(s) - 1 || ip < 0 {
        return -1, ip;
    }
    char := rune(s[ip]);
    if char == '\n' {
        return current_command(s, ip + 1);
    }

    return char, ip;
}

func main() {
    source, err := os.ReadFile("hello_world.bf");
    if err != nil {
        panic("could not read file");
    }

    reader := bufio.NewReader(os.Stdin);

    source_str := string(source);
    fmt.Println(source_str);
    ip := 0;
    dp := 0;
    var tape [30_000]byte;

    char, ip := current_command(source_str, ip);
    for char != -1 {
        switch char {
        case '>':
            dp++;
        case '<':
            dp--;
        case '+':
            tape[dp]++;
        case '-':
            tape[dp]--;
        case '.':
            fmt.Printf("%c", tape[dp]);
        case ',':
            fmt.Print("> ");
            char, _, err = reader.ReadRune();
            if err != nil {
                panic("could not read input");
            }
            tape[dp] = byte(char);
        case '[':
            if tape[dp] == 0 {
                for char != -1 {
                    char, ip = current_command(source_str, ip + 1);
                    if char == ']' {
                        break;
                    }
                }
            }
        case ']':
            if tape[dp] != 0 {
                for char != -1 {
                    char, ip = current_command(source_str, ip - 1);
                    if char == '[' {
                        ip--;
                        break;
                    }
                }
            }

        default:
            panic("unrecognized char: " + string(char));
        }

        char, ip = current_command(source_str, ip + 1);
    }
}
