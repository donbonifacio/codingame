
# Solutions for Day 1.1
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
