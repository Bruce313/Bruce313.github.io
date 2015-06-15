#!/bin/bash

for md in md/*;do
    base=`basename $md`
    node rep.js $md "html/$base.html"
done
