# operating system

In order to develop applications on your computer, you will need to learn the basic concepts and commands of operating systems.

## Basics
An operating system is an application that is the main program running on your computer. It is responsible for bascially everything. It communicates with all devices such as monitor, mouse, keyboard, speakers, network, and storage. It runs all your programs (processes) and determines how much resources (cpu & memory) should be given to each. It keep track of users and their permissions.

The core of an operating system is the _kernal_. Its job is take files (programs) and run them (processes). It also has a set of low-level commands called _syscalls_ (system calls). Programs can call a command to request for memory, start new threads, and write to files. The kernal is not directly responsible for rendering the screen. It is instead in charge of running the program that renders the screen.

## Terminal
The terminal (or shell) is the space where the user can input instructions for the operating system. It can do everything that can be with a mouse and more.

Definitions:
- `$` or `%`: start of user input
    - The character depends on the shell profile that's current running
    - A line in the terminal that does not start with `$` means it's an output from the program
    - _NOTE: `$` is also used to access variables in shell. See examples below_  

    ```sh
        My-Computer:~ sbkim$ # bash/windows

        sbkim@My-Computer ~ % # zsh/macOS
    ```
- `#`: comment
    - every character after `#` that in the same line is not run by the shell. It is as if those characters weren't typed in
- `.`: current directory
- `..`: parent directory
- `/`: root directory
- `~`: home directory (default directory)
- `directory`: folder
- `absolute path`: path to an object starting from root (`/`)
- `relative path`: path to an object relative to current directory (implicit start with `./`)
- `Up/Down Arrow`: scroll thru history of past commands (easy way rerun commands)
- `environment variables`: variables that are set within terminal that is accessed by the programs
- `\`: continue command into the next line
- `standard input (stdin)`: default input for programs (underlying file)
- `standard output (stdout)`: default ouput for programs (underlying file)
    
    ```sh
        $ echo "hello" || \
        > echo "world"
        hello
        world
    ```

Commands:

`<ARGUMENT>`: reqired argument
`[ARGUMENT]`: optional agrument

- `pwd`: absolute path to working (current) directory

    ```sh
        $ pwd
        /Users/sbkim/Desktop/pool-io/learn
    ```
    
- `ls`: list items in the specified directory (default: `.`)
    - It is an error to call this on a file

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

    ```sh
        $ cd # go to home directory
        $ cd ~ # same as above
        $ cd .. # go to parent directory
        $ cd / # go to root directory
    ```

- `mkdir`: create directory

    ```sh
        $ mkdir example
        $ ls
        example
        ...

        $ mkdir ~/example # relative to home
        $ ls ~
        example
        ...
    ```

- `touch`: create file

    ```sh
        $ touch example.txt
        $ ls
        example.txt

        $ touch /Users/sbkim/Desktop/hello
        $ ls
        ...
        $ ls /Users/sbkim/Desktop
        hello
    ```

- `export <NAME>=<VALUE>`: set environment variable
    - It is 

    ```sh
        $ export 
    ```

- `echo <CONTENT>`: print
    ```sh
        $ echo "hello world"
        hello world

        # hello
    ```

- `echo "<CONTENT>" > <FILEPATH>`: write to file
    - Changes the `stdout` of echo `process` to the file

    ```sh
        $ touch example.txt
        $ echo "hello world" > example.txt
    ```

- `echo "<CONTENT>" >> <FILEPATH>`: append to file

    ```sh
        $ touch example.tx
    ```

- `cat <FILEPATH>`: print file
    
    ```sh
        $ touch example.txt
        $ echo "hello" > example.txt
        $ cat example.txt
        hello world
    ```

-  `grep <PHRASE>`: print line that contain phrase

    ```sh
        $ echo "hello world\nbye world" | grep "bye"
        bye world
    ```

- `cp <SOURCE_PATH> <DEST_PATH>`: copy

    ```sh
        $ touch example.txt && echo "hello world" > example.txt
        $ cat example.txt
        hello world
        $ ls
        example.txt
        ...
        $ cp example.txt copy.txt
        $ ls
        copy.txt
        example.txt
        ...
        $ cat copy.txt
        hello world
    ```

- `mv`: move

    ```sh
        $ touch example.txt && echo "hello world" > example.txt
        $ cat 
        $ ls
        example.txt
        ...
        $ mv example.txt moved.txt
        $ ls
        moved.txt
        ...
        $ cat moved.txt
        hello world
    ```

- `rm`: remove
    - Use flag `-rf` to remove directories
    _Beware: Running `rm -rf /` will delete everything in your computer (including the os)_
    
    ```sh
        $ touch example.txt && echo "hello world" > example.txt
        $ cp example.txt copy.txt
        $ ls
        copy.txt
        example.txt
        ...
        $ rm example.txt
        $ ls
        copy.txt
        ...
    ```

- `./<executable>`: run program

    ```sh
        $ ls
        helloworld
        $ ./helloworld
        hello world
    ```

- `ps`: list processes
    - Use `aux` argument to get processes from all `users`

    ```sh
        $ ps
        PID TTY           TIME CMD
        7267 ttys000    0:00.08 /bin/zsh -il
        8278 ttys000    0:25.20 node /usr/local/bin/yarn run dev
        8279 ttys000    0:38.34 /usr/local/bin/node /Users/fwaygo/Desktop/pool/website/node_modules/.bin/next dev
        7455 ttys001    0:00.29 /bin/zsh -il
        17755 ttys002    0:00.36 /bin/zsh -il
        21504 ttys003    0:00.01 /bin/bash --init-file /Applications/Visual Studio Code.app/Contents/Resources/app/out/vs/workbench/contrib/terminal

        $ ps | grep node
        8278 ttys000    0:25.25 node /usr/local/bin/yarn run dev
        8279 ttys000    0:38.38 /usr/local/bin/node /Users/fwaygo/Desktop/pool/website/node_modules/.bin/next dev
        27933 ttys002    0:00.00 grep node
    ```

## Excerise
1. Open Terminal
    - Use the one connected to vscode
        - For Windows, choose `Git Bash` profile

2. Try out all examples above to get a sense on how it works