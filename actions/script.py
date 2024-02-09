import os
from os.path import isdir, isfile
from json import dump

current_dir = os.getcwd()
all_id = {}
current_id = 0

def is_valid_dir(dir):
    return isdir(dir) and dir.split('/')[-1][0] != "."

def is_markdown_file(file):
    # print('split result: ', file.split(".")[-1] == "md")
    return isfile(file) and file.split(".")[-1] == "md"

def calculate_id(file):
    h = hash("/".join(file.split("/").remove("FR")))
    if h in all_id:
        return all_id[h]
    else:
        all_id[h] = current_id
        current_id += 1
        return all_id[h]

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
                "title": element.split(".md")[0],
                "lang" : "fr" if "FR" in dir.split("/") else "en",
                "id" : calculate_id(f'{dir}/{element}')
            })
        # print('is_valid_dir: ', element, is_valid_dir(f'{dir}/{element}'))
        if is_valid_dir(f'{dir}/{element}'):
            # print(element)
            resolved_files = resolved_files + parse_directory(f'{dir}/{element}')
    return resolved_files

pars = parse_directory(current_dir)

pars.sort(key=lambda x: x["title"])

dump(pars, open('index.json', 'w'))

# check if the directory dist exists
if not isdir(f'{current_dir}/dist'):
    os.mkdir(f'{current_dir}/dist')

# copy the index.json file to the dist folder
os.system(f'cp {current_dir}/index.json {current_dir}/dist/index.json')

# copy the docs folder to the dist folder
# os.system(f'cp -r {current_dir}/docs {current_dir}/dist/docs')

# copy all the markdown files to the dist folder

for file in pars:

    # verify if all subdirectories exists
    all_dirs = file["link"].split("/")
    all_dirs.pop()
    p = current_dir + "/dist"
    for dir in all_dirs:
        p = p + "/" + dir
        if not isdir(p):
            os.mkdir(p)

    os.system(f'cp {current_dir}/{file["link"]} {current_dir}/dist/{file["link"]}')
