# operating system

In order to develop applications on your computer, you will need to learn the basic concepts and commands.

## Basics
An operating system is an application that is the main program running on your computer. It is responsible for bascially everything. It communicates with all devices such as monitor, mouse, keyboard, speakers, network, and storage. It runs all your programs (processes) and determines how much resources (cpu & memory) should be given to each. It keep track of users and their permissions.

The core of an operating system is the _kernal_. Its job is take files (programs) and run them (processes). It also has a set of low-level commands called _syscalls_ (system calls). Programs can call a command to request for memory, start new threads, and write to files. The kernal is not directly responsible for rendering the screen. It is instead in charge of running the program that renders the screen.

## Terminal
The terminal (or shell) is the space where the user can input instructions for the operating system. It can do everything that can be with a mouse and more.

Definitions:
    `$` or `%`: start of user input
        - The character depends on the shell profile that's current running
        - A line in the terminal that does not start with `$` means it's an output from the program
    `#`: comment
        - every character after `#` that in the same line is not run by the shell. It is as if those characters weren't typed in
    `.`: current directory
    `..`: parent directory
    `/`: root directory
    `directory`: folder
    `absolute path`: path to an object starting from root (`/`)
    `relative path`: path to an object relative to current directory (implicit start with `./`)
    `\`: continue command into the next line
        Ex.
            ```sh
                $ 
            ```

Commands:
    - `pwd`: absolute path to working (current) directory
        Ex.
            ```sh
                $ pwd
                /Users/sbkim/Desktop/pool-io/learn
            ```
    - `ls`: list items in the specified directory (default: `.`)
        - It is an error to call this on a file
        Ex.
            ```sh
                $ ls
                0-start
                1-os
                ...

                $ ls . # equal to above
                0-start
                1-os
                ...

                $ ls /Users/sbkim/Desktop # absolute path
                learn
                ...

                $ ls 0-start
            ```
    - `cd`: change directory


## Excerise
1. Open Terminal
    - Use the one connected to vscode
    - For Windows, choose `Git Bash` profile

2. 