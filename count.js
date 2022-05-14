
// count the number of warnings and errors in the current file
// and return the results in a new text file

const fs = require('fs')
const path = require('path')
const totalWarningsAndErrors = {warnings: 0, errors: 0}

for (const file of fs.readdirSync(path.join(__dirname, 'logs'))) {
    // console.log(file) 
    contents = fs.readFileSync(path.join(__dirname, 'logs', file), 'utf8')
    const lines = contents.split('\n')
    for (const line of lines) {
        if (line.includes('WARN')) {
            totalWarningsAndErrors.warnings++
        }
        if (line.includes('ERR')) {
            totalWarningsAndErrors.errors++
        }
    }
}

// console.log(totalWarningsAndErrors) 

const createResultsFile = () => {
    fs.writeFileSync(path.join(__dirname, 'results.txt'), JSON.stringify(totalWarningsAndErrors))
}

createResultsFile()