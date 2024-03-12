import sys

from handler import MyHandler

print('argument list', sys.argv)

handler = MyHandler()

words, word_start_times, word_end_times, duration_sec = handler.inference()

with open("/data/words.txt", "w") as file:
    file.writelines(words)

