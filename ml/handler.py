import torch
from transformers import Wav2Vec2Processor, Wav2Vec2ForCTC
import librosa
from itertools import groupby

class MyHandler():
    def __init__(self):
        super(MyHandler, self).__init__()
        self.processor = Wav2Vec2Processor.from_pretrained("jonatasgrosman/wav2vec2-large-xlsr-53-russian")
        self.model = Wav2Vec2ForCTC.from_pretrained("jonatasgrosman/wav2vec2-large-xlsr-53-russian")

    def inference(self, data):
        with torch.no_grad():
            speech_array, sampling_rate = librosa.load(data, sr=16_000)
            data = self.processor(speech_array, sampling_rate=sampling_rate, return_tensors="pt", padding=True)
            logits = self.model(data.input_values, attention_mask=data.attention_mask).logits
            predicted_ids = torch.argmax(logits, dim=-1)
            transcription = self.processor.batch_decode(predicted_ids)[0]
            words = [w for w in transcription.split(' ') if len(w) > 0]
            predicted_ids = predicted_ids[0].tolist()
            duration_sec = data.input_values.shape[1] / 16000
            ids_w_time = [(i / len(predicted_ids) * duration_sec, _id) for i, _id in enumerate(predicted_ids)]
            ids_w_time = [i for i in ids_w_time if i[1] != self.processor.tokenizer.pad_token_id]
            split_ids_w_time = [list(group) for k, group in groupby(ids_w_time, lambda x: x[1] == self.processor.tokenizer.word_delimiter_token_id) if not k]
            assert len(split_ids_w_time) == len(words)
            word_start_times = []
            word_end_times = []
            for cur_ids_w_time, cur_word in zip(split_ids_w_time, words):
              _times = [_time for _time, _id in cur_ids_w_time]
              word_start_times.append(min(_times))
              word_end_times.append(max(_times))

            return words, word_start_times, word_end_times, duration_sec
