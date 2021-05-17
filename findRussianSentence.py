from IPython import embed
import re

flag = True
lines = ""
count = 0
with open("source/source.txt") as f:
    while flag:
        line = f.readline()
        if line.endswith("\n"):
            flag =True
        else:
            flag = False
        
        result = re.findall(r"\.",line)
        if len(result)>5:
            line = ""
        lines = lines + line
        count = count + 1
        splits = re.split(r"(\. |\.\n)",lines)
        splits = [split for split in splits if len(split) > 6]
    embed()
       
   
    


