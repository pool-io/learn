# git

[git]() is a version control system that is widely used by programmers. It keeps track of all changes made and enables you view its history. `git` is _initialized_ in a directory by running `git init` which creates a `.git` directory within it. This directory is already initialized since it is a part of the `/learn` directory which has the `.git` directory. You can check the existance of `.git` by running `ls -la`

![git](https://miro.medium.com/max/1400/1*69prApOy8_cZPnmRHGHQZg.png)

## Commits
A `commit` is a snapshot of the directory in a point in time. It is the the basic building blocks to a directory's history in git. In addition to the holding the state of all files in the directory at that point in time, it also keeps track of the author, timestamp, parent commit, and user defined message.

In order to commit any saved changes, that file much be added to `staging`. This intermediated space is needed such that changes across mulitple files can be added to the same commit. A entired directory can also be added which would result in all the files also being added.

```sh
    $ ls
    example.txt
    $ cat example.txt
    hello world
    $ git status
    On branch master
    Your branch is up to date with 'origin/master'.

    nothing to commit, working tree clean
    $
    $ echo "bye world" > example.txt
    $ git add example.txt
    $ git status
    On branch master
    Your branch is up to date with 'origin/master'.

    Changes to be committed:
    (use "git restore --staged <file>..." to unstage)
            modified:  example.txt
    $
    $ git commit -m "change hello to bye"
    $ git status
    On branch master
    Your branch is up to date with 'origin/master'.

    nothing to commit, working tree clean
```

To see your git history: `git log`

To checkout a specific commit: `git checkout <COMMIT_HASH>`

To get back `git checkout <NAME_OF_BRANCH>`

2. Branches
A `branch` references a lineage of commits. However, it only points to a single commit, referred to as the `head`, which is then followed by going to each commit's parent.

![git branches](https://the-turing-way.netlify.app/_images/sub-branch.png)

To create a branch: `git checkout -b <NAME_OF_BRANCH>`

Omit the `-b` flag if the branch already exists: `git checkout <NAME_OF_BRANCH>`

3. Remote