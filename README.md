# Island Ecology Simulation

This Python script simulates the ecology of an island, focusing on the growth and spread of plants over time.

## Introduction

The simulation models the growth of plants on a virtual island, considering factors such as initial plant distribution, growth rate, and spread to adjacent cells. The goal is to observe how the plant population evolves over time and understand the dynamics of the ecosystem.

## Features

- **Random Initialization**: The initial distribution of plants is randomized across the island.
- **Plant Growth**: Plants grow over time according to a specified growth rate.
- **Spread Mechanism**: Plants spread to adjacent cells, contributing to the overall plant density.
- **Visualization**: The simulation provides visualizations using Matplotlib, including a matrix plot of plant distribution and a line plot of total plant count over time.

## Requirements

- Python 3.x
- NumPy
- Matplotlib

## Usage

1. Clone the repository:

   ```bash
   git clone https://github.com/your_username/island-ecology-simulation.git
   ```

2. Navigate to the project directory:

   ```bash
   cd island-ecology-simulation
   ```

3. Run the Python script:

   ```bash
   python island_ecology_simulation.py
   ```

4. View the animation in your web browser.

## Parameters

- **Island Size**: The dimensions of the island (e.g., 20x20).
- **Initial Plant Distribution**: The number of plants initially distributed across the island.
- **Growth Rate**: The rate at which plants grow over time.
- **Animation Frames**: The number of frames for the animation.
- **Animation Interval**: The interval between frames in milliseconds.

## Example

```python
# Define simulation parameters
islandSizes = 20
numberOfPlants = 5
growthRate = 0.1

# Run the simulation
python island_ecology_simulation.py
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by ecology and simulation studies.
- Thanks to contributors and open-source libraries.
