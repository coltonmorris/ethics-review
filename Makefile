run:
	screencapture -i tmp.jpg && tesseract tmp.jpg ocr_output -l eng && ./ethics-review
