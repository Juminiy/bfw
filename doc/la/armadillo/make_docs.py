import json 
import os

def parse_to_json(file_name):
    with open(file_name, 'r') as f:
        doc = f.read()

    lines = doc.strip().split('\n')
    json_objects = []
    for line in lines:
        if ';' in line:
            key, value = line.split(';', 1)
            key = key.rstrip("\t\n").lstrip("\t\n").strip()
            value = value.rstrip("\t\n").lstrip("\t\n").strip()
            json_object = {key: value}
            json_objects.append(json_object)
        else:
            print(f"Warning: Skipping line '{line}' because it does not contain a ';' symbol.")
    
    with open(file_name+'.json', 'w', encoding='utf-8') as f:
        json.dump(json_objects, f, ensure_ascii=False, indent=4)
    print("JSON对象已成功保存到",file_name,".json文件中")



current_dir = os.getcwd()

for filename in os.listdir(current_dir):
    if filename.endswith('.txt'):
        parse_to_json(filename)