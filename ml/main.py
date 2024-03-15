import json
import sys

from handler import MyHandler

print('argument list', sys.argv)

handler = MyHandler()

words, word_start_times, word_end_times, duration_sec = handler.inference()

data = (words, word_start_times, word_end_times, duration_sec)

with open('/data/speech.json', 'w') as file:
    json.dump(data, file)
