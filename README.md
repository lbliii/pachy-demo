# How to Create a Log Processing Pipeline 

The following guide creates a [Pachyderm](https://www.pachyderm.com/) pipeline that processes the number of WARNING and ERROR messages that occur in a log directory's files. You can use this pipeline to analyze system behavior, discover performance trends, and monitor your logs. 

## Before You Start 

Make sure all of the following tools are installed on your machine:

- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [Minikube](https://minikube.sigs.k8s.io/docs/start/)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Virtualbox](https://www.virtualbox.org/wiki/Downloads)
- [Helm](https://helm.sh/docs/intro/install/)
- [Pachctl](https://docs.pachyderm.com/latest/getting-started/local-installation/#install-pachctl)

## 1. Create a Repository 

[Pachyderm repositories](https://docs.pachyderm.com/latest/concepts/data-concepts/repo/#repository) are a top-level data object that accept all file types. Similar to Git, you can make changes to your repository by pushing [commits](https://docs.pachyderm.com/latest/concepts/data-concepts/commit/#commit) containing files and folders. We will use this repository to send new log files to get processed.

1. Open a terminal. 
2. Use the following command to build your repo: `pachctl create repo lb-pachy-project`

**💡 TIP**: You can use the command `pachctl list repo` to see all of your Pachyderm repositories. 

```
NAME              CREATED        SIZE (MASTER) DESCRIPTION                           
lb-pachy-project           31 hours ago   ≤ 292.3KiB  
```

## 2. Upload Data (Log Files)

Uploading data to your repository requires use of the [pachctl put file](https://docs.pachyderm.com/latest/reference/pachctl/pachctl_put_file) command. Using this command, you can put **files**, **images**, **data**, or whole **directories** into your repository.

1. Find or create example log files that include warning and error messages. 
2. Use the following command to commit a log file to the `lb-pachy-project` repo: 
   ```
   pachctl put file lb-pachy-project@master:log1.txt -f /Users/lblane/Documents/pachy/logs/log1.txt
   ```
3. Verify the log file was added: `pachctl list file lb-pachy-project@master`.
   ```
   NAME      TYPE SIZE     
   /log1.txt file 292.3KiB 
   ```
4. Optionally, you can view the file: 
   ```
   pachctl get file lb-pachy-project@master:log1.txt | open -f -a TextEdit.app
   ```
5. Repeat for as many files or directories necessary. 

**✏️ NOTE**: Pachyderm commits are similar to git commits. Here's a quick breakdown of the format:
   ```
   {repoName}@{branch}:{fileName}.{extension} -f /Path/of/file.txt
   ```

**💡 TIP**: Have a directory of log files? You can use `pachctl put file -r repo@branch -f {dirName}` to upload it with one command. 

## 3. Create a Pipeline 

✅ We've got our repository set up. 

✅ We've committed log data.

Now, let's create a [pipeline](https://docs.pachyderm.com/latest/concepts/pipeline-concepts/pipeline/#pipeline). A pipeline reads, transforms, and outputs data. To use a pipeline, you must define a pipeline schema (either in `JSON` or `YAML`). Pipelines also create repos using the pipeline's name.

**💡 TIP**: Use a prefix, such as `p-`, in the pipeline's name to visually distinguish pipeline-dependent output repos versus input repos when using the `pachctl list repo` command.

   ```
   lblane@lb-lane ~ % pachctl list repo                                                             
   NAME               CREATED           SIZE (MASTER) DESCRIPTION                                  
   p-lb-pachy-project About an hour ago ≤ 29B         Output repo for pipeline p-lb-pachy-project. 
   lb-pachy-project   About an hour ago ≤ 24.84KiB  
   ```

### Define Pipeline Schema 

#### JSON Version

```json 
{
    "pipeline": {
      "name": "p-lb-pachy-project" // Displayed when using the following command: pachctl list pipeline
    },
    "description": "A pipeline that counts WARNING and ERROR occurrences in one or many log files.", // Displayed when using the following command: pachctl list pipeline
    "transform": {
      "cmd": [ "./count" ], // The command that executes the data transformation & output
      "image": "lbliii/lb-pachy-project:5.0" // The Docker image containing the scripts/logic needed to transform the data. 
    },
    "input": {
      "pfs": { // The Pachyderm file system
        "repo": "lb-pachy-project", // The repository name 
        "glob": "/*" // A global pattern used to return all matching files; example: pachctl glob file <repo>@<branch-or-commit>:<pattern> [flags]
      }
    }
  }
```

#### YAML Version

```yaml
pipeline:
  name: p-lb-pachy-project
description: A pipeline that counts WARNING and ERROR occurrences in one or many log files.
transform:
  cmd:
    - ./count 
  image: lbliii/lb-pachy-project:5.0
input:
  pfs:
    repo: lb-pachy-project
    glob: /*
```

#### About Count.go

The following code block contains all of the details of our executable `count.go` file:  

```go 
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var (
	errorCount   int
	warningCount int
)

func main() {
	traverseLogs()
}

func traverseLogs() {

	files, err := ioutil.ReadDir("/pfs/lb-pachy-project")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".txt") {
			content, err := ioutil.ReadFile("/pfs/lb-pachy-project/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			countWarningsAndErrors(string(content))
		}
	}
}

func countWarningsAndErrors(content string) {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.Contains(line, "ERR") {
			errorCount++
		} else if strings.Contains(line, "WARN") {
			warningCount++
		}
	}
	createResultsFile(errorCount, warningCount)
}

func createResultsFile(errorCount int, warningCount int) {
	results := "errorCount: " + fmt.Sprint(errorCount) + "\n" + "warningCount: " + fmt.Sprint(warningCount)
	file := ioutil.WriteFile("/pfs/out/results.txt", []byte(results), 0644)
	if file != nil {
		log.Fatal(file)
	}

}
```
**💡 TIP**: You can also check out examples in [Javascript](https://github.com/lbliii/pachy-demo/blob/main/count.js) and [Python](https://github.com/lbliii/pachy-demo/blob/main/count.py). 

### Submit Pipeline Schema to Pachyderm 

Submitting a pipeline schema to Pachyderm requires using the [pachctl create pipeline](https://docs.pachyderm.com/latest/reference/pachctl/pachctl_create_pipeline/) command. Using this command, you can push `JSON`, `YAML`, `Jsonnet`, and local Docker images to Pachyderm.

1. Open a terminal.
2. Run the following command: 
   ``` 
   pachctl create pipeline -f https://raw.githubusercontent.com/lbliii/pachy-demo/main/lb-pipeline.json
   ```

Once submitted, the pipeline automatically runs a [job](https://docs.pachyderm.com/latest/concepts/pipeline-concepts/job/#job) using the code that transforms and outputs your data.  

**✏️ NOTE**: You can use the command `pachctl list job` to see a list of jobs. 

## Check the Output 

You can download and view the output of the `results.txt` file using the following command: 

```
pachctl get file lb-pachy-project@master:results.txt | open -f -a TextEdit.app
```