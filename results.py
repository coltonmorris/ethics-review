import requests
from bs4 import BeautifulSoup

output = ""
f = open("ocr_output.txt")
data = f.readlines()
for line in data:
    # parse input, assign values to variables
    value = line
    output += value
f.close()

output = output.split("\n\n")
output[0] = output[0].replace("\n", " ")

for i in range(1,len(output)):
    print output[0]+" "+output[i]
    r = requests.get('http://www.google.com/search',
            params={'q':'"'+output[0]+" "+output[i]+'"',
                "tbs":"li:1"}
            )

    soup = BeautifulSoup(r.text, "html.parser")
    print soup.find('div',{'id':'resultStats'}).text

