# Solutions for Day 7.1
module Day7
  module_function

  def fuel_cost(data)
    positions = data.split(',').map(&:strip).map(&:to_i)
    cheapest = median(positions)
    positions
      .map { |pos| (pos - cheapest).abs }
      .sum
  end

  def median(array)
    sorted = array.sort
    len = sorted.length
    ((sorted[(len - 1) / 2] + sorted[len / 2]) / 2.0).round
  end

  def expensive_fuel_cost(data)
    positions = data.split(',').map(&:strip).map(&:to_i)
    cheapest = average(positions)
    positions
      .map { |pos| (pos - cheapest).abs.downto(1).reduce(&:+) }
      .compact
      .sum
  end

  def average(array)
    sum = array.sum
    (sum.to_f / array.length).round
  end
end
