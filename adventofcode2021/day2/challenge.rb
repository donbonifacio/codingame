
# Solutions for Day 2.1
module Day2
  module_function

  VECTOR = {
    'forward' => [1, 0],
    'down' => [0, 1],
    'up' => [0, -1]
  }.freeze

  def position_factor(instructions)
    instructions
      .lines
      .map { |line| line.split(/\s/) }
      .map(&method(:build_vector))
      .reduce([0, 0], &method(:apply_vector))
      .reduce(&:*)
  end

  def build_vector(data)
    dir = VECTOR.fetch(data.first)
    strength = data.last.to_i
    [
      dir.first * strength,
      dir.last * strength
    ]
  end

  def apply_vector(position, vector)
    [
      position.first + vector.first,
      position.last + vector.last
    ]
  end
end

# Solutions for Day 2.2
module Day2Part2
  module_function

  VECTOR = {
    'forward' => [1, 0, 0],
    'down' => [0, 0, 1],
    'up' => [0, 0, -1]
  }.freeze

  def position_factor(instructions)
    instructions
      .lines
      .map { |line| line.split(/\s/) }
      .map(&method(:build_vector))
      .reduce([0, 0, 0], &method(:apply_vector))
      .take(2)
      .reduce(&:*)
  end

  def build_vector(data)
    dir = VECTOR.fetch(data.first)
    strength = data.last.to_i
    [
      dir[0] * strength,
      dir[1] * strength,
      dir[2] * strength
    ]
  end

  def apply_vector(position, vector)
    [
      position[0] + vector[0],
      position[1] + (vector[0] * position[2]),
      position[2] + vector[2]
    ]
  end
end
