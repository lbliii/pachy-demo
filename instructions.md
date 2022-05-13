## 1. Create a Repo 

1. Use the following command to build your repo:
    `pachctl create repo lb-demo`

{{% tip %}}
Use the command `pachctl list repo` to see all of your Pachyderm repositories. 
{{% /tip%}}

## 2. Upload Data (Log Files)

1. Use the following command to add a log file to the `lb-demo` repo: 
    `pachctl put file lb-demo@master:log1.txt -f /Users/lblane/Documents/pachy/logs/log1.txt`
2. Verify the log file was added:
    `pachctl list file lb-demo@master`
3. View
    `pachctl get file lb-demo@master:log1.txt | open -f -a Preview.app` 


## 3. Build Docker Image 

`docker build -t lbliii/lb-demo:1.0 .` 

## 4. Create a Pipeline 


