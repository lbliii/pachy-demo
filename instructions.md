## 1. Create a Repo 

Use the following command to build your repo:
    `pachctl create repo lb-demo`

**TIP**: Use the command `pachctl list repo` to see all of your Pachyderm repositories. 

## 2. Upload Data (Log Files)

1. Find or create example log files that include warning and error messages. 
2. Use the following command to commit a log file to the `lb-demo` repo: 
    `pachctl put file lb-demo@master:log1.txt -f /Users/lblane/Documents/pachy/logs/log1.txt`
3. Verify the log file was added:
    `pachctl list file lb-demo@master`
4. View
    `pachctl get file lb-demo@master:log1.txt | open -f -a TextEdit.app` 

**INFO**: Pachyderm commits are similar to git commits; `{reponame}@{branch}:filename.extension -f /Path/of/file.txt` 

## 3. Create Executable 

Pachyderm is language agnostic, so you can build images however you like. For this exercise, we need to create an executable file that traverses a log directory and counts all instances of warning and error messages; after that, it must output those results to a `results.txt` file.

### Go 

1. Create a Go file (e.g., `count.go`).
2. Define the package (e.g., `package main`). 
3.  // Go mod file init 
4. Create the following global vars:
    - **Int**: `errorCount` 
    - **Int**: `warningCount`
5. Create a function that gets a list of log files located in the `pfs/{repo-name}/logs` directory. 
6. Iterate through each file and parse the content into individual lines.
7. Check each line for variations of `WARN`, `WARNING`, `ERR`, `ERROR`. 
8. Add++ to each global variable for each occurrence.
9. Output a `results.txt` file in the `pfs/out` directory that contains the warning and error totals. 

## 4. Build Docker Image 

### Go 

[Official Documentation](https://docs.docker.com/language/golang/build-images/)

`docker build -t lbliii/lb-demo:1.0 .` 

## 4. Create a Pipeline 

`pachctl create pipeline -f /Users/lblane/Documents/GitHub/pachy-demo/lb-pipeline.json`
