output = {}
f = open("ocr_output.txt")
data = f.readlines()
for line in data:
    # parse input, assign values to variables
    key, value = line.split("\n")
    output[key.strip()] = value.strip()
f.close()

print output
