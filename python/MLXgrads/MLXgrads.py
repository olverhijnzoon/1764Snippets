import mlx.core
import time

#  1764 Snippets
## MLX Gradients

examples = 1764
features = 42
trainingIterations = 5000
learningRate = 0.05

# True weight vector s the ideal set of weights that perfectly maps the input features to the output labels in the generated dataset. In a real-world scenario, this would be unknown, but in this synthetic example, we generate it ourselves for the purpose of creating the dataset.
trueWeightVector = mlx.core.random.normal((features,))

# Matrix of size (examples,features) which represents the input features  
trainingData = mlx.core.random.normal((examples, features))

# Target values "labels" for the examples in X (or that "correspond to each example"); dot product @
outputLabels = (trainingData @ trueWeightVector) > 0

# set of weights that the model learns during the training process
learnedWeightVector = 1e-2 * mlx.core.random.normal((features,))

# log loss function for a binary classification task
def lossFunction(learnedWeightVector):
    # raw predictions "logits" from the model
    logOdds = trainingData @ learnedWeightVector
    # discrepancy between the true label and the predicted probability is lower when the models predictions are closer to the true labels
    return mlx.core.mean(mlx.core.logaddexp(0.0, logOdds) - outputLabels * logOdds)

# this derivative "gradient" is used in the training loop to update the models parameters
gradientFunction = mlx.core.grad(lossFunction)

# start time measurement
startTime = time.time()

for _ in range(trainingIterations):
    gradient = gradientFunction(learnedWeightVector)
    learnedWeightVector = learnedWeightVector - learningRate * gradient
    mlx.core.eval(learnedWeightVector)

# stop time measurement
finishTime = time.time()

# calculate throughput
throughput = trainingIterations / (finishTime - startTime)
print(f"throughput {throughput:.0f} (it/s)")

# calculate loss
loss = lossFunction(learnedWeightVector)
print(f"loss {loss.item():.3f}")

# generates the final predictions of the binary classification model and calculates accuracy
finalPredictions = (trainingData @ learnedWeightVector) > 0
accuracy = mlx.core.mean(finalPredictions == outputLabels)
print(f"accuracy {accuracy.item():.3f} ")