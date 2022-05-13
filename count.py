// count the number of warnings and errors in the current file
// and return the results in a new text file

def get_warnings(file):
    warnings = 0
    with open(file, 'r') as f:
        for line in f:
            if 'WARN' in line:
                warnings += 1
    return warnings

def get_errors(file):
    errors = 0
    with open(file, 'r') as f:
        for line in f:
            if 'ERROR' in line:
                errors += 1
    return errors

