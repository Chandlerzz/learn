import re
russian_word = re.compile(u'[\u0411-\u04ff]+')
with open(r"source/source.txt") as file:
    for i in range(10):
        line = file.readline()
        wordarr = russian_word.findall(line)
        print(wordarr)
