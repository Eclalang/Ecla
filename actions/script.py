import os
from os.path import isdir, isfile
from json import dump

current_dir = os.getcwd()

def is_valid_dir(dir):
    return isdir(dir) and dir.split('/')[-1][0] != "."

def is_markdown_file(file):
    # print('split result: ', file.split(".")[-1] == "md")
    return isfile(file) and file.split(".")[-1] == "md"

def parse_directory(dir):
    resolved_files = []
    # print('parsing', dir)
    toparse_dir = os.listdir(dir)
    # print('toparse_dir: ',toparse_dir)
    for element in toparse_dir:
        # print('is_markdown_file: ', element, is_markdown_file(f'{dir}/{element}'))
        if is_markdown_file(f'{dir}/{element}'):
            # print(f'{dir}/{element}'.split(f'{current_dir}/')[1])
            resolved_files.append({
                "link": f'{dir}/{element}'.split(f'{current_dir}/')[1],
                "title": element.split(".md")[0]
            })
        # print('is_valid_dir: ', element, is_valid_dir(f'{dir}/{element}'))
        if is_valid_dir(f'{dir}/{element}'):
            # print(element)
            resolved_files = resolved_files + parse_directory(f'{dir}/{element}')
    return resolved_files

dump(parse_directory(current_dir), open('index.json', 'w'))
