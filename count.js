
// count the number of warnings and errors in the current file
// and return the results in a new text file

const fs = require('fs')
const path = require('path')

const count = (file) => {
    const filePath = path.join(__dirname, file)
    const content = fs.readFileSync(filePath, 'utf8')
    const lines = content.split('\n')
    const errors = lines.filter(line => line.includes('ERR')).length
    const warnings = lines.filter(line => line.includes('WARN')).length
    const results = `${errors} errors, ${warnings} warnings`
    const resultsPath = path.join(__dirname, 'results.txt')
    fs.writeFileSync(resultsPath, results)
    }

// for each file in /pfs/lb-demo, count the number of errors and warnings

for (const file of fs.readdirSync(path.join(__dirname, 'lb-demo'))) {
    count(file)
}