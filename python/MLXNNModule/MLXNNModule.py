import mlx.core as mx
import mlx.nn as nn
import time

# Neural Network Model using mlx.nn.Module
class BinaryClassifier(nn.Module):
    def __init__(self, features):
        super().__init__()
        self.layer1 = nn.Linear(features, 128)   # First linear layer
        self.activation = nn.ReLU()             # Activation function
        self.dropout = nn.Dropout(p=0.5)        # Dropout layer
        self.layer2 = nn.Linear(128, 1)         # Second linear layer

    def __call__(self, x):
        x = self.layer1(x)
        x = self.activation(x)
        x = self.dropout(x)
        x = self.layer2(x)
        return x

# Parameters
examples = 1764
features = 42
trainingIterations = 5000
learningRate = 0.05

# True weight vector for dataset generation
trueWeightVector = mx.random.normal((features,))

# Training data
trainingData = mx.random.normal((examples, features))
outputLabels = (trainingData @ trueWeightVector) > 0

# Initialize the model
model = BinaryClassifier(features)

# Print model structure
print("Model Structure:")
print(model)
print("\nInitial Parameters:")
for layer_name, layer_params in model.parameters().items():
    for param_name, param in layer_params.items():
        print(f"{layer_name}.{param_name}: {param}")

# Initialize weights
for layer_params in model.parameters().values():
    for param in layer_params.values():
        param[:] = 1e-2 * mx.random.normal(param.shape)

def lossFunction(params, trainingData, outputLabels):
    model.update(params)
    logits = model(trainingData).squeeze()
    return mx.mean(mx.logaddexp(0.0, logits) - outputLabels * logits)

# Gradient function
loss_and_grad_fn = nn.value_and_grad(model, lossFunction)

# Training loop
startTime = time.time()
for i in range(trainingIterations):
    loss, grads = loss_and_grad_fn(model.trainable_parameters(), trainingData, outputLabels)

    for layer_name in model.trainable_parameters():
        layer_params = model.trainable_parameters()[layer_name]
        layer_grads = grads[layer_name]

        for param_name in layer_params:
            param = layer_params[param_name]
            grad = layer_grads[param_name]
            param[:] = param - learningRate * grad

    if i % 1000 == 0:
        print(f"Iteration {i}, Loss: {loss.item():.3f}")

    mx.eval(model.parameters())

finishTime = time.time()
throughput = trainingIterations / (finishTime - startTime)
print(f"\nThroughput: {throughput:.0f} iterations/second")

# Final Evaluation
loss = lossFunction(model, trainingData, outputLabels)
print(f"\nFinal Loss: {loss.item():.3f}")

# Accuracy
predictions = (model(trainingData).squeeze() > 0)
accuracy = mx.mean(predictions == outputLabels)
print(f"Accuracy: {accuracy.item():.3f}")

# Print updated parameters
print("\nUpdated Parameters:")
for layer_name, layer_params in model.parameters().items():
    for param_name, param in layer_params.items():
        print(f"{layer_name}.{param_name}: {param}")
