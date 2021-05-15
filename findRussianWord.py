# python 默认编码为unicode
#这里匹配的是俄语字符所有unicode码
import re
russian_word = re.compile(u'[\u0411-\u04ff]+')
with open(r"source/source.txt") as file:
    words={}
    wordkey=[]
    flag = True
    while flag:
        line = file.readline()
        # check if endwith \n 
        if line.endswith("\n"):
            flag = True
        else:
            flag = False
        wordarr = russian_word.findall(line)
        for word in wordarr:
            if word in wordkey:
                words[word] = words[word]+1
            else:
                words[word] = 1
                wordkey.append(word)
               
