run:
	screencapture -i tmp.jpg && tesseract tmp.jpg ocr_output -l eng && ./ethics-review

auto:
	screencapture -R1130,200,490,480 square.jpg && tesseract square.jpg ocr_output -l eng && ./ethics-review
