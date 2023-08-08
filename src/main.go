package main

import (
    "fmt" 
    "os"
    "bufio"
)

func current_command(s string, ip *int) rune {
    //fmt.Printf("CMD:\tip: %d\tchar: %c\n", *ip, rune(s[*ip]));
    if *ip >= len(s) - 1 { // newline
        return -1;
    }
    
    char := rune(s[*ip]);
    if char == '\n' || char == ' ' {
        *ip++;
        char = current_command(s, ip);
    }
    return char;
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

    char := current_command(source_str, &ip);
    for char != -1 {
        // fmt.Printf("LOOP:\tip: %d\tDP: %d\ttape[DP]: %d\tchar: %c\n", ip, dp, tape[dp], char);
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
            char, _, err = reader.ReadRune();
            if err == nil {
                panic("could not read input");
            }
            tape[dp] = byte(char);
        case '[':
            if tape[dp] == 0 {
                for char != -1 {
                    ip += 1;
                    char = current_command(source_str, &ip);
                    if char == ']' {
                        break;
                    }
                }
            }
        case ']':
            if tape[dp] != 0 {
                for char != -1 {
                    ip -= 1;
                    char = current_command(source_str, &ip);
                    if char == '[' {
                        ip -= 1;
                        break;
                    }
                }
            }

        default:
            panic("unrecognized char: " + string(char));
        }

        ip += 1;
        char = current_command(source_str, &ip);
    }
}
