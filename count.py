import os

def main():
    warnings = 0 
    errors = 0 
    files = os.listdir('/pfs/lb-pachy-project')
    # print(files)
    for file in files:
        with open('/pfs/lb-pachy-project/' + file, 'r') as f:
            lines = f.readlines()
            for line in lines:
                # TODO: edge cases? 
                if 'WARN' in line:
                    warnings += 1
                if 'ERR' in line:
                    errors += 1
    write_results(warnings, errors)
    # print('Warnings: ' + str(warnings))
    # print('Errors: ' + str(errors))


# create a results.txt file and write the total warnings and errors.

def write_results(warnings, errors):
    with open('/pfs/out/results.txt', 'w') as f:
        f.write('Warnings: ' + str(warnings) + '\n')
        f.write('Errors: ' + str(errors))

  
main()

