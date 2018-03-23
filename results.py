import requests
from bs4 import BeautifulSoup
import argparse

parser = argparse.ArgumentParser(description='Get Google Count.')
parser.add_argument('word', help='word to count')
args = parser.parse_args()

r = requests.get('http://www.google.com/search',
  params={'q':'"'+args.word+'"',
    "tbs":"li:1"}
)

soup = BeautifulSoup(r.text, "html.parser")
print soup.find('div',{'id':'resultStats'}).text
