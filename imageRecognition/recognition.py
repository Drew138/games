#!/usr/local/bin/python

import Image
import pytesseract
import sys


def ocr_core(file):
    text = pytesseract.image_to_string(file)
    return text


print(ocr_core(sys.argv[1]))
