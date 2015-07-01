#!/usr/bin/python
#encoding=utf-8
import os
import markdown2
mdDir = './md'
htmlDir = './html'

temp = open('./template.html').read()
for root, dirs, files in os.walk(mdDir):
    for f in files:
        html = markdown2.markdown_path(os.path.join(root, f))
        gen = temp.replace('CONTENT', html).encode('utf-8')
        open(os.path.join(htmlDir, f + '.html'), 'w').write(gen)
