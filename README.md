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
2. Use the following command to build your repo: `pachctl create repo lb-demo`

**💡 TIP**: You can use the command `pachctl list repo` to see all of your Pachyderm repositories. 

```
NAME              CREATED        SIZE (MASTER) DESCRIPTION                           
lb-demo           31 hours ago   ≤ 292.3KiB  
```

## 2. Upload Data (Log Files)

Uploading data to your repository requires use of the [pachctl put file](https://docs.pachyderm.com/latest/reference/pachctl/pachctl_put_file) command. Using this command, you can put **files**, **images**, **data**, or whole **directories** into your repository.

1. Find or create example log files that include warning and error messages. 
2. Use the following command to commit a log file to the `lb-demo` repo: `pachctl put file lb-demo@master:log1.txt -f /Users/lblane/Documents/pachy/logs/log1.txt`.
3. Verify the log file was added: `pachctl list file lb-demo@master`.
   ```
   NAME      TYPE SIZE     
   /log1.txt file 292.3KiB 
   ```
4. Optionally, you can view the file: `pachctl get file lb-demo@master:log1.txt | open -f -a TextEdit.app`.
5. Repeat for as many files or directories necessary. 

**✏️ NOTE**: Pachyderm commits are similar to git commits. Here's a quick breakdown of the format:
   ```
   {reponame}@{branch}:filename.extension -f /Path/of/file.txt
   ```

**💡 TIP**: Have a directory of log files? You can use `pachctl put file -r repo@branch -f {dirName}` to upload it with one command. 

## 3. Create a Pipeline 

✅ We've got our repository set up. 

✅ We've committed log data.

Now, let's create a [pipeline](https://docs.pachyderm.com/latest/concepts/pipeline-concepts/pipeline/#pipeline). A pipeline reads, transforms, and outputs data. To use a pipeline, you must define a pipeline schema (either in `JSON` or `YAML`). 


### Pipeline Template 

```json 
{
    "pipeline": {
      "name": "lb-pipeline" // Displayed when using the following command: pachctl list pipeline
    },
    "description": "A pipeline that counts WARNING and ERROR occurrences in one or many log files.", // Displayed when using the following command: pachctl list pipeline
    "transform": {
      "cmd": [ "go run", "/count.go" ], // The command that executes the data transformation & output
      "image": "lbliii/lb-demo:1.0" // The Docker image containing the scripts/logic needed to transform the data. 
    },
    "input": {
      "pfs": { // The Pachyderm file system
        "repo": "lb-demo", // The repository name 
        "glob": "/*" 
      }
    }
  }
```

