# Solutions for Day 7.1
module Day7
  module_function

  def fuel_cost(data)
    positions = data.split(',').map(&:strip).map(&:to_i)
    cheapest = median(positions)
    positions
      .map { |pos| (pos - cheapest).abs }
      .reduce(&:+)
  end

  def median(array)
    sorted = array.sort
    len = sorted.length
    ((sorted[(len - 1) / 2] + sorted[len / 2]) / 2.0).round
  end
end
