#!/bin/bash

# Define the snippet name
snippetname="rockPaperScissors"

# Create the folder
mkdir "$snippetname"

# Go to the folder
cd "$snippetname"

# Create main.py file
cat <<EOF >$snippetname.py
import numpy as np
from keras.models import Sequential
from keras.layers import LSTM, Dense
from keras.utils import to_categorical

class RPS_AI:
    def __init__(self):
        self.model = self.build_model()
        self.history = []
        self.train_X = []
        self.train_y = []
        self.choices = ['rock', 'paper', 'scissors']

    def build_model(self):
        model = Sequential()
        model.add(LSTM(50, activation='relu', input_shape=(3, 2)))
        model.add(Dense(3, activation='softmax'))
        model.compile(optimizer='adam', loss='categorical_crossentropy')
        return model

    def predict(self):
        if len(self.history) < 3:
            return np.random.choice(self.choices)
        else:
            last_sequence = np.array(self.history[-3:]).reshape((1, 3, 2))
            prediction = self.model.predict(last_sequence, verbose=0)
            predicted_move = self.choices[np.argmax(prediction)]
            # Determine the winning move based on the predicted move
            if predicted_move == 'rock':
                return 'paper'
            elif predicted_move == 'scissors':
                return 'rock'
            else:  # predicted_move == 'paper'
                return 'scissors'

    def update_history(self, move, reward):
        self.history.append((self.choices.index(move), reward))
        if len(self.history) > 4:
            self.train_X.append([x for x in self.history[-5:-2]])
            self.train_y.append(to_categorical(self.history[-2][0], num_classes=3))
            self.model.fit(np.array(self.train_X), np.array(self.train_y), epochs=200, verbose=0)

def play():
    ai = RPS_AI()
    ai_score = 0
    player_score = 0
    for _ in range(100):  # play 100 rounds
        player_move = input("Enter your move (rock, paper, scissors): ")
        ai_move = ai.predict()
        print(f"AI played: {ai_move}")
        # determine winner...
        if player_move == 'rock' and ai_move == 'scissors':
            player_score += 1
            reward = 1
        elif player_move == 'scissors' and ai_move == 'paper':
            player_score += 1
            reward = 1
        elif player_move == 'paper' and ai_move == 'rock':
            player_score += 1
            reward = 1
        elif ai_move == 'rock' and player_move == 'scissors':
            ai_score += 1
            reward = -1
        elif ai_move == 'scissors' and player_move == 'paper':
            ai_score += 1
            reward = -1
        elif ai_move == 'paper' and player_move == 'rock':
            ai_score += 1
            reward = -1
        else:
            reward = 0
        ai.update_history(player_move, reward)  # directly pass the player's move and the reward
        print(f"-")
        print(f"Current score: Player {player_score}, AI {ai_score}")  # print the score after each round

print(play())
EOF

# Create Makefile
cat <<EOF >Makefile
BINARY=$snippetname

run_py: build_py
	python3 \$(BINARY).py

build_py:
	pip3 install tensorflow
	echo "Python script ready to be run"

clean_py:
	rm -f \$(BINARY).py

clean_binaries:
	rm -f ./*pyc
EOF

make
