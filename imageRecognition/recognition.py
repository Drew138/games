#!/usr/local/bin/python

from PIL import Image
import pytesseract
import sys
import io


def ocr_core(image):
    image = Image.open(io.BytesIO(image))
    text = pytesseract.image_to_string(image)
    return text


# sys.argv[1]
print(ocr_core())
